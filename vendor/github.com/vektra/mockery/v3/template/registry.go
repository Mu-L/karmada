package template

import (
	"context"
	"fmt"
	"go/types"
	"sort"

	"github.com/rs/zerolog"
	"github.com/vektra/mockery/v3/internal/stackerr"
	"golang.org/x/tools/go/packages"
)

// Registry encapsulates types information for the source and mock
// destination package. For the mock package, it tracks the list of
// imports and ensures there are no conflicts in the imported package
// qualifiers.
type Registry struct {
	dstPkgPath       string
	srcPkg           *packages.Package
	srcPkgName       string
	imports          map[string]*Package
	importQualifiers map[string]*Package
	// inPackage specifies whether this registry is considered to be in the same
	// package as the srcPkg. This is needed because of the way that Go package
	// qualifiers work. For example, test files for a package are allowed to have
	// a different package name than the files under test, in which case they are
	// considered to be in a separate package. In such a case, `inPackage` should
	// be set to false such that calls to AddImport for the original source package
	// are not ignored. Otherwise if it's set to true, AddImport ignores imports
	// for the package in which the file already resides.
	inPackage bool
}

// New loads the source package info and returns a new instance of
// Registry.
func NewRegistry(srcPkg *packages.Package, dstPkgPath string, inPackage bool) (*Registry, error) {
	return &Registry{
		dstPkgPath:       dstPkgPath,
		srcPkg:           srcPkg,
		imports:          make(map[string]*Package),
		importQualifiers: make(map[string]*Package),
		inPackage:        inPackage,
	}, nil
}

func (r Registry) SrcPkg() *packages.Package {
	return r.srcPkg
}

// SrcPkgName returns the name of the source package.
func (r Registry) SrcPkgName() string {
	return r.srcPkg.Name
}

// LookupInterface returns the underlying interface definition of the
// given interface name.
func (r Registry) LookupInterface(name string) (*types.Interface, *types.TypeParamList, error) {
	obj := r.SrcPkg().Types.Scope().Lookup(name)
	if obj == nil {
		return nil, nil, stackerr.NewStackErr(fmt.Errorf("interface not found: %s", name))
	}

	if !types.IsInterface(obj.Type()) {
		return nil, nil, fmt.Errorf("%s (%s) is not an interface", name, obj.Type())
	}

	var tparams *types.TypeParamList
	named, ok := obj.Type().(*types.Named)
	if ok {
		tparams = named.TypeParams()
	}

	return obj.Type().Underlying().(*types.Interface).Complete(), tparams, nil
}

// MethodScope returns a new MethodScope.
func (r *Registry) MethodScope() *MethodScope {
	return NewMethodScope(r)
}

type fakeTypesPackage struct {
	name string
	path string
}

func (f fakeTypesPackage) Name() string {
	return f.name
}

func (f fakeTypesPackage) Path() string {
	return f.path
}

// addImport adds the given package to the set of imports. It generates a
// suitable alias if there are any conflicts with previously imported
// packages. pkgName must be set to the unaliased package name.
func (r *Registry) AddImport(pkgName string, pkgPath string) *Package {
	// Note: Yes this method is a little weird. This is intended to be used
	// by templates that want to add their own imports. Instead of requiring the
	// templates to pass in a ctx and TypesPackage instance, we create this new
	// AddImport method that wraps around r.addImport. r.addImport still exists
	// because mockery will add its own imports based on the existing types, and
	// in that case we want it to pass ctx (that contains the logger) and the
	// real types.Package type.
	return r.addImport(context.Background(), fakeTypesPackage{
		name: pkgName,
		path: pkgPath,
	})
}

func (r *Registry) addImport(ctx context.Context, pkg TypesPackage) *Package {
	path := pkg.Path()
	logContext := zerolog.Ctx(ctx).With().
		Str("method", "AddImport").
		Str("dst-pkg-path", r.dstPkgPath).
		Str("pkg-path", path).
		Bool("inpackage", r.inPackage)
	if r.srcPkg != nil {
		logContext = logContext.Str("src-pkg-path", r.srcPkg.PkgPath)
	}
	log := logContext.Logger()

	if r.srcPkg != nil && path == r.srcPkg.PkgPath && r.inPackage {
		log.Debug().Msg("package path equals src-pkg-path, not adding import")
		return nil
	} else {
		log.Debug().Msg("package path does not equal src-pkg-path, adding import")
	}

	if imprt, ok := r.imports[path]; ok {
		return imprt
	}

	imprt := Package{pkg: pkg}
	originalQualifier := imprt.Qualifier()
	var aliasSuggestion string = imprt.Qualifier()
	for i := 0; ; i++ {
		if _, conflict := r.importQualifiers[aliasSuggestion]; conflict {
			aliasSuggestion = fmt.Sprintf("%s%d", imprt.Qualifier(), i)
			continue
		}
		if originalQualifier != aliasSuggestion {
			imprt.Alias = aliasSuggestion
		}
		break
	}

	r.imports[path] = &imprt
	r.importQualifiers[imprt.Qualifier()] = &imprt
	return &imprt
}

// Imports returns the list of imported packages. The list is sorted by
// path.
func (r Registry) Imports() Packages {
	imports := make([]*Package, 0, len(r.imports))
	for _, imprt := range r.imports {
		imports = append(imports, imprt)
	}
	sort.Slice(imports, func(i, j int) bool {
		return imports[i].Path() < imports[j].Path()
	})
	return imports
}

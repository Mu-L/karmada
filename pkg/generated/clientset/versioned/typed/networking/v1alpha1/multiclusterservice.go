/*
Copyright The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/networking/v1alpha1"
	scheme "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// MultiClusterServicesGetter has a method to return a MultiClusterServiceInterface.
// A group's client should implement this interface.
type MultiClusterServicesGetter interface {
	MultiClusterServices(namespace string) MultiClusterServiceInterface
}

// MultiClusterServiceInterface has methods to work with MultiClusterService resources.
type MultiClusterServiceInterface interface {
	Create(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.CreateOptions) (*v1alpha1.MultiClusterService, error)
	Update(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (*v1alpha1.MultiClusterService, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (*v1alpha1.MultiClusterService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MultiClusterService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MultiClusterServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterService, err error)
	MultiClusterServiceExpansion
}

// multiClusterServices implements MultiClusterServiceInterface
type multiClusterServices struct {
	*gentype.ClientWithList[*v1alpha1.MultiClusterService, *v1alpha1.MultiClusterServiceList]
}

// newMultiClusterServices returns a MultiClusterServices
func newMultiClusterServices(c *NetworkingV1alpha1Client, namespace string) *multiClusterServices {
	return &multiClusterServices{
		gentype.NewClientWithList[*v1alpha1.MultiClusterService, *v1alpha1.MultiClusterServiceList](
			"multiclusterservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.MultiClusterService { return &v1alpha1.MultiClusterService{} },
			func() *v1alpha1.MultiClusterServiceList { return &v1alpha1.MultiClusterServiceList{} }),
	}
}

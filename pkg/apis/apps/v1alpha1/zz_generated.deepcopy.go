//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservedWorkload) DeepCopyInto(out *ObservedWorkload) {
	*out = *in
	out.Workload = in.Workload
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservedWorkload.
func (in *ObservedWorkload) DeepCopy() *ObservedWorkload {
	if in == nil {
		return nil
	}
	out := new(ObservedWorkload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRebalancer) DeepCopyInto(out *WorkloadRebalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRebalancer.
func (in *WorkloadRebalancer) DeepCopy() *WorkloadRebalancer {
	if in == nil {
		return nil
	}
	out := new(WorkloadRebalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadRebalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRebalancerList) DeepCopyInto(out *WorkloadRebalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadRebalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRebalancerList.
func (in *WorkloadRebalancerList) DeepCopy() *WorkloadRebalancerList {
	if in == nil {
		return nil
	}
	out := new(WorkloadRebalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadRebalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRebalancerSpec) DeepCopyInto(out *WorkloadRebalancerSpec) {
	*out = *in
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRebalancerSpec.
func (in *WorkloadRebalancerSpec) DeepCopy() *WorkloadRebalancerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadRebalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRebalancerStatus) DeepCopyInto(out *WorkloadRebalancerStatus) {
	*out = *in
	if in.ObservedWorkloads != nil {
		in, out := &in.ObservedWorkloads, &out.ObservedWorkloads
		*out = make([]ObservedWorkload, len(*in))
		copy(*out, *in)
	}
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRebalancerStatus.
func (in *WorkloadRebalancerStatus) DeepCopy() *WorkloadRebalancerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadRebalancerStatus)
	in.DeepCopyInto(out)
	return out
}

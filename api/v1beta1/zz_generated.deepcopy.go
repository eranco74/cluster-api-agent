//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentBootstrapConfig) DeepCopyInto(out *AgentBootstrapConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentBootstrapConfig.
func (in *AgentBootstrapConfig) DeepCopy() *AgentBootstrapConfig {
	if in == nil {
		return nil
	}
	out := new(AgentBootstrapConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentBootstrapConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentBootstrapConfigList) DeepCopyInto(out *AgentBootstrapConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AgentBootstrapConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentBootstrapConfigList.
func (in *AgentBootstrapConfigList) DeepCopy() *AgentBootstrapConfigList {
	if in == nil {
		return nil
	}
	out := new(AgentBootstrapConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentBootstrapConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentBootstrapConfigSpec) DeepCopyInto(out *AgentBootstrapConfigSpec) {
	*out = *in
	if in.InfraEnvRef != nil {
		in, out := &in.InfraEnvRef, &out.InfraEnvRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.PullSecretRef != nil {
		in, out := &in.PullSecretRef, &out.PullSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentBootstrapConfigSpec.
func (in *AgentBootstrapConfigSpec) DeepCopy() *AgentBootstrapConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AgentBootstrapConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentBootstrapConfigStatus) DeepCopyInto(out *AgentBootstrapConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentBootstrapConfigStatus.
func (in *AgentBootstrapConfigStatus) DeepCopy() *AgentBootstrapConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AgentBootstrapConfigStatus)
	in.DeepCopyInto(out)
	return out
}
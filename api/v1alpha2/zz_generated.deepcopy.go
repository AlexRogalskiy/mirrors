//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMirror) DeepCopyInto(out *SecretMirror) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMirror.
func (in *SecretMirror) DeepCopy() *SecretMirror {
	if in == nil {
		return nil
	}
	out := new(SecretMirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretMirror) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMirrorDestination) DeepCopyInto(out *SecretMirrorDestination) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Vault = in.Vault
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMirrorDestination.
func (in *SecretMirrorDestination) DeepCopy() *SecretMirrorDestination {
	if in == nil {
		return nil
	}
	out := new(SecretMirrorDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMirrorList) DeepCopyInto(out *SecretMirrorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretMirror, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMirrorList.
func (in *SecretMirrorList) DeepCopy() *SecretMirrorList {
	if in == nil {
		return nil
	}
	out := new(SecretMirrorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretMirrorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMirrorSource) DeepCopyInto(out *SecretMirrorSource) {
	*out = *in
	out.Vault = in.Vault
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMirrorSource.
func (in *SecretMirrorSource) DeepCopy() *SecretMirrorSource {
	if in == nil {
		return nil
	}
	out := new(SecretMirrorSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMirrorSpec) DeepCopyInto(out *SecretMirrorSpec) {
	*out = *in
	out.Source = in.Source
	in.Destination.DeepCopyInto(&out.Destination)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMirrorSpec.
func (in *SecretMirrorSpec) DeepCopy() *SecretMirrorSpec {
	if in == nil {
		return nil
	}
	out := new(SecretMirrorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMirrorStatus) DeepCopyInto(out *SecretMirrorStatus) {
	*out = *in
	in.LastSyncTime.DeepCopyInto(&out.LastSyncTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMirrorStatus.
func (in *SecretMirrorStatus) DeepCopy() *SecretMirrorStatus {
	if in == nil {
		return nil
	}
	out := new(SecretMirrorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAppRoleAuthSpec) DeepCopyInto(out *VaultAppRoleAuthSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAppRoleAuthSpec.
func (in *VaultAppRoleAuthSpec) DeepCopy() *VaultAppRoleAuthSpec {
	if in == nil {
		return nil
	}
	out := new(VaultAppRoleAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthSpec) DeepCopyInto(out *VaultAuthSpec) {
	*out = *in
	out.AppRole = in.AppRole
	out.Token = in.Token
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthSpec.
func (in *VaultAuthSpec) DeepCopy() *VaultAuthSpec {
	if in == nil {
		return nil
	}
	out := new(VaultAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSpec) DeepCopyInto(out *VaultSpec) {
	*out = *in
	out.Auth = in.Auth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSpec.
func (in *VaultSpec) DeepCopy() *VaultSpec {
	if in == nil {
		return nil
	}
	out := new(VaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultTokenAuthSpec) DeepCopyInto(out *VaultTokenAuthSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultTokenAuthSpec.
func (in *VaultTokenAuthSpec) DeepCopy() *VaultTokenAuthSpec {
	if in == nil {
		return nil
	}
	out := new(VaultTokenAuthSpec)
	in.DeepCopyInto(out)
	return out
}

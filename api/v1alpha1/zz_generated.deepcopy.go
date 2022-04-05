// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EventTypes) DeepCopyInto(out *EventTypes) {
	{
		in := &in
		*out = make(EventTypes, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypes.
func (in EventTypes) DeepCopy() EventTypes {
	if in == nil {
		return nil
	}
	out := new(EventTypes)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in IgnoreReasons) DeepCopyInto(out *IgnoreReasons) {
	{
		in := &in
		*out = make(IgnoreReasons, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreReasons.
func (in IgnoreReasons) DeepCopy() IgnoreReasons {
	if in == nil {
		return nil
	}
	out := new(IgnoreReasons)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in InvolvedObjectKinds) DeepCopyInto(out *InvolvedObjectKinds) {
	{
		in := &in
		*out = make(InvolvedObjectKinds, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvolvedObjectKinds.
func (in InvolvedObjectKinds) DeepCopy() InvolvedObjectKinds {
	if in == nil {
		return nil
	}
	out := new(InvolvedObjectKinds)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in InvolvedObjectNames) DeepCopyInto(out *InvolvedObjectNames) {
	{
		in := &in
		*out = make(InvolvedObjectNames, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvolvedObjectNames.
func (in InvolvedObjectNames) DeepCopy() InvolvedObjectNames {
	if in == nil {
		return nil
	}
	out := new(InvolvedObjectNames)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	if in.Observes != nil {
		in, out := &in.Observes, &out.Observes
		*out = make([]Observe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Listeners) DeepCopyInto(out *Listeners) {
	{
		in := &in
		*out = make(Listeners, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listeners.
func (in Listeners) DeepCopy() Listeners {
	if in == nil {
		return nil
	}
	out := new(Listeners)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loudspeaker) DeepCopyInto(out *Loudspeaker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loudspeaker.
func (in *Loudspeaker) DeepCopy() *Loudspeaker {
	if in == nil {
		return nil
	}
	out := new(Loudspeaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Loudspeaker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoudspeakerList) DeepCopyInto(out *LoudspeakerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Loudspeaker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoudspeakerList.
func (in *LoudspeakerList) DeepCopy() *LoudspeakerList {
	if in == nil {
		return nil
	}
	out := new(LoudspeakerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoudspeakerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoudspeakerSpec) DeepCopyInto(out *LoudspeakerSpec) {
	*out = *in
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make(Listeners, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoudspeakerSpec.
func (in *LoudspeakerSpec) DeepCopy() *LoudspeakerSpec {
	if in == nil {
		return nil
	}
	out := new(LoudspeakerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoudspeakerStatus) DeepCopyInto(out *LoudspeakerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoudspeakerStatus.
func (in *LoudspeakerStatus) DeepCopy() *LoudspeakerStatus {
	if in == nil {
		return nil
	}
	out := new(LoudspeakerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observe) DeepCopyInto(out *Observe) {
	*out = *in
	if in.IgnoreReasons != nil {
		in, out := &in.IgnoreReasons, &out.IgnoreReasons
		*out = make(IgnoreReasons, len(*in))
		copy(*out, *in)
	}
	if in.InvolvedObjectNames != nil {
		in, out := &in.InvolvedObjectNames, &out.InvolvedObjectNames
		*out = make(InvolvedObjectNames, len(*in))
		copy(*out, *in)
	}
	if in.InvolvedObjectKinds != nil {
		in, out := &in.InvolvedObjectKinds, &out.InvolvedObjectKinds
		*out = make(InvolvedObjectKinds, len(*in))
		copy(*out, *in)
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make(EventTypes, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observe.
func (in *Observe) DeepCopy() *Observe {
	if in == nil {
		return nil
	}
	out := new(Observe)
	in.DeepCopyInto(out)
	return out
}

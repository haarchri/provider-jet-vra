//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreVsphere) DeepCopyInto(out *DatastoreVsphere) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreVsphere.
func (in *DatastoreVsphere) DeepCopy() *DatastoreVsphere {
	if in == nil {
		return nil
	}
	out := new(DatastoreVsphere)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreVsphere) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreVsphereList) DeepCopyInto(out *DatastoreVsphereList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatastoreVsphere, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreVsphereList.
func (in *DatastoreVsphereList) DeepCopy() *DatastoreVsphereList {
	if in == nil {
		return nil
	}
	out := new(DatastoreVsphereList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreVsphereList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreVsphereObservation) DeepCopyInto(out *DatastoreVsphereObservation) {
	*out = *in
	if in.CloudAccountIds != nil {
		in, out := &in.CloudAccountIds, &out.CloudAccountIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.ExternalRegionID != nil {
		in, out := &in.ExternalRegionID, &out.ExternalRegionID
		*out = new(string)
		**out = **in
	}
	if in.FreeSizeGb != nil {
		in, out := &in.FreeSizeGb, &out.FreeSizeGb
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreVsphereObservation.
func (in *DatastoreVsphereObservation) DeepCopy() *DatastoreVsphereObservation {
	if in == nil {
		return nil
	}
	out := new(DatastoreVsphereObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreVsphereParameters) DeepCopyInto(out *DatastoreVsphereParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreVsphereParameters.
func (in *DatastoreVsphereParameters) DeepCopy() *DatastoreVsphereParameters {
	if in == nil {
		return nil
	}
	out := new(DatastoreVsphereParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreVsphereSpec) DeepCopyInto(out *DatastoreVsphereSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreVsphereSpec.
func (in *DatastoreVsphereSpec) DeepCopy() *DatastoreVsphereSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreVsphereSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreVsphereStatus) DeepCopyInto(out *DatastoreVsphereStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreVsphereStatus.
func (in *DatastoreVsphereStatus) DeepCopy() *DatastoreVsphereStatus {
	if in == nil {
		return nil
	}
	out := new(DatastoreVsphereStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksObservation) DeepCopyInto(out *LinksObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
	if in.Hrefs != nil {
		in, out := &in.Hrefs, &out.Hrefs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Rel != nil {
		in, out := &in.Rel, &out.Rel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksObservation.
func (in *LinksObservation) DeepCopy() *LinksObservation {
	if in == nil {
		return nil
	}
	out := new(LinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksParameters) DeepCopyInto(out *LinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksParameters.
func (in *LinksParameters) DeepCopy() *LinksParameters {
	if in == nil {
		return nil
	}
	out := new(LinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsObservation) DeepCopyInto(out *TagsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsObservation.
func (in *TagsObservation) DeepCopy() *TagsObservation {
	if in == nil {
		return nil
	}
	out := new(TagsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsParameters) DeepCopyInto(out *TagsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsParameters.
func (in *TagsParameters) DeepCopy() *TagsParameters {
	if in == nil {
		return nil
	}
	out := new(TagsParameters)
	in.DeepCopyInto(out)
	return out
}

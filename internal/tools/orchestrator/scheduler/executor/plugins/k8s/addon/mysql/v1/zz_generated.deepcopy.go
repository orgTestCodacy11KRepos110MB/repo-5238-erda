// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mysql) DeepCopyInto(out *Mysql) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mysql.
func (in *Mysql) DeepCopy() *Mysql {
	if in == nil {
		return nil
	}
	out := new(Mysql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mysql) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlList) DeepCopyInto(out *MysqlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mysql, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlList.
func (in *MysqlList) DeepCopy() *MysqlList {
	if in == nil {
		return nil
	}
	out := new(MysqlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MysqlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlSolo) DeepCopyInto(out *MysqlSolo) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlSolo.
func (in *MysqlSolo) DeepCopy() *MysqlSolo {
	if in == nil {
		return nil
	}
	out := new(MysqlSolo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlSoloSpec) DeepCopyInto(out *MysqlSoloSpec) {
	*out = *in
	if in.SourceId != nil {
		in, out := &in.SourceId, &out.SourceId
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlSoloSpec.
func (in *MysqlSoloSpec) DeepCopy() *MysqlSoloSpec {
	if in == nil {
		return nil
	}
	out := new(MysqlSoloSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlSoloStatus) DeepCopyInto(out *MysqlSoloStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlSoloStatus.
func (in *MysqlSoloStatus) DeepCopy() *MysqlSoloStatus {
	if in == nil {
		return nil
	}
	out := new(MysqlSoloStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlSpec) DeepCopyInto(out *MysqlSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.PrimaryId != nil {
		in, out := &in.PrimaryId, &out.PrimaryId
		*out = new(int)
		**out = **in
	}
	if in.AutoSwitch != nil {
		in, out := &in.AutoSwitch, &out.AutoSwitch
		*out = new(bool)
		**out = **in
	}
	out.StorageSize = in.StorageSize.DeepCopy()
	if in.Solos != nil {
		in, out := &in.Solos, &out.Solos
		*out = make([]MysqlSoloSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExporterFlags != nil {
		in, out := &in.ExporterFlags, &out.ExporterFlags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlSpec.
func (in *MysqlSpec) DeepCopy() *MysqlSpec {
	if in == nil {
		return nil
	}
	out := new(MysqlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlStatus) DeepCopyInto(out *MysqlStatus) {
	*out = *in
	out.Version = in.Version
	if in.Solos != nil {
		in, out := &in.Solos, &out.Solos
		*out = make([]MysqlSolo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.MysqlSoloStatus = in.MysqlSoloStatus
	if in.WriteId != nil {
		in, out := &in.WriteId, &out.WriteId
		*out = new(int)
		**out = **in
	}
	if in.ReadId != nil {
		in, out := &in.ReadId, &out.ReadId
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlStatus.
func (in *MysqlStatus) DeepCopy() *MysqlStatus {
	if in == nil {
		return nil
	}
	out := new(MysqlStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlVersion) DeepCopyInto(out *MysqlVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlVersion.
func (in *MysqlVersion) DeepCopy() *MysqlVersion {
	if in == nil {
		return nil
	}
	out := new(MysqlVersion)
	in.DeepCopyInto(out)
	return out
}

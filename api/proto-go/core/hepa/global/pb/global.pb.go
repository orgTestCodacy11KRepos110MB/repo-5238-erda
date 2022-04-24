// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: global.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetFeaturesResponse) Reset() {
	*x = GetFeaturesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeaturesResponse) ProtoMessage() {}

func (x *GetFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeaturesResponse.ProtoReflect.Descriptor instead.
func (*GetFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{0}
}

func (x *GetFeaturesResponse) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetFeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *GetFeaturesRequest) Reset() {
	*x = GetFeaturesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeaturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeaturesRequest) ProtoMessage() {}

func (x *GetFeaturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeaturesRequest.ProtoReflect.Descriptor instead.
func (*GetFeaturesRequest) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeaturesRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type CreateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateTenantResponse) Reset() {
	*x = CreateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantResponse) ProtoMessage() {}

func (x *CreateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantResponse.ProtoReflect.Descriptor instead.
func (*CreateTenantResponse) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTenantResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type CreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantGroup     string `protobuf:"bytes,2,opt,name=tenantGroup,proto3" json:"tenantGroup,omitempty"`
	Az              string `protobuf:"bytes,3,opt,name=az,proto3" json:"az,omitempty"`
	Env             string `protobuf:"bytes,4,opt,name=env,proto3" json:"env,omitempty"`
	ProjectId       string `protobuf:"bytes,5,opt,name=projectId,proto3" json:"projectId,omitempty"`
	ProjectName     string `protobuf:"bytes,6,opt,name=projectName,proto3" json:"projectName,omitempty"`
	AdminAddr       string `protobuf:"bytes,7,opt,name=adminAddr,proto3" json:"adminAddr,omitempty"`
	GatewayEndpoint string `protobuf:"bytes,8,opt,name=gatewayEndpoint,proto3" json:"gatewayEndpoint,omitempty"`
	InnerAddr       string `protobuf:"bytes,9,opt,name=innerAddr,proto3" json:"innerAddr,omitempty"`
	ServiceName     string `protobuf:"bytes,10,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	InstanceId      string `protobuf:"bytes,11,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTenantRequest) GetTenantGroup() string {
	if x != nil {
		return x.TenantGroup
	}
	return ""
}

func (x *CreateTenantRequest) GetAz() string {
	if x != nil {
		return x.Az
	}
	return ""
}

func (x *CreateTenantRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *CreateTenantRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateTenantRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *CreateTenantRequest) GetAdminAddr() string {
	if x != nil {
		return x.AdminAddr
	}
	return ""
}

func (x *CreateTenantRequest) GetGatewayEndpoint() string {
	if x != nil {
		return x.GatewayEndpoint
	}
	return ""
}

func (x *CreateTenantRequest) GetInnerAddr() string {
	if x != nil {
		return x.InnerAddr
	}
	return ""
}

func (x *CreateTenantRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *CreateTenantRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type GetTenantGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTenantGroupResponse) Reset() {
	*x = GetTenantGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantGroupResponse) ProtoMessage() {}

func (x *GetTenantGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantGroupResponse.ProtoReflect.Descriptor instead.
func (*GetTenantGroupResponse) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{4}
}

func (x *GetTenantGroupResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetTenantGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Env       string `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *GetTenantGroupRequest) Reset() {
	*x = GetTenantGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantGroupRequest) ProtoMessage() {}

func (x *GetTenantGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantGroupRequest.ProtoReflect.Descriptor instead.
func (*GetTenantGroupRequest) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{5}
}

func (x *GetTenantGroupRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GetTenantGroupRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

type HealthModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthModule) Reset() {
	*x = HealthModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthModule) ProtoMessage() {}

func (x *HealthModule) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthModule.ProtoReflect.Descriptor instead.
func (*HealthModule) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{6}
}

func (x *HealthModule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HealthModule) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HealthModule) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetHealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHealthRequest) Reset() {
	*x = GetHealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthRequest) ProtoMessage() {}

func (x *GetHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthRequest.ProtoReflect.Descriptor instead.
func (*GetHealthRequest) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{7}
}

type GetHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Modules []*HealthModule `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *GetHealthResponse) Reset() {
	*x = GetHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthResponse) ProtoMessage() {}

func (x *GetHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthResponse.ProtoReflect.Descriptor instead.
func (*GetHealthResponse) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{8}
}

func (x *GetHealthResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetHealthResponse) GetModules() []*HealthModule {
	if x != nil {
		return x.Modules
	}
	return nil
}

var File_global_proto protoreflect.FileDescriptor

var file_global_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x98, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xd1, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x7a,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x7a, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x47, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x54, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x32, 0xe4, 0x04, 0x0a, 0x0d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x27, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70,
	0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0xb0, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x2e, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x3b, 0x12, 0x39, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x3d, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x7d, 0x26, 0x65, 0x6e, 0x76, 0x3d, 0x7b, 0x65, 0x6e, 0x76, 0x7d, 0x12, 0x85, 0x01, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2a, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x14,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x70, 0x61, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_global_proto_rawDescOnce sync.Once
	file_global_proto_rawDescData = file_global_proto_rawDesc
)

func file_global_proto_rawDescGZIP() []byte {
	file_global_proto_rawDescOnce.Do(func() {
		file_global_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_proto_rawDescData)
	})
	return file_global_proto_rawDescData
}

var file_global_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_global_proto_goTypes = []interface{}{
	(*GetFeaturesResponse)(nil),    // 0: erda.core.hepa.global.GetFeaturesResponse
	(*GetFeaturesRequest)(nil),     // 1: erda.core.hepa.global.GetFeaturesRequest
	(*CreateTenantResponse)(nil),   // 2: erda.core.hepa.global.CreateTenantResponse
	(*CreateTenantRequest)(nil),    // 3: erda.core.hepa.global.CreateTenantRequest
	(*GetTenantGroupResponse)(nil), // 4: erda.core.hepa.global.GetTenantGroupResponse
	(*GetTenantGroupRequest)(nil),  // 5: erda.core.hepa.global.GetTenantGroupRequest
	(*HealthModule)(nil),           // 6: erda.core.hepa.global.HealthModule
	(*GetHealthRequest)(nil),       // 7: erda.core.hepa.global.GetHealthRequest
	(*GetHealthResponse)(nil),      // 8: erda.core.hepa.global.GetHealthResponse
	nil,                            // 9: erda.core.hepa.global.GetFeaturesResponse.DataEntry
}
var file_global_proto_depIdxs = []int32{
	9, // 0: erda.core.hepa.global.GetFeaturesResponse.data:type_name -> erda.core.hepa.global.GetFeaturesResponse.DataEntry
	6, // 1: erda.core.hepa.global.GetHealthResponse.modules:type_name -> erda.core.hepa.global.HealthModule
	7, // 2: erda.core.hepa.global.GlobalService.GetHealth:input_type -> erda.core.hepa.global.GetHealthRequest
	5, // 3: erda.core.hepa.global.GlobalService.GetTenantGroup:input_type -> erda.core.hepa.global.GetTenantGroupRequest
	3, // 4: erda.core.hepa.global.GlobalService.CreateTenant:input_type -> erda.core.hepa.global.CreateTenantRequest
	1, // 5: erda.core.hepa.global.GlobalService.GetFeatures:input_type -> erda.core.hepa.global.GetFeaturesRequest
	8, // 6: erda.core.hepa.global.GlobalService.GetHealth:output_type -> erda.core.hepa.global.GetHealthResponse
	4, // 7: erda.core.hepa.global.GlobalService.GetTenantGroup:output_type -> erda.core.hepa.global.GetTenantGroupResponse
	2, // 8: erda.core.hepa.global.GlobalService.CreateTenant:output_type -> erda.core.hepa.global.CreateTenantResponse
	0, // 9: erda.core.hepa.global.GlobalService.GetFeatures:output_type -> erda.core.hepa.global.GetFeaturesResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_global_proto_init() }
func file_global_proto_init() {
	if File_global_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_global_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeaturesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeaturesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthModule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_global_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_global_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_global_proto_goTypes,
		DependencyIndexes: file_global_proto_depIdxs,
		MessageInfos:      file_global_proto_msgTypes,
	}.Build()
	File_global_proto = out.File
	file_global_proto_rawDesc = nil
	file_global_proto_goTypes = nil
	file_global_proto_depIdxs = nil
}
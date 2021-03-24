// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: plugin.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PluginNameRequest) Reset() {
	*x = PluginNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginNameRequest) ProtoMessage() {}

func (x *PluginNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginNameRequest.ProtoReflect.Descriptor instead.
func (*PluginNameRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{0}
}

type PluginNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginName string `protobuf:"bytes,1,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
}

func (x *PluginNameResponse) Reset() {
	*x = PluginNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginNameResponse) ProtoMessage() {}

func (x *PluginNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginNameResponse.ProtoReflect.Descriptor instead.
func (*PluginNameResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginNameResponse) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

type PolicyNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PolicyNamesRequest) Reset() {
	*x = PolicyNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyNamesRequest) ProtoMessage() {}

func (x *PolicyNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyNamesRequest.ProtoReflect.Descriptor instead.
func (*PolicyNamesRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{2}
}

type PolicyNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyNames []string `protobuf:"bytes,1,rep,name=policy_names,json=policyNames,proto3" json:"policy_names,omitempty"`
}

func (x *PolicyNamesResponse) Reset() {
	*x = PolicyNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyNamesResponse) ProtoMessage() {}

func (x *PolicyNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyNamesResponse.ProtoReflect.Descriptor instead.
func (*PolicyNamesResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *PolicyNamesResponse) GetPolicyNames() []string {
	if x != nil {
		return x.PolicyNames
	}
	return nil
}

type ConfigurePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginConfig *structpb.Struct `protobuf:"bytes,1,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
}

func (x *ConfigurePluginRequest) Reset() {
	*x = ConfigurePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurePluginRequest) ProtoMessage() {}

func (x *ConfigurePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurePluginRequest.ProtoReflect.Descriptor instead.
func (*ConfigurePluginRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigurePluginRequest) GetPluginConfig() *structpb.Struct {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

type ConfigurePluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigurePluginResponse) Reset() {
	*x = ConfigurePluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurePluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurePluginResponse) ProtoMessage() {}

func (x *ConfigurePluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurePluginResponse.ProtoReflect.Descriptor instead.
func (*ConfigurePluginResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{5}
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyName       string `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	AdmissionRequest []byte `protobuf:"bytes,2,opt,name=admission_request,json=admissionRequest,proto3" json:"admission_request,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateRequest) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *ValidateRequest) GetAdmissionRequest() []byte {
	if x != nil {
		return x.AdmissionRequest
	}
	return nil
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceViolations []*ResourceViolation `protobuf:"bytes,1,rep,name=resource_violations,json=resourceViolations,proto3" json:"resource_violations,omitempty"`
	PatchOperations    []*PatchOperation    `protobuf:"bytes,2,rep,name=patch_operations,json=patchOperations,proto3" json:"patch_operations,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateResponse) GetResourceViolations() []*ResourceViolation {
	if x != nil {
		return x.ResourceViolations
	}
	return nil
}

func (x *ValidateResponse) GetPatchOperations() []*PatchOperation {
	if x != nil {
		return x.PatchOperations
	}
	return nil
}

type ResourceViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceKind string `protobuf:"bytes,2,opt,name=resource_kind,json=resourceKind,proto3" json:"resource_kind,omitempty"`
	Namespace    string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Violation    string `protobuf:"bytes,4,opt,name=violation,proto3" json:"violation,omitempty"`
	Policy       string `protobuf:"bytes,5,opt,name=policy,proto3" json:"policy,omitempty"`
	Error        string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResourceViolation) Reset() {
	*x = ResourceViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceViolation) ProtoMessage() {}

func (x *ResourceViolation) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceViolation.ProtoReflect.Descriptor instead.
func (*ResourceViolation) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{8}
}

func (x *ResourceViolation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ResourceViolation) GetResourceKind() string {
	if x != nil {
		return x.ResourceKind
	}
	return ""
}

func (x *ResourceViolation) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ResourceViolation) GetViolation() string {
	if x != nil {
		return x.Violation
	}
	return ""
}

func (x *ResourceViolation) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *ResourceViolation) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type PatchOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string          `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Path  string          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PatchOperation) Reset() {
	*x = PatchOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchOperation) ProtoMessage() {}

func (x *PatchOperation) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchOperation.ProtoReflect.Descriptor instead.
func (*PatchOperation) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{9}
}

func (x *PatchOperation) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *PatchOperation) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PatchOperation) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_plugin_proto protoreflect.FileDescriptor

var file_plugin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x12, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x14, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x56, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5f, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x10, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x62, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0xa5, 0x02, 0x0a, 0x0b, 0x4b, 0x52, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x12, 0x41, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_proto_rawDescOnce sync.Once
	file_plugin_proto_rawDescData = file_plugin_proto_rawDesc
)

func file_plugin_proto_rawDescGZIP() []byte {
	file_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_proto_rawDescData)
	})
	return file_plugin_proto_rawDescData
}

var file_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_plugin_proto_goTypes = []interface{}{
	(*PluginNameRequest)(nil),       // 0: proto.PluginNameRequest
	(*PluginNameResponse)(nil),      // 1: proto.PluginNameResponse
	(*PolicyNamesRequest)(nil),      // 2: proto.PolicyNamesRequest
	(*PolicyNamesResponse)(nil),     // 3: proto.PolicyNamesResponse
	(*ConfigurePluginRequest)(nil),  // 4: proto.ConfigurePluginRequest
	(*ConfigurePluginResponse)(nil), // 5: proto.ConfigurePluginResponse
	(*ValidateRequest)(nil),         // 6: proto.ValidateRequest
	(*ValidateResponse)(nil),        // 7: proto.ValidateResponse
	(*ResourceViolation)(nil),       // 8: proto.ResourceViolation
	(*PatchOperation)(nil),          // 9: proto.PatchOperation
	(*structpb.Struct)(nil),         // 10: google.protobuf.Struct
	(*structpb.Value)(nil),          // 11: google.protobuf.Value
}
var file_plugin_proto_depIdxs = []int32{
	10, // 0: proto.ConfigurePluginRequest.plugin_config:type_name -> google.protobuf.Struct
	8,  // 1: proto.ValidateResponse.resource_violations:type_name -> proto.ResourceViolation
	9,  // 2: proto.ValidateResponse.patch_operations:type_name -> proto.PatchOperation
	11, // 3: proto.PatchOperation.value:type_name -> google.protobuf.Value
	0,  // 4: proto.KRailPlugin.PluginName:input_type -> proto.PluginNameRequest
	2,  // 5: proto.KRailPlugin.PolicyNames:input_type -> proto.PolicyNamesRequest
	4,  // 6: proto.KRailPlugin.ConfigurePlugin:input_type -> proto.ConfigurePluginRequest
	6,  // 7: proto.KRailPlugin.Validate:input_type -> proto.ValidateRequest
	1,  // 8: proto.KRailPlugin.PluginName:output_type -> proto.PluginNameResponse
	3,  // 9: proto.KRailPlugin.PolicyNames:output_type -> proto.PolicyNamesResponse
	5,  // 10: proto.KRailPlugin.ConfigurePlugin:output_type -> proto.ConfigurePluginResponse
	7,  // 11: proto.KRailPlugin.Validate:output_type -> proto.ValidateResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_plugin_proto_init() }
func file_plugin_proto_init() {
	if File_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginNameRequest); i {
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
		file_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginNameResponse); i {
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
		file_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyNamesRequest); i {
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
		file_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyNamesResponse); i {
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
		file_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurePluginRequest); i {
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
		file_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurePluginResponse); i {
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
		file_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
		file_plugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceViolation); i {
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
		file_plugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchOperation); i {
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
			RawDescriptor: file_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_proto_msgTypes,
	}.Build()
	File_plugin_proto = out.File
	file_plugin_proto_rawDesc = nil
	file_plugin_proto_goTypes = nil
	file_plugin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KRailPluginClient is the client API for KRailPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KRailPluginClient interface {
	PluginName(ctx context.Context, in *PluginNameRequest, opts ...grpc.CallOption) (*PluginNameResponse, error)
	PolicyNames(ctx context.Context, in *PolicyNamesRequest, opts ...grpc.CallOption) (*PolicyNamesResponse, error)
	ConfigurePlugin(ctx context.Context, in *ConfigurePluginRequest, opts ...grpc.CallOption) (*ConfigurePluginResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type kRailPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewKRailPluginClient(cc grpc.ClientConnInterface) KRailPluginClient {
	return &kRailPluginClient{cc}
}

func (c *kRailPluginClient) PluginName(ctx context.Context, in *PluginNameRequest, opts ...grpc.CallOption) (*PluginNameResponse, error) {
	out := new(PluginNameResponse)
	err := c.cc.Invoke(ctx, "/proto.KRailPlugin/PluginName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kRailPluginClient) PolicyNames(ctx context.Context, in *PolicyNamesRequest, opts ...grpc.CallOption) (*PolicyNamesResponse, error) {
	out := new(PolicyNamesResponse)
	err := c.cc.Invoke(ctx, "/proto.KRailPlugin/PolicyNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kRailPluginClient) ConfigurePlugin(ctx context.Context, in *ConfigurePluginRequest, opts ...grpc.CallOption) (*ConfigurePluginResponse, error) {
	out := new(ConfigurePluginResponse)
	err := c.cc.Invoke(ctx, "/proto.KRailPlugin/ConfigurePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kRailPluginClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/proto.KRailPlugin/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KRailPluginServer is the server API for KRailPlugin service.
type KRailPluginServer interface {
	PluginName(context.Context, *PluginNameRequest) (*PluginNameResponse, error)
	PolicyNames(context.Context, *PolicyNamesRequest) (*PolicyNamesResponse, error)
	ConfigurePlugin(context.Context, *ConfigurePluginRequest) (*ConfigurePluginResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

// UnimplementedKRailPluginServer can be embedded to have forward compatible implementations.
type UnimplementedKRailPluginServer struct {
}

func (*UnimplementedKRailPluginServer) PluginName(context.Context, *PluginNameRequest) (*PluginNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginName not implemented")
}
func (*UnimplementedKRailPluginServer) PolicyNames(context.Context, *PolicyNamesRequest) (*PolicyNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyNames not implemented")
}
func (*UnimplementedKRailPluginServer) ConfigurePlugin(context.Context, *ConfigurePluginRequest) (*ConfigurePluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigurePlugin not implemented")
}
func (*UnimplementedKRailPluginServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterKRailPluginServer(s *grpc.Server, srv KRailPluginServer) {
	s.RegisterService(&_KRailPlugin_serviceDesc, srv)
}

func _KRailPlugin_PluginName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KRailPluginServer).PluginName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KRailPlugin/PluginName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KRailPluginServer).PluginName(ctx, req.(*PluginNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KRailPlugin_PolicyNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KRailPluginServer).PolicyNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KRailPlugin/PolicyNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KRailPluginServer).PolicyNames(ctx, req.(*PolicyNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KRailPlugin_ConfigurePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KRailPluginServer).ConfigurePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KRailPlugin/ConfigurePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KRailPluginServer).ConfigurePlugin(ctx, req.(*ConfigurePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KRailPlugin_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KRailPluginServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KRailPlugin/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KRailPluginServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KRailPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KRailPlugin",
	HandlerType: (*KRailPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginName",
			Handler:    _KRailPlugin_PluginName_Handler,
		},
		{
			MethodName: "PolicyNames",
			Handler:    _KRailPlugin_PolicyNames_Handler,
		},
		{
			MethodName: "ConfigurePlugin",
			Handler:    _KRailPlugin_ConfigurePlugin_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _KRailPlugin_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: synapsis/v1/service.proto

package synapsisv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synapsis_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synapsis_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_synapsis_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_synapsis_v1_service_proto protoreflect.FileDescriptor

var file_synapsis_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x9f, 0x08, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1f, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x73,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x59, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x22, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x29, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73,
	0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26,
	0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x28, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb1, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x79,
	0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x61, 0x66, 0x69, 0x64, 0x70, 0x72, 0x61, 0x79,
	0x6f, 0x67, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0b, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x17, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_synapsis_v1_service_proto_rawDescOnce sync.Once
	file_synapsis_v1_service_proto_rawDescData = file_synapsis_v1_service_proto_rawDesc
)

func file_synapsis_v1_service_proto_rawDescGZIP() []byte {
	file_synapsis_v1_service_proto_rawDescOnce.Do(func() {
		file_synapsis_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_synapsis_v1_service_proto_rawDescData)
	})
	return file_synapsis_v1_service_proto_rawDescData
}

var file_synapsis_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synapsis_v1_service_proto_goTypes = []interface{}{
	(*PingResponse)(nil),                  // 0: synapsis.v1.PingResponse
	(*emptypb.Empty)(nil),                 // 1: google.protobuf.Empty
	(*CreateUserRequest)(nil),             // 2: synapsis.v1.CreateUserRequest
	(*GetUserByIdRequest)(nil),            // 3: synapsis.v1.GetUserByIdRequest
	(*LoginRequest)(nil),                  // 4: synapsis.v1.LoginRequest
	(*LogoutRequest)(nil),                 // 5: synapsis.v1.LogoutRequest
	(*ForgotPasswordRequest)(nil),         // 6: synapsis.v1.ForgotPasswordRequest
	(*ResetPasswordRequest)(nil),          // 7: synapsis.v1.ResetPasswordRequest
	(*CreateProductCategoryRequest)(nil),  // 8: synapsis.v1.CreateProductCategoryRequest
	(*GetProductCategoryRequest)(nil),     // 9: synapsis.v1.GetProductCategoryRequest
	(*UpdateProductCategoryRequest)(nil),  // 10: synapsis.v1.UpdateProductCategoryRequest
	(*DeleteProductCategoryRequest)(nil),  // 11: synapsis.v1.DeleteProductCategoryRequest
	(*GetProductCategoriesRequest)(nil),   // 12: synapsis.v1.GetProductCategoriesRequest
	(*CreateUserResponse)(nil),            // 13: synapsis.v1.CreateUserResponse
	(*GetUserByIdResponse)(nil),           // 14: synapsis.v1.GetUserByIdResponse
	(*LoginResponse)(nil),                 // 15: synapsis.v1.LoginResponse
	(*ForgotPasswordResponse)(nil),        // 16: synapsis.v1.ForgotPasswordResponse
	(*CreateProductCategoryResponse)(nil), // 17: synapsis.v1.CreateProductCategoryResponse
	(*GetProductCategoryResponse)(nil),    // 18: synapsis.v1.GetProductCategoryResponse
	(*GetProductCategoriesResponse)(nil),  // 19: synapsis.v1.GetProductCategoriesResponse
}
var file_synapsis_v1_service_proto_depIdxs = []int32{
	1,  // 0: synapsis.v1.SynapsisService.Ping:input_type -> google.protobuf.Empty
	2,  // 1: synapsis.v1.SynapsisService.CreateUser:input_type -> synapsis.v1.CreateUserRequest
	3,  // 2: synapsis.v1.SynapsisService.GetUserById:input_type -> synapsis.v1.GetUserByIdRequest
	4,  // 3: synapsis.v1.SynapsisService.Login:input_type -> synapsis.v1.LoginRequest
	5,  // 4: synapsis.v1.SynapsisService.Logout:input_type -> synapsis.v1.LogoutRequest
	6,  // 5: synapsis.v1.SynapsisService.ForgotPassword:input_type -> synapsis.v1.ForgotPasswordRequest
	7,  // 6: synapsis.v1.SynapsisService.ResetPassword:input_type -> synapsis.v1.ResetPasswordRequest
	8,  // 7: synapsis.v1.SynapsisService.CreateProductCategory:input_type -> synapsis.v1.CreateProductCategoryRequest
	9,  // 8: synapsis.v1.SynapsisService.GetProductCategory:input_type -> synapsis.v1.GetProductCategoryRequest
	10, // 9: synapsis.v1.SynapsisService.UpdateProductCategory:input_type -> synapsis.v1.UpdateProductCategoryRequest
	11, // 10: synapsis.v1.SynapsisService.DeleteProductCategory:input_type -> synapsis.v1.DeleteProductCategoryRequest
	12, // 11: synapsis.v1.SynapsisService.GetProductCategories:input_type -> synapsis.v1.GetProductCategoriesRequest
	0,  // 12: synapsis.v1.SynapsisService.Ping:output_type -> synapsis.v1.PingResponse
	13, // 13: synapsis.v1.SynapsisService.CreateUser:output_type -> synapsis.v1.CreateUserResponse
	14, // 14: synapsis.v1.SynapsisService.GetUserById:output_type -> synapsis.v1.GetUserByIdResponse
	15, // 15: synapsis.v1.SynapsisService.Login:output_type -> synapsis.v1.LoginResponse
	1,  // 16: synapsis.v1.SynapsisService.Logout:output_type -> google.protobuf.Empty
	16, // 17: synapsis.v1.SynapsisService.ForgotPassword:output_type -> synapsis.v1.ForgotPasswordResponse
	1,  // 18: synapsis.v1.SynapsisService.ResetPassword:output_type -> google.protobuf.Empty
	17, // 19: synapsis.v1.SynapsisService.CreateProductCategory:output_type -> synapsis.v1.CreateProductCategoryResponse
	18, // 20: synapsis.v1.SynapsisService.GetProductCategory:output_type -> synapsis.v1.GetProductCategoryResponse
	18, // 21: synapsis.v1.SynapsisService.UpdateProductCategory:output_type -> synapsis.v1.GetProductCategoryResponse
	1,  // 22: synapsis.v1.SynapsisService.DeleteProductCategory:output_type -> google.protobuf.Empty
	19, // 23: synapsis.v1.SynapsisService.GetProductCategories:output_type -> synapsis.v1.GetProductCategoriesResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_synapsis_v1_service_proto_init() }
func file_synapsis_v1_service_proto_init() {
	if File_synapsis_v1_service_proto != nil {
		return
	}
	file_synapsis_v1_user_proto_init()
	file_synapsis_v1_product_category_proto_init()
	file_synapsis_v1_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synapsis_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_synapsis_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_synapsis_v1_service_proto_goTypes,
		DependencyIndexes: file_synapsis_v1_service_proto_depIdxs,
		MessageInfos:      file_synapsis_v1_service_proto_msgTypes,
	}.Build()
	File_synapsis_v1_service_proto = out.File
	file_synapsis_v1_service_proto_rawDesc = nil
	file_synapsis_v1_service_proto_goTypes = nil
	file_synapsis_v1_service_proto_depIdxs = nil
}

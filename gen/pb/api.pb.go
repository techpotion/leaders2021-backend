// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: api.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x14, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x67, 0x65, 0x6f, 0x6a, 0x73, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc8, 0x07, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x92, 0x41,
	0x10, 0x0a, 0x0e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x20, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x84, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x92, 0x41, 0x10, 0x0a, 0x0e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x20, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0xb3, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12,
	0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x92, 0x41, 0x19, 0x0a, 0x17, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x20, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x20, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x72,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x92, 0x41, 0x0a, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f,
	0x6e, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f,
	0x6e, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x12,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6f,
	0x4a, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x47, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x79, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x92, 0x41, 0x0a, 0x0a, 0x08, 0x47, 0x65, 0x6f,
	0x4a, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x72, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x4a,
	0x73, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x47, 0x65,
	0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x92, 0x41, 0x0a, 0x0a,
	0x08, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x10, 0x50, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x42, 0xf4, 0x04, 0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0xe8, 0x04, 0x12,
	0x61, 0x0a, 0x14, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x20, 0x32, 0x30, 0x32, 0x31, 0x20,
	0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0b, 0x54, 0x65, 0x63, 0x68, 0x20,
	0x50, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x70,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x76, 0x61, 0x72, 0x6c, 0x61, 0x6d, 0x6f, 0x77, 0x2e,
	0x63, 0x6f, 0x6c, 0x40, 0x79, 0x61, 0x68, 0x6f, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x1a, 0x12, 0x38, 0x39, 0x2e, 0x31, 0x37, 0x38, 0x2e, 0x32, 0x33, 0x39, 0x2e, 0x38,
	0x34, 0x3a, 0x33, 0x32, 0x30, 0x31, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x67,
	0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x60, 0x0a, 0x46, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x62, 0x6f, 0x64, 0x79, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x50, 0x49, 0x27, 0x73,
	0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x71, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x6a,
	0x0a, 0x50, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x72, 0x69, 0x65, 0x73, 0x20, 0x74,
	0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x67, 0x0a, 0x03, 0x34, 0x30,
	0x33, 0x12, 0x60, 0x0a, 0x46, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x14, 0x1a,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x4a, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x43, 0x0a, 0x29, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x36, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x2f, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_goTypes = []interface{}{
	(*SportsObjects_ListRequest)(nil),          // 0: api.SportsObjects.ListRequest
	(*SportsObjects_GetRequest)(nil),           // 1: api.SportsObjects.GetRequest
	(*SportsObjectsDetailed_ListRequest)(nil),  // 2: api.SportsObjectsDetailed.ListRequest
	(*GeoJsons_Request)(nil),                   // 3: api.GeoJsons.Request
	(*PolygonAnalytics_Request)(nil),           // 4: api.PolygonAnalytics.Request
	(*SportsObjects_ListResponse)(nil),         // 5: api.SportsObjects.ListResponse
	(*SportsObjects_GetResponse)(nil),          // 6: api.SportsObjects.GetResponse
	(*SportsObjectsDetailed_ListResponse)(nil), // 7: api.SportsObjectsDetailed.ListResponse
	(*GeoJsons_Response)(nil),                  // 8: api.GeoJsons.Response
	(*PolygonAnalytics_Response)(nil),          // 9: api.PolygonAnalytics.Response
}
var file_api_proto_depIdxs = []int32{
	0, // 0: api.ApiService.ListSportsObjects:input_type -> api.SportsObjects.ListRequest
	1, // 1: api.ApiService.GetSportsObject:input_type -> api.SportsObjects.GetRequest
	2, // 2: api.ApiService.ListSportsObjectsDetailed:input_type -> api.SportsObjectsDetailed.ListRequest
	3, // 3: api.ApiService.GetGeoJsonRegions:input_type -> api.GeoJsons.Request
	3, // 4: api.ApiService.GetGeoJsonDensityHeatmap:input_type -> api.GeoJsons.Request
	3, // 5: api.ApiService.GetGeoJsonObjects:input_type -> api.GeoJsons.Request
	4, // 6: api.ApiService.PolygonAnalytics:input_type -> api.PolygonAnalytics.Request
	5, // 7: api.ApiService.ListSportsObjects:output_type -> api.SportsObjects.ListResponse
	6, // 8: api.ApiService.GetSportsObject:output_type -> api.SportsObjects.GetResponse
	7, // 9: api.ApiService.ListSportsObjectsDetailed:output_type -> api.SportsObjectsDetailed.ListResponse
	8, // 10: api.ApiService.GetGeoJsonRegions:output_type -> api.GeoJsons.Response
	8, // 11: api.ApiService.GetGeoJsonDensityHeatmap:output_type -> api.GeoJsons.Response
	8, // 12: api.ApiService.GetGeoJsonObjects:output_type -> api.GeoJsons.Response
	9, // 13: api.ApiService.PolygonAnalytics:output_type -> api.PolygonAnalytics.Response
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_sports_objects_proto_init()
	file_geojsons_proto_init()
	file_analytics_proto_init()
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: common.proto

package pb

import (
	options1 "github.com/TheSDTM/protoc-gen-gorm/options"
	validate "github.com/envoyproxy/protoc-gen-validate/validate"
	options "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	annotations "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/annotations.proto.

var E_Http = annotations.E_Http

// Symbols defined in public import of github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2/options/annotations.proto.

var E_Openapiv2Swagger = options.E_Openapiv2Swagger
var E_Openapiv2Operation = options.E_Openapiv2Operation
var E_Openapiv2Schema = options.E_Openapiv2Schema
var E_Openapiv2Tag = options.E_Openapiv2Tag
var E_Openapiv2Field = options.E_Openapiv2Field

// Symbols defined in public import of github.com/TheSDTM/protoc-gen-gorm/options/gorm.proto.

type FieldWritePermission = options1.FieldWritePermission

const FieldWritePermission_FieldWritePermissionUpdateOnly = options1.FieldWritePermission_FieldWritePermissionUpdateOnly
const FieldWritePermission_FieldWritePermissionCreateOnly = options1.FieldWritePermission_FieldWritePermissionCreateOnly
const FieldWritePermission_FieldWritePermissionNo = options1.FieldWritePermission_FieldWritePermissionNo
const FieldWritePermission_FieldWritePermissionCreateUpdate = options1.FieldWritePermission_FieldWritePermissionCreateUpdate

var FieldWritePermission_name = options1.FieldWritePermission_name
var FieldWritePermission_value = options1.FieldWritePermission_value

type GormFileOptions = options1.GormFileOptions
type GormMessageOptions = options1.GormMessageOptions
type ExtraField = options1.ExtraField
type GormFieldOptions = options1.GormFieldOptions
type GormFieldOptions_HasOne = options1.GormFieldOptions_HasOne
type GormFieldOptions_BelongsTo = options1.GormFieldOptions_BelongsTo
type GormFieldOptions_HasMany = options1.GormFieldOptions_HasMany
type GormFieldOptions_ManyToMany = options1.GormFieldOptions_ManyToMany
type GormTag = options1.GormTag
type HasOneOptions = options1.HasOneOptions
type BelongsToOptions = options1.BelongsToOptions
type HasManyOptions = options1.HasManyOptions
type ManyToManyOptions = options1.ManyToManyOptions

var E_FileOpts = options1.E_FileOpts
var E_Opts = options1.E_Opts
var E_Field = options1.E_Field

// Symbols defined in public import of github.com/envoyproxy/protoc-gen-validate/validate/validate.proto.

type KnownRegex = validate.KnownRegex

const KnownRegex_UNKNOWN = validate.KnownRegex_UNKNOWN
const KnownRegex_HTTP_HEADER_NAME = validate.KnownRegex_HTTP_HEADER_NAME
const KnownRegex_HTTP_HEADER_VALUE = validate.KnownRegex_HTTP_HEADER_VALUE

var KnownRegex_name = validate.KnownRegex_name
var KnownRegex_value = validate.KnownRegex_value

type FieldRules = validate.FieldRules
type FieldRules_Float = validate.FieldRules_Float
type FieldRules_Double = validate.FieldRules_Double
type FieldRules_Int32 = validate.FieldRules_Int32
type FieldRules_Int64 = validate.FieldRules_Int64
type FieldRules_Uint32 = validate.FieldRules_Uint32
type FieldRules_Uint64 = validate.FieldRules_Uint64
type FieldRules_Sint32 = validate.FieldRules_Sint32
type FieldRules_Sint64 = validate.FieldRules_Sint64
type FieldRules_Fixed32 = validate.FieldRules_Fixed32
type FieldRules_Fixed64 = validate.FieldRules_Fixed64
type FieldRules_Sfixed32 = validate.FieldRules_Sfixed32
type FieldRules_Sfixed64 = validate.FieldRules_Sfixed64
type FieldRules_Bool = validate.FieldRules_Bool
type FieldRules_String_ = validate.FieldRules_String_
type FieldRules_Bytes = validate.FieldRules_Bytes
type FieldRules_Enum = validate.FieldRules_Enum
type FieldRules_Repeated = validate.FieldRules_Repeated
type FieldRules_Map = validate.FieldRules_Map
type FieldRules_Any = validate.FieldRules_Any
type FieldRules_Duration = validate.FieldRules_Duration
type FieldRules_Timestamp = validate.FieldRules_Timestamp
type FloatRules = validate.FloatRules
type DoubleRules = validate.DoubleRules
type Int32Rules = validate.Int32Rules
type Int64Rules = validate.Int64Rules
type UInt32Rules = validate.UInt32Rules
type UInt64Rules = validate.UInt64Rules
type SInt32Rules = validate.SInt32Rules
type SInt64Rules = validate.SInt64Rules
type Fixed32Rules = validate.Fixed32Rules
type Fixed64Rules = validate.Fixed64Rules
type SFixed32Rules = validate.SFixed32Rules
type SFixed64Rules = validate.SFixed64Rules
type BoolRules = validate.BoolRules
type StringRules = validate.StringRules

const Default_StringRules_Strict = validate.Default_StringRules_Strict

type StringRules_Email = validate.StringRules_Email
type StringRules_Hostname = validate.StringRules_Hostname
type StringRules_Ip = validate.StringRules_Ip
type StringRules_Ipv4 = validate.StringRules_Ipv4
type StringRules_Ipv6 = validate.StringRules_Ipv6
type StringRules_Uri = validate.StringRules_Uri
type StringRules_UriRef = validate.StringRules_UriRef
type StringRules_Address = validate.StringRules_Address
type StringRules_Uuid = validate.StringRules_Uuid
type StringRules_WellKnownRegex = validate.StringRules_WellKnownRegex
type BytesRules = validate.BytesRules
type BytesRules_Ip = validate.BytesRules_Ip
type BytesRules_Ipv4 = validate.BytesRules_Ipv4
type BytesRules_Ipv6 = validate.BytesRules_Ipv6
type EnumRules = validate.EnumRules
type MessageRules = validate.MessageRules
type RepeatedRules = validate.RepeatedRules
type MapRules = validate.MapRules
type AnyRules = validate.AnyRules
type DurationRules = validate.DurationRules
type TimestampRules = validate.TimestampRules

var E_Disabled = validate.E_Disabled
var E_Ignored = validate.E_Ignored
var E_Required = validate.E_Required
var E_Rules = validate.E_Rules

// Symbols defined in public import of google/protobuf/any.proto.

type Any = anypb.Any

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber     uint32 `protobuf:"varint,1,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	ResultsPerPage uint32 `protobuf:"varint,2,opt,name=resultsPerPage,proto3" json:"resultsPerPage,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Pagination) GetResultsPerPage() uint32 {
	if x != nil {
		return x.ResultsPerPage
	}
	return 0
}

type ListStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListStats) Reset() {
	*x = ListStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStats) ProtoMessage() {}

func (x *ListStats) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStats.ProtoReflect.Descriptor instead.
func (*ListStats) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *ListStats) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float32 `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Point) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Point) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type SportsObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId                     uint32 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	ObjectName                   string `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	ObjectPoint                  *Point `protobuf:"bytes,3,opt,name=object_point,json=objectPoint,proto3" json:"object_point,omitempty"`
	DepartmentalOrganizationId   uint32 `protobuf:"varint,4,opt,name=departmental_organization_id,json=departmentalOrganizationId,proto3" json:"departmental_organization_id,omitempty"`
	DepartmentalOrganizationName string `protobuf:"bytes,5,opt,name=departmental_organization_name,json=departmentalOrganizationName,proto3" json:"departmental_organization_name,omitempty"`
	Availability                 uint32 `protobuf:"varint,9,opt,name=availability,proto3" json:"availability,omitempty"`
	SportKind                    string `protobuf:"bytes,10,opt,name=sport_kind,json=sportKind,proto3" json:"sport_kind,omitempty"`
}

func (x *SportsObject) Reset() {
	*x = SportsObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObject) ProtoMessage() {}

func (x *SportsObject) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObject.ProtoReflect.Descriptor instead.
func (*SportsObject) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *SportsObject) GetObjectId() uint32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *SportsObject) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *SportsObject) GetObjectPoint() *Point {
	if x != nil {
		return x.ObjectPoint
	}
	return nil
}

func (x *SportsObject) GetDepartmentalOrganizationId() uint32 {
	if x != nil {
		return x.DepartmentalOrganizationId
	}
	return 0
}

func (x *SportsObject) GetDepartmentalOrganizationName() string {
	if x != nil {
		return x.DepartmentalOrganizationName
	}
	return ""
}

func (x *SportsObject) GetAvailability() uint32 {
	if x != nil {
		return x.Availability
	}
	return 0
}

func (x *SportsObject) GetSportKind() string {
	if x != nil {
		return x.SportKind
	}
	return ""
}

type SportsObjectDetailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId                     uint32 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	ObjectName                   string `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	ObjectPoint                  *Point `protobuf:"bytes,3,opt,name=object_point,json=objectPoint,proto3" json:"object_point,omitempty"`
	DepartmentalOrganizationId   uint32 `protobuf:"varint,4,opt,name=departmental_organization_id,json=departmentalOrganizationId,proto3" json:"departmental_organization_id,omitempty"`
	DepartmentalOrganizationName string `protobuf:"bytes,5,opt,name=departmental_organization_name,json=departmentalOrganizationName,proto3" json:"departmental_organization_name,omitempty"`
	SportsAreaId                 uint32 `protobuf:"varint,6,opt,name=sports_area_id,json=sportsAreaId,proto3" json:"sports_area_id,omitempty"`
	SportsAreaName               string `protobuf:"bytes,7,opt,name=sports_area_name,json=sportsAreaName,proto3" json:"sports_area_name,omitempty"`
	SportsAreaType               string `protobuf:"bytes,8,opt,name=sports_area_type,json=sportsAreaType,proto3" json:"sports_area_type,omitempty"`
	Availability                 uint32 `protobuf:"varint,9,opt,name=availability,proto3" json:"availability,omitempty"`
	SportKind                    string `protobuf:"bytes,10,opt,name=sport_kind,json=sportKind,proto3" json:"sport_kind,omitempty"`
}

func (x *SportsObjectDetailed) Reset() {
	*x = SportsObjectDetailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjectDetailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjectDetailed) ProtoMessage() {}

func (x *SportsObjectDetailed) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjectDetailed.ProtoReflect.Descriptor instead.
func (*SportsObjectDetailed) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *SportsObjectDetailed) GetObjectId() uint32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *SportsObjectDetailed) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *SportsObjectDetailed) GetObjectPoint() *Point {
	if x != nil {
		return x.ObjectPoint
	}
	return nil
}

func (x *SportsObjectDetailed) GetDepartmentalOrganizationId() uint32 {
	if x != nil {
		return x.DepartmentalOrganizationId
	}
	return 0
}

func (x *SportsObjectDetailed) GetDepartmentalOrganizationName() string {
	if x != nil {
		return x.DepartmentalOrganizationName
	}
	return ""
}

func (x *SportsObjectDetailed) GetSportsAreaId() uint32 {
	if x != nil {
		return x.SportsAreaId
	}
	return 0
}

func (x *SportsObjectDetailed) GetSportsAreaName() string {
	if x != nil {
		return x.SportsAreaName
	}
	return ""
}

func (x *SportsObjectDetailed) GetSportsAreaType() string {
	if x != nil {
		return x.SportsAreaType
	}
	return ""
}

func (x *SportsObjectDetailed) GetAvailability() uint32 {
	if x != nil {
		return x.Availability
	}
	return 0
}

func (x *SportsObjectDetailed) GetSportKind() string {
	if x != nil {
		return x.SportKind
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x1a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x53, 0x44, 0x54, 0x4d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x21,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x33, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x3a, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0xf6, 0x02, 0x0a, 0x0c, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x15, 0x0a, 0x13, 0x40,
	0x01, 0x60, 0x01, 0x6a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x40, 0x0a, 0x1c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x44, 0x0a, 0x1e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x3a, 0x13, 0xba, 0xb9, 0x19, 0x02,
	0x08, 0x01, 0xba, 0xb9, 0x19, 0x09, 0x1a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x81, 0x04, 0x0a, 0x14, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x15, 0x0a, 0x13,
	0x40, 0x01, 0x60, 0x01, 0x6a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x40, 0x0a, 0x1c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x44, 0x0a, 0x1e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x3a, 0x1c, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0xba, 0xb9, 0x19,
	0x12, 0x1a, 0x10, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x50, 0x00, 0x50,
	0x01, 0x50, 0x02, 0x50, 0x03, 0x50, 0x04, 0x50, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_proto_goTypes = []interface{}{
	(*Pagination)(nil),           // 0: api.Pagination
	(*ListStats)(nil),            // 1: api.ListStats
	(*Point)(nil),                // 2: api.Point
	(*SportsObject)(nil),         // 3: api.SportsObject
	(*SportsObjectDetailed)(nil), // 4: api.SportsObjectDetailed
}
var file_common_proto_depIdxs = []int32{
	2, // 0: api.SportsObject.object_point:type_name -> api.Point
	2, // 1: api.SportsObjectDetailed.object_point:type_name -> api.Point
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStats); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObject); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjectDetailed); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}

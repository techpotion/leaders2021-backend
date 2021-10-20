// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: sports_objects.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SportsObjects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SportsObjects) Reset() {
	*x = SportsObjects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjects) ProtoMessage() {}

func (x *SportsObjects) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjects.ProtoReflect.Descriptor instead.
func (*SportsObjects) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{0}
}

type SportsObjectsDetailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SportsObjectsDetailed) Reset() {
	*x = SportsObjectsDetailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjectsDetailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjectsDetailed) ProtoMessage() {}

func (x *SportsObjectsDetailed) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjectsDetailed.ProtoReflect.Descriptor instead.
func (*SportsObjectsDetailed) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{1}
}

type SportsObjects_ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// filters
	ObjectNames                   []string       `protobuf:"bytes,2,rep,name=objectNames,proto3" json:"objectNames,omitempty"`                                         // optional
	DepartmentalOrganizationIds   []uint32       `protobuf:"varint,3,rep,packed,name=departmentalOrganizationIds,proto3" json:"departmentalOrganizationIds,omitempty"` // optional
	DepartmentalOrganizationNames []string       `protobuf:"bytes,4,rep,name=departmentalOrganizationNames,proto3" json:"departmentalOrganizationNames,omitempty"`     // optional
	Availabilities                []Availability `protobuf:"varint,7,rep,packed,name=availabilities,proto3,enum=api.Availability" json:"availabilities,omitempty"`     // optional
}

func (x *SportsObjects_ListRequest) Reset() {
	*x = SportsObjects_ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjects_ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjects_ListRequest) ProtoMessage() {}

func (x *SportsObjects_ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjects_ListRequest.ProtoReflect.Descriptor instead.
func (*SportsObjects_ListRequest) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SportsObjects_ListRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *SportsObjects_ListRequest) GetObjectNames() []string {
	if x != nil {
		return x.ObjectNames
	}
	return nil
}

func (x *SportsObjects_ListRequest) GetDepartmentalOrganizationIds() []uint32 {
	if x != nil {
		return x.DepartmentalOrganizationIds
	}
	return nil
}

func (x *SportsObjects_ListRequest) GetDepartmentalOrganizationNames() []string {
	if x != nil {
		return x.DepartmentalOrganizationNames
	}
	return nil
}

func (x *SportsObjects_ListRequest) GetAvailabilities() []Availability {
	if x != nil {
		return x.Availabilities
	}
	return nil
}

type SportsObjects_ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportsObjects []*SportsObject `protobuf:"bytes,1,rep,name=sportsObjects,proto3" json:"sportsObjects,omitempty"`
	ListStats     *ListStats      `protobuf:"bytes,2,opt,name=listStats,proto3" json:"listStats,omitempty"`
}

func (x *SportsObjects_ListResponse) Reset() {
	*x = SportsObjects_ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjects_ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjects_ListResponse) ProtoMessage() {}

func (x *SportsObjects_ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjects_ListResponse.ProtoReflect.Descriptor instead.
func (*SportsObjects_ListResponse) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SportsObjects_ListResponse) GetSportsObjects() []*SportsObject {
	if x != nil {
		return x.SportsObjects
	}
	return nil
}

func (x *SportsObjects_ListResponse) GetListStats() *ListStats {
	if x != nil {
		return x.ListStats
	}
	return nil
}

type SportsObjects_GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId uint32 `protobuf:"varint,2,opt,name=objectId,proto3" json:"objectId,omitempty"`
}

func (x *SportsObjects_GetRequest) Reset() {
	*x = SportsObjects_GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjects_GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjects_GetRequest) ProtoMessage() {}

func (x *SportsObjects_GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjects_GetRequest.ProtoReflect.Descriptor instead.
func (*SportsObjects_GetRequest) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SportsObjects_GetRequest) GetObjectId() uint32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

type SportsObjects_GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportsObject *SportsObject `protobuf:"bytes,1,opt,name=sportsObject,proto3" json:"sportsObject,omitempty"`
}

func (x *SportsObjects_GetResponse) Reset() {
	*x = SportsObjects_GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjects_GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjects_GetResponse) ProtoMessage() {}

func (x *SportsObjects_GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjects_GetResponse.ProtoReflect.Descriptor instead.
func (*SportsObjects_GetResponse) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{0, 3}
}

func (x *SportsObjects_GetResponse) GetSportsObject() *SportsObject {
	if x != nil {
		return x.SportsObject
	}
	return nil
}

type SportsObjectsDetailed_ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// filters
	ObjectIds                     []uint32       `protobuf:"varint,2,rep,packed,name=objectIds,proto3" json:"objectIds,omitempty"`                                     // optional
	ObjectNames                   []string       `protobuf:"bytes,3,rep,name=objectNames,proto3" json:"objectNames,omitempty"`                                         // optional
	DepartmentalOrganizationIds   []uint32       `protobuf:"varint,4,rep,packed,name=departmentalOrganizationIds,proto3" json:"departmentalOrganizationIds,omitempty"` // optional
	DepartmentalOrganizationNames []string       `protobuf:"bytes,5,rep,name=departmentalOrganizationNames,proto3" json:"departmentalOrganizationNames,omitempty"`     // optional
	SportsAreaNames               []string       `protobuf:"bytes,6,rep,name=sportsAreaNames,proto3" json:"sportsAreaNames,omitempty"`                                 // optional
	SportsAreaTypes               []string       `protobuf:"bytes,7,rep,name=sportsAreaTypes,proto3" json:"sportsAreaTypes,omitempty"`                                 // optional
	Availabilities                []Availability `protobuf:"varint,8,rep,packed,name=availabilities,proto3,enum=api.Availability" json:"availabilities,omitempty"`     // optional
	SportKinds                    []string       `protobuf:"bytes,9,rep,name=sportKinds,proto3" json:"sportKinds,omitempty"`                                           // optional
}

func (x *SportsObjectsDetailed_ListRequest) Reset() {
	*x = SportsObjectsDetailed_ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjectsDetailed_ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjectsDetailed_ListRequest) ProtoMessage() {}

func (x *SportsObjectsDetailed_ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjectsDetailed_ListRequest.ProtoReflect.Descriptor instead.
func (*SportsObjectsDetailed_ListRequest) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SportsObjectsDetailed_ListRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetObjectIds() []uint32 {
	if x != nil {
		return x.ObjectIds
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetObjectNames() []string {
	if x != nil {
		return x.ObjectNames
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetDepartmentalOrganizationIds() []uint32 {
	if x != nil {
		return x.DepartmentalOrganizationIds
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetDepartmentalOrganizationNames() []string {
	if x != nil {
		return x.DepartmentalOrganizationNames
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetSportsAreaNames() []string {
	if x != nil {
		return x.SportsAreaNames
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetSportsAreaTypes() []string {
	if x != nil {
		return x.SportsAreaTypes
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetAvailabilities() []Availability {
	if x != nil {
		return x.Availabilities
	}
	return nil
}

func (x *SportsObjectsDetailed_ListRequest) GetSportKinds() []string {
	if x != nil {
		return x.SportKinds
	}
	return nil
}

type SportsObjectsDetailed_ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportsObjects []*SportsObjectDetailed `protobuf:"bytes,1,rep,name=sportsObjects,proto3" json:"sportsObjects,omitempty"`
	ListStats     *ListStats              `protobuf:"bytes,2,opt,name=listStats,proto3" json:"listStats,omitempty"`
}

func (x *SportsObjectsDetailed_ListResponse) Reset() {
	*x = SportsObjectsDetailed_ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_objects_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsObjectsDetailed_ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsObjectsDetailed_ListResponse) ProtoMessage() {}

func (x *SportsObjectsDetailed_ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_objects_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsObjectsDetailed_ListResponse.ProtoReflect.Descriptor instead.
func (*SportsObjectsDetailed_ListResponse) Descriptor() ([]byte, []int) {
	return file_sports_objects_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SportsObjectsDetailed_ListResponse) GetSportsObjects() []*SportsObjectDetailed {
	if x != nil {
		return x.SportsObjects
	}
	return nil
}

func (x *SportsObjectsDetailed_ListResponse) GetListStats() *ListStats {
	if x != nil {
		return x.ListStats
	}
	return nil
}

var File_sports_objects_proto protoreflect.FileDescriptor

var file_sports_objects_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x04, 0x0a, 0x0d, 0x53, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0xc4, 0x02, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x0b,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x01,
	0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x52, 0x0a,
	0x1b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0d, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x2a, 0x06, 0x28, 0xa0,
	0x8d, 0x06, 0x40, 0x01, 0x52, 0x1b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x12, 0x44, 0x0a, 0x1d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x1a, 0x75, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0d, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x6c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09,
	0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04,
	0x28, 0xa0, 0x8d, 0x06, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x44,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0xae, 0x05, 0x0a, 0x15, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x1a, 0x95,
	0x04, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0d, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x2a, 0x06, 0x28, 0xa0,
	0x8d, 0x06, 0x40, 0x01, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12,
	0x2f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03,
	0xd0, 0x01, 0x01, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x52, 0x0a, 0x1b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x2a,
	0x06, 0x28, 0xc0, 0x9a, 0x0c, 0x40, 0x01, 0x52, 0x1b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x1d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0f, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0xd0,
	0x01, 0x01, 0x52, 0x0f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42,
	0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x01, 0x52, 0x0f, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x4b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a,
	0x92, 0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x01, 0x52, 0x0a, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x1a, 0x7d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x0d, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sports_objects_proto_rawDescOnce sync.Once
	file_sports_objects_proto_rawDescData = file_sports_objects_proto_rawDesc
)

func file_sports_objects_proto_rawDescGZIP() []byte {
	file_sports_objects_proto_rawDescOnce.Do(func() {
		file_sports_objects_proto_rawDescData = protoimpl.X.CompressGZIP(file_sports_objects_proto_rawDescData)
	})
	return file_sports_objects_proto_rawDescData
}

var file_sports_objects_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sports_objects_proto_goTypes = []interface{}{
	(*SportsObjects)(nil),                      // 0: api.SportsObjects
	(*SportsObjectsDetailed)(nil),              // 1: api.SportsObjectsDetailed
	(*SportsObjects_ListRequest)(nil),          // 2: api.SportsObjects.ListRequest
	(*SportsObjects_ListResponse)(nil),         // 3: api.SportsObjects.ListResponse
	(*SportsObjects_GetRequest)(nil),           // 4: api.SportsObjects.GetRequest
	(*SportsObjects_GetResponse)(nil),          // 5: api.SportsObjects.GetResponse
	(*SportsObjectsDetailed_ListRequest)(nil),  // 6: api.SportsObjectsDetailed.ListRequest
	(*SportsObjectsDetailed_ListResponse)(nil), // 7: api.SportsObjectsDetailed.ListResponse
	(*Pagination)(nil),                         // 8: api.Pagination
	(Availability)(0),                          // 9: api.Availability
	(*SportsObject)(nil),                       // 10: api.SportsObject
	(*ListStats)(nil),                          // 11: api.ListStats
	(*SportsObjectDetailed)(nil),               // 12: api.SportsObjectDetailed
}
var file_sports_objects_proto_depIdxs = []int32{
	8,  // 0: api.SportsObjects.ListRequest.pagination:type_name -> api.Pagination
	9,  // 1: api.SportsObjects.ListRequest.availabilities:type_name -> api.Availability
	10, // 2: api.SportsObjects.ListResponse.sportsObjects:type_name -> api.SportsObject
	11, // 3: api.SportsObjects.ListResponse.listStats:type_name -> api.ListStats
	10, // 4: api.SportsObjects.GetResponse.sportsObject:type_name -> api.SportsObject
	8,  // 5: api.SportsObjectsDetailed.ListRequest.pagination:type_name -> api.Pagination
	9,  // 6: api.SportsObjectsDetailed.ListRequest.availabilities:type_name -> api.Availability
	12, // 7: api.SportsObjectsDetailed.ListResponse.sportsObjects:type_name -> api.SportsObjectDetailed
	11, // 8: api.SportsObjectsDetailed.ListResponse.listStats:type_name -> api.ListStats
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_sports_objects_proto_init() }
func file_sports_objects_proto_init() {
	if File_sports_objects_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sports_objects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjects); i {
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
		file_sports_objects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjectsDetailed); i {
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
		file_sports_objects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjects_ListRequest); i {
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
		file_sports_objects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjects_ListResponse); i {
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
		file_sports_objects_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjects_GetRequest); i {
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
		file_sports_objects_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjects_GetResponse); i {
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
		file_sports_objects_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjectsDetailed_ListRequest); i {
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
		file_sports_objects_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsObjectsDetailed_ListResponse); i {
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
			RawDescriptor: file_sports_objects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sports_objects_proto_goTypes,
		DependencyIndexes: file_sports_objects_proto_depIdxs,
		MessageInfos:      file_sports_objects_proto_msgTypes,
	}.Build()
	File_sports_objects_proto = out.File
	file_sports_objects_proto_rawDesc = nil
	file_sports_objects_proto_goTypes = nil
	file_sports_objects_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: analytics.proto

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

type PolygonAnalytics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PolygonAnalytics) Reset() {
	*x = PolygonAnalytics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonAnalytics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonAnalytics) ProtoMessage() {}

func (x *PolygonAnalytics) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonAnalytics.ProtoReflect.Descriptor instead.
func (*PolygonAnalytics) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0}
}

type PolygonParkAnalytics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PolygonParkAnalytics) Reset() {
	*x = PolygonParkAnalytics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonParkAnalytics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonParkAnalytics) ProtoMessage() {}

func (x *PolygonParkAnalytics) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonParkAnalytics.ProtoReflect.Descriptor instead.
func (*PolygonParkAnalytics) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1}
}

type PolygonAnalytics_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygon *Polygon `protobuf:"bytes,1,opt,name=polygon,proto3" json:"polygon,omitempty"`
	// filters
	SportsAreaType string       `protobuf:"bytes,2,opt,name=sportsAreaType,proto3" json:"sportsAreaType,omitempty"`                    // optional
	Availability   Availability `protobuf:"varint,3,opt,name=availability,proto3,enum=api.Availability" json:"availability,omitempty"` // optional
	SportKind      string       `protobuf:"bytes,4,opt,name=sportKind,proto3" json:"sportKind,omitempty"`                              // optional
}

func (x *PolygonAnalytics_Request) Reset() {
	*x = PolygonAnalytics_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonAnalytics_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonAnalytics_Request) ProtoMessage() {}

func (x *PolygonAnalytics_Request) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonAnalytics_Request.ProtoReflect.Descriptor instead.
func (*PolygonAnalytics_Request) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PolygonAnalytics_Request) GetPolygon() *Polygon {
	if x != nil {
		return x.Polygon
	}
	return nil
}

func (x *PolygonAnalytics_Request) GetSportsAreaType() string {
	if x != nil {
		return x.SportsAreaType
	}
	return ""
}

func (x *PolygonAnalytics_Request) GetAvailability() Availability {
	if x != nil {
		return x.Availability
	}
	return Availability_UNKNOWN
}

func (x *PolygonAnalytics_Request) GetSportKind() string {
	if x != nil {
		return x.SportKind
	}
	return ""
}

type PolygonAnalytics_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreasSquare     float64  `protobuf:"fixed64,1,opt,name=areasSquare,proto3" json:"areasSquare,omitempty"`
	AreasAmount     uint32   `protobuf:"varint,2,opt,name=areasAmount,proto3" json:"areasAmount,omitempty"`
	SportsAmount    uint32   `protobuf:"varint,3,opt,name=sportsAmount,proto3" json:"sportsAmount,omitempty"`
	SportsKinds     []string `protobuf:"bytes,4,rep,name=sportsKinds,proto3" json:"sportsKinds,omitempty"`
	AreaTypes       []string `protobuf:"bytes,5,rep,name=areaTypes,proto3" json:"areaTypes,omitempty"`
	AreaTypesAmount uint32   `protobuf:"varint,6,opt,name=areaTypesAmount,proto3" json:"areaTypesAmount,omitempty"`
}

func (x *PolygonAnalytics_Response) Reset() {
	*x = PolygonAnalytics_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonAnalytics_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonAnalytics_Response) ProtoMessage() {}

func (x *PolygonAnalytics_Response) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonAnalytics_Response.ProtoReflect.Descriptor instead.
func (*PolygonAnalytics_Response) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PolygonAnalytics_Response) GetAreasSquare() float64 {
	if x != nil {
		return x.AreasSquare
	}
	return 0
}

func (x *PolygonAnalytics_Response) GetAreasAmount() uint32 {
	if x != nil {
		return x.AreasAmount
	}
	return 0
}

func (x *PolygonAnalytics_Response) GetSportsAmount() uint32 {
	if x != nil {
		return x.SportsAmount
	}
	return 0
}

func (x *PolygonAnalytics_Response) GetSportsKinds() []string {
	if x != nil {
		return x.SportsKinds
	}
	return nil
}

func (x *PolygonAnalytics_Response) GetAreaTypes() []string {
	if x != nil {
		return x.AreaTypes
	}
	return nil
}

func (x *PolygonAnalytics_Response) GetAreaTypesAmount() uint32 {
	if x != nil {
		return x.AreaTypesAmount
	}
	return 0
}

type PolygonParkAnalytics_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygon *Polygon `protobuf:"bytes,1,opt,name=polygon,proto3" json:"polygon,omitempty"`
	// filters
	HasSportground bool `protobuf:"varint,2,opt,name=hasSportground,proto3" json:"hasSportground,omitempty"`
}

func (x *PolygonParkAnalytics_Request) Reset() {
	*x = PolygonParkAnalytics_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonParkAnalytics_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonParkAnalytics_Request) ProtoMessage() {}

func (x *PolygonParkAnalytics_Request) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonParkAnalytics_Request.ProtoReflect.Descriptor instead.
func (*PolygonParkAnalytics_Request) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PolygonParkAnalytics_Request) GetPolygon() *Polygon {
	if x != nil {
		return x.Polygon
	}
	return nil
}

func (x *PolygonParkAnalytics_Request) GetHasSportground() bool {
	if x != nil {
		return x.HasSportground
	}
	return false
}

type PolygonParkAnalytics_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parks     []*Park    `protobuf:"bytes,1,rep,name=parks,proto3" json:"parks,omitempty"`
	ListStats *ListStats `protobuf:"bytes,2,opt,name=listStats,proto3" json:"listStats,omitempty"`
}

func (x *PolygonParkAnalytics_Response) Reset() {
	*x = PolygonParkAnalytics_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonParkAnalytics_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonParkAnalytics_Response) ProtoMessage() {}

func (x *PolygonParkAnalytics_Response) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonParkAnalytics_Response.ProtoReflect.Descriptor instead.
func (*PolygonParkAnalytics_Response) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PolygonParkAnalytics_Response) GetParks() []*Park {
	if x != nil {
		return x.Parks
	}
	return nil
}

func (x *PolygonParkAnalytics_Response) GetListStats() *ListStats {
	if x != nil {
		return x.ListStats
	}
	return nil
}

var File_analytics_proto protoreflect.FileDescriptor

var file_analytics_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a, 0x10, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0xc2, 0x01, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c,
	0x79, 0x67, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x0e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x01, 0x52,
	0x0e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x35, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4b,
	0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0xd0, 0x01, 0x01, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xdc,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x72, 0x65, 0x61, 0x73, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x73, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x72, 0x65, 0x61, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4b, 0x69, 0x6e,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x65, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61, 0x72,
	0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xcc, 0x01,
	0x0a, 0x14, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x59, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x1a, 0x59, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x52, 0x05, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x2c,
	0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x08, 0x5a, 0x06,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_proto_rawDescOnce sync.Once
	file_analytics_proto_rawDescData = file_analytics_proto_rawDesc
)

func file_analytics_proto_rawDescGZIP() []byte {
	file_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_proto_rawDescData)
	})
	return file_analytics_proto_rawDescData
}

var file_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_analytics_proto_goTypes = []interface{}{
	(*PolygonAnalytics)(nil),              // 0: api.PolygonAnalytics
	(*PolygonParkAnalytics)(nil),          // 1: api.PolygonParkAnalytics
	(*PolygonAnalytics_Request)(nil),      // 2: api.PolygonAnalytics.Request
	(*PolygonAnalytics_Response)(nil),     // 3: api.PolygonAnalytics.Response
	(*PolygonParkAnalytics_Request)(nil),  // 4: api.PolygonParkAnalytics.Request
	(*PolygonParkAnalytics_Response)(nil), // 5: api.PolygonParkAnalytics.Response
	(*Polygon)(nil),                       // 6: api.Polygon
	(Availability)(0),                     // 7: api.Availability
	(*Park)(nil),                          // 8: api.Park
	(*ListStats)(nil),                     // 9: api.ListStats
}
var file_analytics_proto_depIdxs = []int32{
	6, // 0: api.PolygonAnalytics.Request.polygon:type_name -> api.Polygon
	7, // 1: api.PolygonAnalytics.Request.availability:type_name -> api.Availability
	6, // 2: api.PolygonParkAnalytics.Request.polygon:type_name -> api.Polygon
	8, // 3: api.PolygonParkAnalytics.Response.parks:type_name -> api.Park
	9, // 4: api.PolygonParkAnalytics.Response.listStats:type_name -> api.ListStats
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_analytics_proto_init() }
func file_analytics_proto_init() {
	if File_analytics_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonAnalytics); i {
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
		file_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonParkAnalytics); i {
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
		file_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonAnalytics_Request); i {
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
		file_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonAnalytics_Response); i {
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
		file_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonParkAnalytics_Request); i {
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
		file_analytics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonParkAnalytics_Response); i {
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
			RawDescriptor: file_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_proto_depIdxs,
		MessageInfos:      file_analytics_proto_msgTypes,
	}.Build()
	File_analytics_proto = out.File
	file_analytics_proto_rawDesc = nil
	file_analytics_proto_goTypes = nil
	file_analytics_proto_depIdxs = nil
}

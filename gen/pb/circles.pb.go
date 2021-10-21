// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: circles.proto

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

type Intersections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Intersections) Reset() {
	*x = Intersections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intersections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intersections) ProtoMessage() {}

func (x *Intersections) ProtoReflect() protoreflect.Message {
	mi := &file_circles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intersections.ProtoReflect.Descriptor instead.
func (*Intersections) Descriptor() ([]byte, []int) {
	return file_circles_proto_rawDescGZIP(), []int{0}
}

type Intersections_ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygon *Polygon `protobuf:"bytes,1,opt,name=polygon,proto3" json:"polygon,omitempty"`
	// filters
	Availability Availability `protobuf:"varint,2,opt,name=availability,proto3,enum=api.Availability" json:"availability,omitempty"` // optional
}

func (x *Intersections_ListRequest) Reset() {
	*x = Intersections_ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intersections_ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intersections_ListRequest) ProtoMessage() {}

func (x *Intersections_ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_circles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intersections_ListRequest.ProtoReflect.Descriptor instead.
func (*Intersections_ListRequest) Descriptor() ([]byte, []int) {
	return file_circles_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Intersections_ListRequest) GetPolygon() *Polygon {
	if x != nil {
		return x.Polygon
	}
	return nil
}

func (x *Intersections_ListRequest) GetAvailability() Availability {
	if x != nil {
		return x.Availability
	}
	return Availability_UNKNOWN
}

type Intersections_ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intersections []*CircleIntersections `protobuf:"bytes,1,rep,name=intersections,proto3" json:"intersections,omitempty"`
	ListStats     *ListStats             `protobuf:"bytes,2,opt,name=listStats,proto3" json:"listStats,omitempty"`
}

func (x *Intersections_ListResponse) Reset() {
	*x = Intersections_ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intersections_ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intersections_ListResponse) ProtoMessage() {}

func (x *Intersections_ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_circles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intersections_ListResponse.ProtoReflect.Descriptor instead.
func (*Intersections_ListResponse) Descriptor() ([]byte, []int) {
	return file_circles_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Intersections_ListResponse) GetIntersections() []*CircleIntersections {
	if x != nil {
		return x.Intersections
	}
	return nil
}

func (x *Intersections_ListResponse) GetListStats() *ListStats {
	if x != nil {
		return x.ListStats
	}
	return nil
}

var File_circles_proto protoreflect.FileDescriptor

var file_circles_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x6c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0c, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x1a, 0x7c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x42, 0x08, 0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_circles_proto_rawDescOnce sync.Once
	file_circles_proto_rawDescData = file_circles_proto_rawDesc
)

func file_circles_proto_rawDescGZIP() []byte {
	file_circles_proto_rawDescOnce.Do(func() {
		file_circles_proto_rawDescData = protoimpl.X.CompressGZIP(file_circles_proto_rawDescData)
	})
	return file_circles_proto_rawDescData
}

var file_circles_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_circles_proto_goTypes = []interface{}{
	(*Intersections)(nil),              // 0: api.Intersections
	(*Intersections_ListRequest)(nil),  // 1: api.Intersections.ListRequest
	(*Intersections_ListResponse)(nil), // 2: api.Intersections.ListResponse
	(*Polygon)(nil),                    // 3: api.Polygon
	(Availability)(0),                  // 4: api.Availability
	(*CircleIntersections)(nil),        // 5: api.CircleIntersections
	(*ListStats)(nil),                  // 6: api.ListStats
}
var file_circles_proto_depIdxs = []int32{
	3, // 0: api.Intersections.ListRequest.polygon:type_name -> api.Polygon
	4, // 1: api.Intersections.ListRequest.availability:type_name -> api.Availability
	5, // 2: api.Intersections.ListResponse.intersections:type_name -> api.CircleIntersections
	6, // 3: api.Intersections.ListResponse.listStats:type_name -> api.ListStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_circles_proto_init() }
func file_circles_proto_init() {
	if File_circles_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_circles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intersections); i {
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
		file_circles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intersections_ListRequest); i {
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
		file_circles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intersections_ListResponse); i {
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
			RawDescriptor: file_circles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_circles_proto_goTypes,
		DependencyIndexes: file_circles_proto_depIdxs,
		MessageInfos:      file_circles_proto_msgTypes,
	}.Build()
	File_circles_proto = out.File
	file_circles_proto_rawDesc = nil
	file_circles_proto_goTypes = nil
	file_circles_proto_depIdxs = nil
}
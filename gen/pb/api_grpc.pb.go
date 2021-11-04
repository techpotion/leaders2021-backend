// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	// Sports Objects
	ListSportsObjects(ctx context.Context, in *SportsObjects_ListRequest, opts ...grpc.CallOption) (*SportsObjects_ListResponse, error)
	ListSportsObjectsFromDetailed(ctx context.Context, in *SportsObjectsDetailed_ListRequest, opts ...grpc.CallOption) (*SportsObjects_ListResponse, error)
	GetSportsObject(ctx context.Context, in *SportsObjects_GetRequest, opts ...grpc.CallOption) (*SportsObjects_GetResponse, error)
	// Sports Objects Detailed
	ListSportsObjectsDetailed(ctx context.Context, in *SportsObjectsDetailed_ListRequest, opts ...grpc.CallOption) (*SportsObjectsDetailed_ListResponse, error)
	// GeoJSONs
	GetGeoJsonRegions(ctx context.Context, in *GeoJsons_Request, opts ...grpc.CallOption) (*GeoJsons_Response, error)
	GetGeoJsonDensityHeatmap(ctx context.Context, in *GeoJsons_Request, opts ...grpc.CallOption) (*GeoJsons_Response, error)
	GetGeoJsonSportsObjects(ctx context.Context, in *GeoJsons_Request, opts ...grpc.CallOption) (*GeoJsons_Response, error)
	// Analytics
	PolygonAnalytics(ctx context.Context, in *PolygonAnalytics_Request, opts ...grpc.CallOption) (*PolygonAnalytics_Response, error)
	PolygonParkAnalytics(ctx context.Context, in *PolygonParkAnalytics_Request, opts ...grpc.CallOption) (*PolygonParkAnalytics_Response, error)
	PolygonPollutionAnalytics(ctx context.Context, in *PolygonPollutionAnalytics_Request, opts ...grpc.CallOption) (*PolygonPollutionAnalytics_Response, error)
	PolygonSubwayAnalytics(ctx context.Context, in *PolygonSubwayAnalytics_Request, opts ...grpc.CallOption) (*PolygonSubwayAnalytics_Response, error)
	PolygonAnalyticsDashboard(ctx context.Context, in *PolygonAnalyticsDashboard_Request, opts ...grpc.CallOption) (*PolygonAnalyticsDashboard_Response, error)
	// Marks
	GetMark(ctx context.Context, in *Marks_GetRequest, opts ...grpc.CallOption) (*Marks_GetResponse, error)
	// Filters
	ListObjectsNames(ctx context.Context, in *ObjectsNames_ListRequest, opts ...grpc.CallOption) (*ObjectsNames_ListResponse, error)
	ListDepartmentalOrganizationsIds(ctx context.Context, in *DepartmentalOrganizationsIds_ListRequest, opts ...grpc.CallOption) (*DepartmentalOrganizationsIds_ListResponse, error)
	ListDepartmentalOrganizationsNames(ctx context.Context, in *DepartmentalOrganizationsNames_ListRequest, opts ...grpc.CallOption) (*DepartmentalOrganizationsNames_ListResponse, error)
	ListSportsAreaNames(ctx context.Context, in *SportsAreaNames_ListRequest, opts ...grpc.CallOption) (*SportsAreaNames_ListResponse, error)
	ListSportsAreaTypes(ctx context.Context, in *SportsAreaTypes_ListRequest, opts ...grpc.CallOption) (*SportsAreaTypes_ListResponse, error)
	ListSportKinds(ctx context.Context, in *SportKinds_ListRequest, opts ...grpc.CallOption) (*SportKinds_ListResponse, error)
	// Circles
	ListIntersections(ctx context.Context, in *Intersections_ListRequest, opts ...grpc.CallOption) (*Intersections_ListResponse, error)
	ListUnions(ctx context.Context, in *Unions_ListRequest, opts ...grpc.CallOption) (*Unions_ListResponse, error)
	// Densities
	GetDensity(ctx context.Context, in *Densities_GetRequest, opts ...grpc.CallOption) (*Densities_GetResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) ListSportsObjects(ctx context.Context, in *SportsObjects_ListRequest, opts ...grpc.CallOption) (*SportsObjects_ListResponse, error) {
	out := new(SportsObjects_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListSportsObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListSportsObjectsFromDetailed(ctx context.Context, in *SportsObjectsDetailed_ListRequest, opts ...grpc.CallOption) (*SportsObjects_ListResponse, error) {
	out := new(SportsObjects_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListSportsObjectsFromDetailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetSportsObject(ctx context.Context, in *SportsObjects_GetRequest, opts ...grpc.CallOption) (*SportsObjects_GetResponse, error) {
	out := new(SportsObjects_GetResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/GetSportsObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListSportsObjectsDetailed(ctx context.Context, in *SportsObjectsDetailed_ListRequest, opts ...grpc.CallOption) (*SportsObjectsDetailed_ListResponse, error) {
	out := new(SportsObjectsDetailed_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListSportsObjectsDetailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetGeoJsonRegions(ctx context.Context, in *GeoJsons_Request, opts ...grpc.CallOption) (*GeoJsons_Response, error) {
	out := new(GeoJsons_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/GetGeoJsonRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetGeoJsonDensityHeatmap(ctx context.Context, in *GeoJsons_Request, opts ...grpc.CallOption) (*GeoJsons_Response, error) {
	out := new(GeoJsons_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/GetGeoJsonDensityHeatmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetGeoJsonSportsObjects(ctx context.Context, in *GeoJsons_Request, opts ...grpc.CallOption) (*GeoJsons_Response, error) {
	out := new(GeoJsons_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/GetGeoJsonSportsObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) PolygonAnalytics(ctx context.Context, in *PolygonAnalytics_Request, opts ...grpc.CallOption) (*PolygonAnalytics_Response, error) {
	out := new(PolygonAnalytics_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/PolygonAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) PolygonParkAnalytics(ctx context.Context, in *PolygonParkAnalytics_Request, opts ...grpc.CallOption) (*PolygonParkAnalytics_Response, error) {
	out := new(PolygonParkAnalytics_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/PolygonParkAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) PolygonPollutionAnalytics(ctx context.Context, in *PolygonPollutionAnalytics_Request, opts ...grpc.CallOption) (*PolygonPollutionAnalytics_Response, error) {
	out := new(PolygonPollutionAnalytics_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/PolygonPollutionAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) PolygonSubwayAnalytics(ctx context.Context, in *PolygonSubwayAnalytics_Request, opts ...grpc.CallOption) (*PolygonSubwayAnalytics_Response, error) {
	out := new(PolygonSubwayAnalytics_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/PolygonSubwayAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) PolygonAnalyticsDashboard(ctx context.Context, in *PolygonAnalyticsDashboard_Request, opts ...grpc.CallOption) (*PolygonAnalyticsDashboard_Response, error) {
	out := new(PolygonAnalyticsDashboard_Response)
	err := c.cc.Invoke(ctx, "/api.ApiService/PolygonAnalyticsDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMark(ctx context.Context, in *Marks_GetRequest, opts ...grpc.CallOption) (*Marks_GetResponse, error) {
	out := new(Marks_GetResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/GetMark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListObjectsNames(ctx context.Context, in *ObjectsNames_ListRequest, opts ...grpc.CallOption) (*ObjectsNames_ListResponse, error) {
	out := new(ObjectsNames_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListObjectsNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListDepartmentalOrganizationsIds(ctx context.Context, in *DepartmentalOrganizationsIds_ListRequest, opts ...grpc.CallOption) (*DepartmentalOrganizationsIds_ListResponse, error) {
	out := new(DepartmentalOrganizationsIds_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListDepartmentalOrganizationsIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListDepartmentalOrganizationsNames(ctx context.Context, in *DepartmentalOrganizationsNames_ListRequest, opts ...grpc.CallOption) (*DepartmentalOrganizationsNames_ListResponse, error) {
	out := new(DepartmentalOrganizationsNames_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListDepartmentalOrganizationsNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListSportsAreaNames(ctx context.Context, in *SportsAreaNames_ListRequest, opts ...grpc.CallOption) (*SportsAreaNames_ListResponse, error) {
	out := new(SportsAreaNames_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListSportsAreaNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListSportsAreaTypes(ctx context.Context, in *SportsAreaTypes_ListRequest, opts ...grpc.CallOption) (*SportsAreaTypes_ListResponse, error) {
	out := new(SportsAreaTypes_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListSportsAreaTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListSportKinds(ctx context.Context, in *SportKinds_ListRequest, opts ...grpc.CallOption) (*SportKinds_ListResponse, error) {
	out := new(SportKinds_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListSportKinds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListIntersections(ctx context.Context, in *Intersections_ListRequest, opts ...grpc.CallOption) (*Intersections_ListResponse, error) {
	out := new(Intersections_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListIntersections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListUnions(ctx context.Context, in *Unions_ListRequest, opts ...grpc.CallOption) (*Unions_ListResponse, error) {
	out := new(Unions_ListResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/ListUnions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetDensity(ctx context.Context, in *Densities_GetRequest, opts ...grpc.CallOption) (*Densities_GetResponse, error) {
	out := new(Densities_GetResponse)
	err := c.cc.Invoke(ctx, "/api.ApiService/GetDensity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	// Sports Objects
	ListSportsObjects(context.Context, *SportsObjects_ListRequest) (*SportsObjects_ListResponse, error)
	ListSportsObjectsFromDetailed(context.Context, *SportsObjectsDetailed_ListRequest) (*SportsObjects_ListResponse, error)
	GetSportsObject(context.Context, *SportsObjects_GetRequest) (*SportsObjects_GetResponse, error)
	// Sports Objects Detailed
	ListSportsObjectsDetailed(context.Context, *SportsObjectsDetailed_ListRequest) (*SportsObjectsDetailed_ListResponse, error)
	// GeoJSONs
	GetGeoJsonRegions(context.Context, *GeoJsons_Request) (*GeoJsons_Response, error)
	GetGeoJsonDensityHeatmap(context.Context, *GeoJsons_Request) (*GeoJsons_Response, error)
	GetGeoJsonSportsObjects(context.Context, *GeoJsons_Request) (*GeoJsons_Response, error)
	// Analytics
	PolygonAnalytics(context.Context, *PolygonAnalytics_Request) (*PolygonAnalytics_Response, error)
	PolygonParkAnalytics(context.Context, *PolygonParkAnalytics_Request) (*PolygonParkAnalytics_Response, error)
	PolygonPollutionAnalytics(context.Context, *PolygonPollutionAnalytics_Request) (*PolygonPollutionAnalytics_Response, error)
	PolygonSubwayAnalytics(context.Context, *PolygonSubwayAnalytics_Request) (*PolygonSubwayAnalytics_Response, error)
	PolygonAnalyticsDashboard(context.Context, *PolygonAnalyticsDashboard_Request) (*PolygonAnalyticsDashboard_Response, error)
	// Marks
	GetMark(context.Context, *Marks_GetRequest) (*Marks_GetResponse, error)
	// Filters
	ListObjectsNames(context.Context, *ObjectsNames_ListRequest) (*ObjectsNames_ListResponse, error)
	ListDepartmentalOrganizationsIds(context.Context, *DepartmentalOrganizationsIds_ListRequest) (*DepartmentalOrganizationsIds_ListResponse, error)
	ListDepartmentalOrganizationsNames(context.Context, *DepartmentalOrganizationsNames_ListRequest) (*DepartmentalOrganizationsNames_ListResponse, error)
	ListSportsAreaNames(context.Context, *SportsAreaNames_ListRequest) (*SportsAreaNames_ListResponse, error)
	ListSportsAreaTypes(context.Context, *SportsAreaTypes_ListRequest) (*SportsAreaTypes_ListResponse, error)
	ListSportKinds(context.Context, *SportKinds_ListRequest) (*SportKinds_ListResponse, error)
	// Circles
	ListIntersections(context.Context, *Intersections_ListRequest) (*Intersections_ListResponse, error)
	ListUnions(context.Context, *Unions_ListRequest) (*Unions_ListResponse, error)
	// Densities
	GetDensity(context.Context, *Densities_GetRequest) (*Densities_GetResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) ListSportsObjects(context.Context, *SportsObjects_ListRequest) (*SportsObjects_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportsObjects not implemented")
}
func (UnimplementedApiServiceServer) ListSportsObjectsFromDetailed(context.Context, *SportsObjectsDetailed_ListRequest) (*SportsObjects_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportsObjectsFromDetailed not implemented")
}
func (UnimplementedApiServiceServer) GetSportsObject(context.Context, *SportsObjects_GetRequest) (*SportsObjects_GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSportsObject not implemented")
}
func (UnimplementedApiServiceServer) ListSportsObjectsDetailed(context.Context, *SportsObjectsDetailed_ListRequest) (*SportsObjectsDetailed_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportsObjectsDetailed not implemented")
}
func (UnimplementedApiServiceServer) GetGeoJsonRegions(context.Context, *GeoJsons_Request) (*GeoJsons_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoJsonRegions not implemented")
}
func (UnimplementedApiServiceServer) GetGeoJsonDensityHeatmap(context.Context, *GeoJsons_Request) (*GeoJsons_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoJsonDensityHeatmap not implemented")
}
func (UnimplementedApiServiceServer) GetGeoJsonSportsObjects(context.Context, *GeoJsons_Request) (*GeoJsons_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoJsonSportsObjects not implemented")
}
func (UnimplementedApiServiceServer) PolygonAnalytics(context.Context, *PolygonAnalytics_Request) (*PolygonAnalytics_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolygonAnalytics not implemented")
}
func (UnimplementedApiServiceServer) PolygonParkAnalytics(context.Context, *PolygonParkAnalytics_Request) (*PolygonParkAnalytics_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolygonParkAnalytics not implemented")
}
func (UnimplementedApiServiceServer) PolygonPollutionAnalytics(context.Context, *PolygonPollutionAnalytics_Request) (*PolygonPollutionAnalytics_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolygonPollutionAnalytics not implemented")
}
func (UnimplementedApiServiceServer) PolygonSubwayAnalytics(context.Context, *PolygonSubwayAnalytics_Request) (*PolygonSubwayAnalytics_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolygonSubwayAnalytics not implemented")
}
func (UnimplementedApiServiceServer) PolygonAnalyticsDashboard(context.Context, *PolygonAnalyticsDashboard_Request) (*PolygonAnalyticsDashboard_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolygonAnalyticsDashboard not implemented")
}
func (UnimplementedApiServiceServer) GetMark(context.Context, *Marks_GetRequest) (*Marks_GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMark not implemented")
}
func (UnimplementedApiServiceServer) ListObjectsNames(context.Context, *ObjectsNames_ListRequest) (*ObjectsNames_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjectsNames not implemented")
}
func (UnimplementedApiServiceServer) ListDepartmentalOrganizationsIds(context.Context, *DepartmentalOrganizationsIds_ListRequest) (*DepartmentalOrganizationsIds_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartmentalOrganizationsIds not implemented")
}
func (UnimplementedApiServiceServer) ListDepartmentalOrganizationsNames(context.Context, *DepartmentalOrganizationsNames_ListRequest) (*DepartmentalOrganizationsNames_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartmentalOrganizationsNames not implemented")
}
func (UnimplementedApiServiceServer) ListSportsAreaNames(context.Context, *SportsAreaNames_ListRequest) (*SportsAreaNames_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportsAreaNames not implemented")
}
func (UnimplementedApiServiceServer) ListSportsAreaTypes(context.Context, *SportsAreaTypes_ListRequest) (*SportsAreaTypes_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportsAreaTypes not implemented")
}
func (UnimplementedApiServiceServer) ListSportKinds(context.Context, *SportKinds_ListRequest) (*SportKinds_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportKinds not implemented")
}
func (UnimplementedApiServiceServer) ListIntersections(context.Context, *Intersections_ListRequest) (*Intersections_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIntersections not implemented")
}
func (UnimplementedApiServiceServer) ListUnions(context.Context, *Unions_ListRequest) (*Unions_ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUnions not implemented")
}
func (UnimplementedApiServiceServer) GetDensity(context.Context, *Densities_GetRequest) (*Densities_GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDensity not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_ListSportsObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportsObjects_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSportsObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListSportsObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSportsObjects(ctx, req.(*SportsObjects_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListSportsObjectsFromDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportsObjectsDetailed_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSportsObjectsFromDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListSportsObjectsFromDetailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSportsObjectsFromDetailed(ctx, req.(*SportsObjectsDetailed_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetSportsObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportsObjects_GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetSportsObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetSportsObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetSportsObject(ctx, req.(*SportsObjects_GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListSportsObjectsDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportsObjectsDetailed_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSportsObjectsDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListSportsObjectsDetailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSportsObjectsDetailed(ctx, req.(*SportsObjectsDetailed_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetGeoJsonRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoJsons_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetGeoJsonRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetGeoJsonRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetGeoJsonRegions(ctx, req.(*GeoJsons_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetGeoJsonDensityHeatmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoJsons_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetGeoJsonDensityHeatmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetGeoJsonDensityHeatmap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetGeoJsonDensityHeatmap(ctx, req.(*GeoJsons_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetGeoJsonSportsObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoJsons_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetGeoJsonSportsObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetGeoJsonSportsObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetGeoJsonSportsObjects(ctx, req.(*GeoJsons_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_PolygonAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolygonAnalytics_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).PolygonAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/PolygonAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).PolygonAnalytics(ctx, req.(*PolygonAnalytics_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_PolygonParkAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolygonParkAnalytics_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).PolygonParkAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/PolygonParkAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).PolygonParkAnalytics(ctx, req.(*PolygonParkAnalytics_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_PolygonPollutionAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolygonPollutionAnalytics_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).PolygonPollutionAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/PolygonPollutionAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).PolygonPollutionAnalytics(ctx, req.(*PolygonPollutionAnalytics_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_PolygonSubwayAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolygonSubwayAnalytics_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).PolygonSubwayAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/PolygonSubwayAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).PolygonSubwayAnalytics(ctx, req.(*PolygonSubwayAnalytics_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_PolygonAnalyticsDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolygonAnalyticsDashboard_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).PolygonAnalyticsDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/PolygonAnalyticsDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).PolygonAnalyticsDashboard(ctx, req.(*PolygonAnalyticsDashboard_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Marks_GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetMark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMark(ctx, req.(*Marks_GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListObjectsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectsNames_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListObjectsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListObjectsNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListObjectsNames(ctx, req.(*ObjectsNames_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListDepartmentalOrganizationsIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepartmentalOrganizationsIds_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListDepartmentalOrganizationsIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListDepartmentalOrganizationsIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListDepartmentalOrganizationsIds(ctx, req.(*DepartmentalOrganizationsIds_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListDepartmentalOrganizationsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepartmentalOrganizationsNames_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListDepartmentalOrganizationsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListDepartmentalOrganizationsNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListDepartmentalOrganizationsNames(ctx, req.(*DepartmentalOrganizationsNames_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListSportsAreaNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportsAreaNames_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSportsAreaNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListSportsAreaNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSportsAreaNames(ctx, req.(*SportsAreaNames_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListSportsAreaTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportsAreaTypes_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSportsAreaTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListSportsAreaTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSportsAreaTypes(ctx, req.(*SportsAreaTypes_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListSportKinds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportKinds_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSportKinds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListSportKinds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSportKinds(ctx, req.(*SportKinds_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListIntersections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Intersections_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListIntersections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListIntersections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListIntersections(ctx, req.(*Intersections_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListUnions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Unions_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListUnions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/ListUnions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListUnions(ctx, req.(*Unions_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetDensity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Densities_GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetDensity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetDensity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetDensity(ctx, req.(*Densities_GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSportsObjects",
			Handler:    _ApiService_ListSportsObjects_Handler,
		},
		{
			MethodName: "ListSportsObjectsFromDetailed",
			Handler:    _ApiService_ListSportsObjectsFromDetailed_Handler,
		},
		{
			MethodName: "GetSportsObject",
			Handler:    _ApiService_GetSportsObject_Handler,
		},
		{
			MethodName: "ListSportsObjectsDetailed",
			Handler:    _ApiService_ListSportsObjectsDetailed_Handler,
		},
		{
			MethodName: "GetGeoJsonRegions",
			Handler:    _ApiService_GetGeoJsonRegions_Handler,
		},
		{
			MethodName: "GetGeoJsonDensityHeatmap",
			Handler:    _ApiService_GetGeoJsonDensityHeatmap_Handler,
		},
		{
			MethodName: "GetGeoJsonSportsObjects",
			Handler:    _ApiService_GetGeoJsonSportsObjects_Handler,
		},
		{
			MethodName: "PolygonAnalytics",
			Handler:    _ApiService_PolygonAnalytics_Handler,
		},
		{
			MethodName: "PolygonParkAnalytics",
			Handler:    _ApiService_PolygonParkAnalytics_Handler,
		},
		{
			MethodName: "PolygonPollutionAnalytics",
			Handler:    _ApiService_PolygonPollutionAnalytics_Handler,
		},
		{
			MethodName: "PolygonSubwayAnalytics",
			Handler:    _ApiService_PolygonSubwayAnalytics_Handler,
		},
		{
			MethodName: "PolygonAnalyticsDashboard",
			Handler:    _ApiService_PolygonAnalyticsDashboard_Handler,
		},
		{
			MethodName: "GetMark",
			Handler:    _ApiService_GetMark_Handler,
		},
		{
			MethodName: "ListObjectsNames",
			Handler:    _ApiService_ListObjectsNames_Handler,
		},
		{
			MethodName: "ListDepartmentalOrganizationsIds",
			Handler:    _ApiService_ListDepartmentalOrganizationsIds_Handler,
		},
		{
			MethodName: "ListDepartmentalOrganizationsNames",
			Handler:    _ApiService_ListDepartmentalOrganizationsNames_Handler,
		},
		{
			MethodName: "ListSportsAreaNames",
			Handler:    _ApiService_ListSportsAreaNames_Handler,
		},
		{
			MethodName: "ListSportsAreaTypes",
			Handler:    _ApiService_ListSportsAreaTypes_Handler,
		},
		{
			MethodName: "ListSportKinds",
			Handler:    _ApiService_ListSportKinds_Handler,
		},
		{
			MethodName: "ListIntersections",
			Handler:    _ApiService_ListIntersections_Handler,
		},
		{
			MethodName: "ListUnions",
			Handler:    _ApiService_ListUnions_Handler,
		},
		{
			MethodName: "GetDensity",
			Handler:    _ApiService_GetDensity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

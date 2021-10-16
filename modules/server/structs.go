package server

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/controllers/analytics"
	"github.com/techpotion/leaders2021-backend/modules/controllers/filters"
	"github.com/techpotion/leaders2021-backend/modules/controllers/geojsons"
	sportsobjects "github.com/techpotion/leaders2021-backend/modules/controllers/sports_objects"
	sportsobjectsdetailed "github.com/techpotion/leaders2021-backend/modules/controllers/sports_objects_detailed"
)

// implementing api service interface generated from proto
type ApiServiceServer struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) ListSportsObjects(ctx context.Context, in *pb.SportsObjects_ListRequest) (*pb.SportsObjects_ListResponse, error) {
	return sportsobjects.List(ctx, in)
}

func (s *ApiServiceServer) GetSportsObject(ctx context.Context, in *pb.SportsObjects_GetRequest) (*pb.SportsObjects_GetResponse, error) {
	return sportsobjects.Get(ctx, in)
}

func (s *ApiServiceServer) ListSportsObjectsDetailed(ctx context.Context, in *pb.SportsObjectsDetailed_ListRequest) (*pb.SportsObjectsDetailed_ListResponse, error) {
	return sportsobjectsdetailed.List(ctx, in)
}

func (s *ApiServiceServer) GetGeoJsonRegions(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return geojsons.GetRegions(ctx, in)
}

func (s *ApiServiceServer) GetGeoJsonDensityHeatmap(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return geojsons.GetDensityHeatmap(ctx, in)
}

func (s *ApiServiceServer) GetGeoJsonSportObjects(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return geojsons.GetGeoJsonSportObjects(ctx, in)
}

func (s *ApiServiceServer) PolygonAnalytics(ctx context.Context, in *pb.PolygonAnalytics_Request) (*pb.PolygonAnalytics_Response, error) {
	return analytics.PolygonAnalytics(ctx, in)
}

func (s *ApiServiceServer) ListObjectsNames(ctx context.Context, in *pb.ObjectsNames_ListRequest) (*pb.ObjectsNames_ListResponse, error) {
	return filters.ListObjectsNames(ctx, in)
}

func (s *ApiServiceServer) ListDepartmentalOrganizationsIds(ctx context.Context, in *pb.DepartmentalOrganizationsIds_ListRequest) (*pb.DepartmentalOrganizationsIds_ListResponse, error) {
	return filters.ListDepartmentalOrganizationsIds(ctx, in)
}

func (s *ApiServiceServer) ListDepartmentalOrganizationsNames(ctx context.Context, in *pb.DepartmentalOrganizationsNames_ListRequest) (*pb.DepartmentalOrganizationsNames_ListResponse, error) {
	return filters.ListDepartmentalOrganizationsNames(ctx, in)
}

func (s *ApiServiceServer) ListSportKinds(ctx context.Context, in *pb.SportKinds_ListRequest) (*pb.SportKinds_ListResponse, error) {
	return filters.ListSportKinds(ctx, in)
}

func (s *ApiServiceServer) ListSportsAreaNames(ctx context.Context, in *pb.SportsAreaNames_ListRequest) (*pb.SportsAreaNames_ListResponse, error) {
	return filters.ListSportsAreaNames(ctx, in)
}

func (s *ApiServiceServer) ListSportsAreaTypes(ctx context.Context, in *pb.SportsAreaTypes_ListRequest) (*pb.SportsAreaTypes_ListResponse, error) {
	return filters.ListSportsAreaTypes(ctx, in)
}

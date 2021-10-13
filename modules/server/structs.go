package server

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/controllers/analytics"
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

func (s *ApiServiceServer) GetGeoJsonObjects(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return geojsons.GetDensityHeatmap(ctx, in)
}

func (s *ApiServiceServer) GetPolygonAnalytics(ctx context.Context, in *pb.PolygonAnalytics_Request) (*pb.PolygonAnalytics_Response, error) {
	return analytics.GetPolygonAnalytics(ctx, in)
}

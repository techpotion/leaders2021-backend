package geojsons

import (
	"context"
	"io/ioutil"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	regionsPath        = "static/moscow_regions.geojson"
	densityHeatmapPath = "static/density.geojson"
	objectsPath        = "static/objects.geojson"
)

func GetRegions(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return handleGeoJson(regionsPath)
}

func GetDensityHeatmap(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return handleGeoJson(densityHeatmapPath)
}

func GetGeoJsonSportsObjects(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return handleGeoJson(objectsPath)
}

func handleGeoJson(path string) (*pb.GeoJsons_Response, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := &pb.GeoJsons_Response{
		GeoJson: string(jsonBytes),
	}

	return resp, nil
}

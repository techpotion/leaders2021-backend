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

// GetRegions is a geojson with moscow regions handler
func GetRegions(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return handleGeoJson(regionsPath)
}

// GetDensityHeatmap is a geojson with moscow population density
func GetDensityHeatmap(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return handleGeoJson(densityHeatmapPath)
}

// GetGeoJsonSportsObjects is a geojson with sports objects density
func GetGeoJsonSportsObjects(ctx context.Context, in *pb.GeoJsons_Request) (*pb.GeoJsons_Response, error) {
	return handleGeoJson(objectsPath)
}

// handleGeoJson is a universal geojson proxy method
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

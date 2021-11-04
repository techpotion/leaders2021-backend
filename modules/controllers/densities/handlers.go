package densities

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func GetDensity(ctx context.Context, in *pb.Densities_GetRequest) (*pb.Densities_GetResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	region := &pb.RegionORM{}
	polygonContaintsQuery := analytics.FormPolygonContainsPointQuery(in.Point)
	result := db.Where(polygonContaintsQuery).Find(&region)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &pb.Densities_GetResponse{Density: region.Density}, nil
}

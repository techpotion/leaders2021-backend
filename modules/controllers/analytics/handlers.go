package analytics

import (
	"context"
	"errors"
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func GetPolygonAnalytics(ctx context.Context, in *pb.PolygonAnalytics_GetRequest) (*pb.PolygonAnalytics_GetResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	filter := &pb.SportsObjectDetailedORM{
		SportsAreaType: in.SportsAreaType,
		Availability:   uint32(in.Availability),
		SportKind:      in.SportKind,
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var objectsList []*pb.SportsObjectDetailedORM

	polygonQuery := analytics.FormPolygonContainsQuery(in.Polygon)

	result := db.Where(filter).Where(polygonQuery).Find(&objectsList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.SportsObjectDetailed
	for _, objectORM := range objectsList {
		converted, err := objectORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	fmt.Println(&pb.SportsObjectsDetailed_ListResponse{
		SportsObjects: convertedList,
		ListStats:     &pb.ListStats{Count: uint32(result.RowsAffected)},
	})
	return nil, nil
}

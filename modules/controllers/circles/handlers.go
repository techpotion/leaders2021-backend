package circles

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ListCircles(ctx context.Context, in *pb.Circles_ListRequest) (*pb.Circles_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var circlesList []*pb.CircleORM

	result := db.Where(pb.CircleORM{Availability: uint32(in.Availability)}).Find(&circlesList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.Circle
	for _, circleORM := range circlesList {
		converted, err := circleORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	return &pb.Circles_ListResponse{
		Intersections: convertedList,
		ListStats:     &pb.ListStats{Count: uint32(len(convertedList))},
	}, nil
}

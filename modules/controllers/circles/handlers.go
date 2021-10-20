package circles

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/circles"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ListCircles(ctx context.Context, in *pb.Circles_ListRequest) (*pb.Circles_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	query := circles.FormIntersectionsQuery(uint32(in.Availability))

	var intersections []circles.CircleIntersectionsORM

	db.Raw(query).Scan(&intersections)

	var convertedList []*pb.CircleIntersections
	for _, intersection := range intersections {
		converted, err := intersection.ToPB(ctx)
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

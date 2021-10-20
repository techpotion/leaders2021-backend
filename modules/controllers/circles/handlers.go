package circles

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/circles"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ListIntersections(ctx context.Context, in *pb.Intersections_ListRequest) (*pb.Intersections_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	query := circles.FormIntersectionsQuery(uint32(in.Availability), in.Polygon)

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
	return &pb.Intersections_ListResponse{
		Intersections: convertedList,
		ListStats:     &pb.ListStats{Count: uint32(len(convertedList))},
	}, nil
}

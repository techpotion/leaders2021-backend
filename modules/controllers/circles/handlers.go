package circles

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/circles"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ListIntersections(ctx context.Context, in *pb.Intersections_ListRequest) (*pb.Intersections_ListResponse, error) {
	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	query := circles.FormIntersectionsQuery(uint32(in.Availability), in.Polygon)

	var intersections []*pb.CirclesClusterORM

	db.Raw(query).Scan(&intersections)

	var convertedList []*pb.CirclesCluster
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

func ListUnions(ctx context.Context, in *pb.Unions_ListRequest) (*pb.Unions_ListResponse, error) {
	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	query := circles.FormUnionsQuery(uint32(in.Availability), in.Polygon)

	var intersections []*pb.CirclesClusterORM

	db.Raw(query).Scan(&intersections)

	var convertedList []*pb.CirclesCluster
	for _, intersection := range intersections {
		converted, err := intersection.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}
	return &pb.Unions_ListResponse{
		Unions:    convertedList,
		ListStats: &pb.ListStats{Count: uint32(len(convertedList))},
	}, nil
}

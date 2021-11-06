package circles

import (
	"context"
	"sync"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/circles"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListIntersections is a list intersections endpoint handler that returns geojsons with intersections
func ListIntersections(ctx context.Context, in *pb.Intersections_ListRequest) (*pb.Intersections_ListResponse, error) {
	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var unions *pb.Unions_ListResponse
	var unionError error
	var wg sync.WaitGroup

	wg.Add(1)
	if in.WithUnion {
		go func() {
			defer wg.Done()
			unions, unionError = ListUnions(ctx, &pb.Unions_ListRequest{
				Polygon:                       in.Polygon,
				DepartmentalOrganizationNames: in.DepartmentalOrganizationNames,
				SportsAreaNames:               in.SportsAreaNames,
				SportsAreaTypes:               in.SportsAreaTypes,
				Availability:                  in.Availability,
				SportKinds:                    in.SportKinds,
			})
		}()
	} else {
		wg.Done()
	}

	query := circles.FormIntersectionsQuery(in.Polygon, uint32(in.Availability), in.SportKinds, in.DepartmentalOrganizationNames, in.SportsAreaNames, in.SportsAreaTypes)

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

	wg.Wait()
	if unionError != nil {
		return nil, status.Error(codes.Internal, unionError.Error())
	}

	var unionsResp []*pb.CirclesCluster
	var unionsListStats *pb.ListStats = &pb.ListStats{Count: 0}
	if unions != nil {
		unionsResp = unions.Unions
		unionsListStats = unions.ListStats
	}

	return &pb.Intersections_ListResponse{
		Intersections:   convertedList,
		ListStats:       &pb.ListStats{Count: uint32(len(convertedList))},
		Unions:          unionsResp,
		UnionsListStats: unionsListStats,
	}, nil
}

// ListUnions is a list unions endpoint handler that returns geojsons with unions
func ListUnions(ctx context.Context, in *pb.Unions_ListRequest) (*pb.Unions_ListResponse, error) {
	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	query := circles.FormUnionsQuery(in.Polygon, uint32(in.Availability), in.SportKinds, in.DepartmentalOrganizationNames, in.SportsAreaNames, in.SportsAreaTypes)

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

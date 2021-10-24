package sportsobjectsdetailed

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"github.com/techpotion/leaders2021-backend/modules/filters"
	sportsobjectsdetailed "github.com/techpotion/leaders2021-backend/modules/sports_objects_detailed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// List is a ListSportsObjectsDetailed handler that returns the list of sports objects detailed by required filters
func List(ctx context.Context, in *pb.SportsObjectsDetailed_ListRequest) (*pb.SportsObjectsDetailed_ListResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Polygon != nil {
		if err := analytics.ValidatePolygon(in.Polygon); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var objectsList []*pb.SportsObjectDetailedORM
	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())

	result := db.Limit(lim).Offset(offset*lim).Scopes(
		filters.ObjectIdsScope(in.ObjectIds, sportsobjectsdetailed.TableName),
		filters.ObjectNamesScope(in.ObjectNames, sportsobjectsdetailed.TableName),
		filters.DepartmentalOrganizationIdsScope(in.DepartmentalOrganizationIds, sportsobjectsdetailed.TableName),
		filters.DepartmentalOrganizationNamesScope(in.DepartmentalOrganizationNames, sportsobjectsdetailed.TableName),
		filters.SportsAreaNamesScope(in.SportsAreaNames, sportsobjectsdetailed.TableName),
		filters.SportsAreaTypeScope(in.SportsAreaTypes, sportsobjectsdetailed.TableName),
		filters.AvailabilitiesScope(in.Availabilities, sportsobjectsdetailed.TableName),
		filters.SportKindsScope(in.SportKinds, sportsobjectsdetailed.TableName),
		filters.PolygonScope(in.Polygon),
	).Find(&objectsList)
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

	return &pb.SportsObjectsDetailed_ListResponse{
		SportsObjects: convertedList,
		ListStats:     &pb.ListStats{Count: uint32(result.RowsAffected)},
	}, nil
}

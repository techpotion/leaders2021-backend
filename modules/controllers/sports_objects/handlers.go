package sportsobjects

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"github.com/techpotion/leaders2021-backend/modules/filters"
	sportsobjects "github.com/techpotion/leaders2021-backend/modules/sports_objects"
	sportsobjectsdetailed "github.com/techpotion/leaders2021-backend/modules/sports_objects_detailed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func List(ctx context.Context, in *pb.SportsObjects_ListRequest) (*pb.SportsObjects_ListResponse, error) {
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

	var objectsList []*pb.SportsObjectORM
	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())

	result := db.Limit(lim).Offset(offset*lim).Scopes(
		filters.ObjectNamesScope(in.ObjectNames, sportsobjects.TableName),
		filters.DepartmentalOrganizationIdsScope(in.DepartmentalOrganizationIds, sportsobjects.TableName),
		filters.DepartmentalOrganizationNamesScope(in.DepartmentalOrganizationNames, sportsobjects.TableName),
		filters.AvailabilitiesScope(in.Availabilities, sportsobjects.TableName),
		filters.PolygonScope(in.Polygon),
	).Find(&objectsList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.SportsObject
	for _, objectORM := range objectsList {
		converted, err := objectORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	return &pb.SportsObjects_ListResponse{
		SportsObjects: convertedList,
		ListStats:     &pb.ListStats{Count: uint32(result.RowsAffected)},
	}, nil
}

func Get(ctx context.Context, in *pb.SportsObjects_GetRequest) (*pb.SportsObjects_GetResponse, error) {
	filter := &pb.SportsObjectORM{
		ObjectId: in.ObjectId,
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var object pb.SportsObjectORM
	result := db.Where(filter).First(&object)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	converted, err := object.ToPB(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.SportsObjects_GetResponse{
		SportsObject: &converted,
	}, nil
}

func ListFromDetailed(ctx context.Context, in *pb.SportsObjectsDetailed_ListRequest) (*pb.SportsObjects_ListResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var objectsList []*pb.SportsObjectORM
	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())

	result := db.Model(&pb.SportsObjectDetailedORM{}).Limit(lim).Offset(offset*lim).Scopes(
		filters.ObjectIdsScope(in.ObjectIds, sportsobjectsdetailed.TableName),
		filters.ObjectNamesScope(in.ObjectNames, sportsobjectsdetailed.TableName),
		filters.DepartmentalOrganizationIdsScope(in.DepartmentalOrganizationIds, sportsobjectsdetailed.TableName),
		filters.DepartmentalOrganizationNamesScope(in.DepartmentalOrganizationNames, sportsobjectsdetailed.TableName),
		filters.SportsAreaNamesScope(in.SportsAreaNames, sportsobjectsdetailed.TableName),
		filters.SportsAreaTypeScope(in.SportsAreaTypes, sportsobjectsdetailed.TableName),
		filters.AvailabilitiesScope(in.Availabilities, sportsobjectsdetailed.TableName),
		filters.SportKindsScope(in.SportKinds, sportsobjectsdetailed.TableName),
		filters.PolygonScope(in.Polygon),
	).Joins("INNER JOIN objects ON objects.object_id = objects_detailed.object_id").Find(&objectsList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.SportsObject
	for _, objectORM := range objectsList {
		converted, err := objectORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	convertedList = sportsobjects.UniqueObjects(convertedList)

	return &pb.SportsObjects_ListResponse{
		SportsObjects: convertedList,
		ListStats:     &pb.ListStats{Count: uint32(len(convertedList))},
	}, nil
}

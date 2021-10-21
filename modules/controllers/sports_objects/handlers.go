package sportsobjects

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"github.com/techpotion/leaders2021-backend/modules/filters"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func List(ctx context.Context, in *pb.SportsObjects_ListRequest) (*pb.SportsObjects_ListResponse, error) {
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

	result := db.Limit(lim).Offset(offset*lim).Scopes(
		filters.ObjectNamesScope(in.ObjectNames),
		filters.DepartmentalOrganizationIdsScope(in.DepartmentalOrganizationIds),
		filters.DepartmentalOrganizationNamesScope(in.DepartmentalOrganizationNames),
		filters.AvailabilitiesScope(in.Availabilities),
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

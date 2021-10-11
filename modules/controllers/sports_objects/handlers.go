package sportsobjects

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func List(ctx context.Context, in *pb.SportsObjects_ListRequest) (*pb.SportsObjects_ListResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	filter := &pb.SportsObjectORM{
		ObjectName:                   in.ObjectName,
		DepartmentalOrganizationId:   in.DepartmentalOrganizationId,
		DepartmentalOrganizationName: in.DepartmentalOrganizationName,
		Availability:                 uint32(in.Availability),
		SportKind:                    in.SportKind,
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var objectsList []*pb.SportsObjectORM
	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber()) * lim

	result := db.Limit(lim).Offset(offset * lim).Where(filter).Find(&objectsList)
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
		ListStats:     &pb.ListStats{Count: uint64(result.RowsAffected)},
	}, nil
}

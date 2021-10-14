package filters

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ListObjectsNames(ctx context.Context, in *pb.ObjectsNames_ListRequest) (*pb.ObjectsNames_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var names []string

	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())
	result := db.Model(&pb.SportsObjectORM{}).Limit(lim).Offset(offset*lim).Where("object_name IS NOT NULL").Distinct().Pluck("object_name", &names)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &pb.ObjectsNames_ListResponse{
		Names:     names,
		ListStats: &pb.ListStats{Count: uint32(len(names))},
	}, nil
}

func ListDepartmentalOrganizationsIds(ctx context.Context, in *pb.DepartmentalOrganizationsIds_ListRequest) (*pb.DepartmentalOrganizationsIds_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var ids []uint32

	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())
	result := db.Model(&pb.SportsObjectORM{}).Limit(lim).Offset(offset*lim).Where("departmental_organization_id IS NOT NULL").Distinct().Pluck("departmental_organization_id", &ids)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &pb.DepartmentalOrganizationsIds_ListResponse{
		Ids:       ids,
		ListStats: &pb.ListStats{Count: uint32(len(ids))},
	}, nil
}

func ListDepartmentalOrganizationsNames(ctx context.Context, in *pb.DepartmentalOrganizationsNames_ListRequest) (*pb.DepartmentalOrganizationsNames_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var names []string

	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())
	result := db.Model(&pb.SportsObjectORM{}).Limit(lim).Offset(offset*lim).Where("departmental_organization_name IS NOT NULL").Distinct().Pluck("departmental_organization_name", &names)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &pb.DepartmentalOrganizationsNames_ListResponse{
		Names:     names,
		ListStats: &pb.ListStats{Count: uint32(len(names))},
	}, nil
}

func ListSportKinds(ctx context.Context, in *pb.SportKinds_ListRequest) (*pb.SportKinds_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var sportKinds []string

	lim := int(in.Pagination.GetResultsPerPage())
	offset := int(in.Pagination.GetPageNumber())
	result := db.Model(&pb.SportsObjectDetailedORM{}).Limit(lim).Offset(offset*lim).Where("sport_kind IS NOT NULL").Distinct().Pluck("sport_kind", &sportKinds)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &pb.SportKinds_ListResponse{
		Kinds:     sportKinds,
		ListStats: &pb.ListStats{Count: uint32(len(sportKinds))},
	}, nil
}

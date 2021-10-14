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
	return nil, nil
}

func ListDepartmentalOrganizationsIds(ctx context.Context, in *pb.DepartmentalOrganizationsIds_ListRequest) (*pb.DepartmentalOrganizationsIds_ListResponse, error) {
	return nil, nil
}

func ListDepartmentalOrganizationsNames(ctx context.Context, in *pb.DepartmentalOrganizationsNames_ListRequest) (*pb.DepartmentalOrganizationsNames_ListResponse, error) {
	return nil, nil
}

func ListSportKinds(ctx context.Context, in *pb.SportKinds_ListRequest) (*pb.SportKinds_ListResponse, error) {
	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var sportKinds []string

	result := db.Model(&pb.SportsObjectDetailedORM{}).Distinct().Pluck("sport_kind", &sportKinds)
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

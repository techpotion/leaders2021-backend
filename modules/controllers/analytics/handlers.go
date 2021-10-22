package analytics

import (
	"context"
	"errors"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func PolygonAnalytics(ctx context.Context, in *pb.PolygonAnalytics_Request) (*pb.PolygonAnalytics_Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	filter := &pb.SportsObjectDetailedORM{
		SportsAreaType: in.SportsAreaType,
		Availability:   uint32(in.Availability),
		SportKind:      in.SportKind,
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var objectsList []*pb.SportsObjectDetailedORM

	polygonQuery := analytics.FormPolygonContainsQuery(in.Polygon)

	result := db.Select("\"sport_kind\", \"sports_area_type\", \"sports_area_square\"").Where(filter).Where(polygonQuery).Find(&objectsList)
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

	region := &pb.RegionORM{}
	polygonCenter := analytics.CalculatePolygonCenter(in.Polygon)
	polygonContaintsQuery := analytics.FormPolygonContainsPointQuery(polygonCenter)
	result = db.Where(polygonContaintsQuery).Find(&region)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var polygonSquare float64
	polygonSquareQuery := analytics.CalculatePolygonSquareQuery(in.Polygon)
	result = db.Raw(polygonSquareQuery).Scan(&polygonSquare)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	personsOnSquare := polygonSquare * region.Density

	areasSquare := analytics.CalculateSquare(convertedList)
	areasAmount := len(convertedList)
	areasSquarePerPerson := float64(areasSquare) / personsOnSquare

	sportsKinds := analytics.UniqueSportsKinds(convertedList)
	sportsAmount := len(sportsKinds)
	sportsAmountPerPerson := float64(sportsAmount) / personsOnSquare

	areaTypes := analytics.UniqueAreaTypes(convertedList)
	areaTypesAmount := len(areaTypes)
	areaTypesAmountPerPerson := float64(areaTypesAmount) / personsOnSquare

	return &pb.PolygonAnalytics_Response{
		AreasSquare:           areasSquare,
		AreasSquarePerPerson:  areasSquarePerPerson,
		AreasAmount:           uint32(areasAmount),
		SportsKinds:           sportsKinds,
		SportsAmount:          uint32(sportsAmount),
		SportsAmountPerPerson: sportsAmountPerPerson,
		AreaTypes:             areaTypes,
		AreaTypesAmount:       uint32(areaTypesAmount),
		AreasAmountPerPerson:  areaTypesAmountPerPerson,
	}, nil
}

func PolygonParkAnalytics(ctx context.Context, in *pb.PolygonParkAnalytics_Request) (*pb.PolygonParkAnalytics_Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var parksList []*pb.ParkORM

	polygonQuery := analytics.FormPolygonIntersectsParkQuery(in.Polygon)

	result := db.Where(pb.ParkORM{HasSportground: false}).Where("has_sportground = ?", in.HasSportground).Where(polygonQuery).Find(&parksList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.Park
	for _, objectORM := range parksList {
		converted, err := objectORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	return &pb.PolygonParkAnalytics_Response{
		Parks:     convertedList,
		ListStats: &pb.ListStats{Count: uint32(result.RowsAffected)},
	}, nil
}

func PolygonPollutionAnalytics(ctx context.Context, in *pb.PolygonPollutionAnalytics_Request) (*pb.PolygonPollutionAnalytics_Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var parksList []*pb.PollutionORM

	polygonQuery := analytics.FormPolygonContainsQuery(in.Polygon)

	result := db.Where(pb.ParkORM{HasSportground: false}).Where("is_polluted = ?", in.IsPolluted).Where(polygonQuery).Find(&parksList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.Pollution
	for _, objectORM := range parksList {
		converted, err := objectORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	if in.ReturnPoints {
		return &pb.PolygonPollutionAnalytics_Response{
			Points:    convertedList,
			ListStats: &pb.ListStats{Count: uint32(result.RowsAffected)},
		}, nil
	} else {
		return &pb.PolygonPollutionAnalytics_Response{}, nil
	}
}

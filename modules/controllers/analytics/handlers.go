package analytics

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"github.com/techpotion/leaders2021-backend/modules/controllers/marks"
	"github.com/techpotion/leaders2021-backend/modules/database"
	"github.com/techpotion/leaders2021-backend/modules/filters"
	sportsobjectsdetailed "github.com/techpotion/leaders2021-backend/modules/sports_objects_detailed"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// PolygonAnalytics is a polygon analytics endpoint handler
func PolygonAnalytics(ctx context.Context, in *pb.PolygonAnalytics_Request) (*pb.PolygonAnalytics_Response, error) {
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

	var objectsList []*pb.SportsObjectDetailedORM

	polygonQuery := analytics.FormPolygonContainsQuery(in.Polygon)

	result := db.Select("\"object_id\", \"sport_kind\", \"sports_area_type\", \"sports_area_square\"").Scopes(
		filters.DepartmentalOrganizationNamesScope(in.DepartmentalOrganizationNames, sportsobjectsdetailed.TableName),
		filters.SportsAreaNamesScope(in.SportsAreaNames, sportsobjectsdetailed.TableName),
		filters.SportsAreaTypeScope(in.SportsAreaTypes, sportsobjectsdetailed.TableName),
		filters.AvailabilitiesScope(in.Availabilities, sportsobjectsdetailed.TableName),
		filters.SportKindsScope(in.SportKinds, sportsobjectsdetailed.TableName),
	).Where(polygonQuery).Find(&objectsList)
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

	if region.Density == 0 {
		region.Density = 5500
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

	sportsObjectsAmount := len(sportsobjectsdetailed.UniqueObjectsFromDetailed(convertedList))
	sportsObjectsAmountPer100k := float64(sportsObjectsAmount) / personsOnSquare

	return &pb.PolygonAnalytics_Response{
		AreasSquare:                areasSquare,
		AreasAmount:                uint32(areasAmount),
		AreasSquarePer100K:         areasSquarePerPerson * 100000,
		SportsKinds:                sportsKinds,
		SportsAmount:               uint32(sportsAmount),
		SportsAmountPer100K:        sportsAmountPerPerson * 100000,
		AreaTypes:                  areaTypes,
		AreaTypesAmount:            uint32(areaTypesAmount),
		AreasAmountPer100K:         areaTypesAmountPerPerson * 100000,
		SportsObjectsAmount:        uint32(sportsObjectsAmount),
		SportsObjectsAmountPer100K: sportsObjectsAmountPer100k * 100000,
		Density:                    region.Density,
	}, nil
}

// PolygonParkAnalytics is a polygon park analytics endpoint handler
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

	parkSquareQuery := fmt.Sprintf("*, %s as square", analytics.FormPolygonSquareQuery())
	result := db.Select(parkSquareQuery).Where(pb.ParkORM{HasSportground: false}).Where("has_sportground = ?", in.HasSportground).Where(polygonQuery).Find(&parksList)
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
		ListStats: &pb.ListStats{Count: uint32(len(convertedList))},
	}, nil
}

// PolygonPollutionAnalytics is a polygon pollution analytics endpoint handler
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

	var totalPoints int64
	result = db.Model(&pb.PollutionORM{}).Where(pb.ParkORM{HasSportground: false}).Where(polygonQuery).Count(&totalPoints)
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

	var perc float32 = 0.
	if totalPoints != 0 {
		perc = float32(len(convertedList)) / float32(totalPoints)
	}
	if in.ReturnPoints {
		return &pb.PolygonPollutionAnalytics_Response{
			Points:              convertedList,
			PollutionPercentage: perc,
			ListStats:           &pb.ListStats{Count: uint32(len(convertedList))},
		}, nil
	} else {
		return &pb.PolygonPollutionAnalytics_Response{}, nil
	}
}

// PolygonPollutionAnalytics is a polygon pollution analytics endpoint handler
func PolygonSubwayAnalytics(ctx context.Context, in *pb.PolygonSubwayAnalytics_Request) (*pb.PolygonSubwayAnalytics_Response, error) {
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

	var subwaysList []*pb.SubwayORM

	polygonQuery := analytics.FormGeometryPolygon(in.Polygon)

	result := db.
		Select(fmt.Sprintf("*, ST_Distance(%s::geography, position::geography) as distance_from_polygon", polygonQuery)).
		Where(fmt.Sprintf("ST_Distance(%s::geography, position::geography) < 500", polygonQuery)).
		Order(fmt.Sprintf("position <-> %s", polygonQuery)).
		Find(&subwaysList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, result.Error.Error())
		}
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.Subway
	for _, subwayORM := range subwaysList {
		converted, err := subwayORM.ToPB(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	return &pb.PolygonSubwayAnalytics_Response{
		Points:    convertedList,
		ListStats: &pb.ListStats{Count: uint32(len(convertedList))},
	}, nil
}

// PolygonAnalyticsDashboard is a polygon pollution analytics dashboard endpoint handler that compiles all the other methods
func PolygonAnalyticsDashboard(ctx context.Context, in *pb.PolygonAnalyticsDashboard_Request) (*pb.PolygonAnalyticsDashboard_Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var basicAnalytics *pb.PolygonAnalytics_Response
	var parkAnalytics *pb.PolygonParkAnalytics_Response
	var pollutionAnalytics *pb.PolygonPollutionAnalytics_Response
	var subwayAnalytics *pb.PolygonSubwayAnalytics_Response
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		basicAnalytics, err = PolygonAnalytics(ctx, &pb.PolygonAnalytics_Request{
			Polygon:                       in.Polygon,
			DepartmentalOrganizationNames: in.DepartmentalOrganizationNames,
			SportsAreaNames:               in.SportsAreaNames,
			SportsAreaTypes:               in.SportsAreaTypes,
			Availabilities:                in.Availabilities,
			SportKinds:                    in.SportKinds,
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		parkAnalytics, err = PolygonParkAnalytics(ctx, &pb.PolygonParkAnalytics_Request{
			Polygon:        in.Polygon,
			HasSportground: false,
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		pollutionAnalytics, err = PolygonPollutionAnalytics(ctx, &pb.PolygonPollutionAnalytics_Request{
			Polygon:      in.Polygon,
			IsPolluted:   true,
			ReturnPoints: true,
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		subwayAnalytics, err = PolygonSubwayAnalytics(ctx, &pb.PolygonSubwayAnalytics_Request{
			Polygon: in.Polygon,
		})
	}()
	wg.Wait()
	fmt.Printf("%+v", basicAnalytics)
	fmt.Printf("%+v", subwayAnalytics)
	fmt.Printf("%+v", pollutionAnalytics)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	mark, err := marks.GetMark(ctx, &pb.Marks_GetRequest{
		AreasAmountPer100K:  float32(basicAnalytics.AreasAmountPer100K),
		SportsAmountPer100K: float32(basicAnalytics.SportsAmountPer100K),
		AreasSquarePer100K:  float32(basicAnalytics.AreasAmountPer100K),
		SubwayDistance:      float32(subwayAnalytics.Points[0].DistanceFromPolygon),
		PollutionPercentage: pollutionAnalytics.PollutionPercentage,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.PolygonAnalyticsDashboard_Response{
		BasicAnalytics:     basicAnalytics,
		ParkAnalytics:      parkAnalytics,
		PollutionAnalytics: pollutionAnalytics,
		SubwayAnalytics:    subwayAnalytics,
		Mark:               mark.Mark,
	}, nil
}

package analytics

import (
	"errors"
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
)

// TODO add tests
func ValidatePolygon(polygon *pb.Polygon) error {
	polyLen := len(polygon.Points)

	if polyLen < 4 {
		return errors.New("polygon must consist of at least 3 different points")
	}

	if polygon.Points[0].Lat != polygon.Points[polyLen-1].Lat &&
		polygon.Points[0].Lng != polygon.Points[polyLen-1].Lng {
		return errors.New("polygon is not closed: first and last point should be the same")
	}

	return nil
}

// TODO add tests
func FormPolygonContainsQuery(polygon *pb.Polygon) string {
	polygonQuery := FormGeometryPolygon(polygon)
	return fmt.Sprintf(`
		ST_Contains(
			%s,
			position
		)`,
		polygonQuery,
	)
}

// TODO add tests
func FormGeometryPolygon(polygon *pb.Polygon) string {
	points := ""
	for _, point := range polygon.Points {
		points += fmt.Sprintf(",%f %f", point.Lng, point.Lat)
	}
	poly := "POLYGON((" + points[1:] + "))"
	return fmt.Sprintf("ST_GeomFromText('%s', 4326)", poly)
}

// TODO add tests
func CalculateSquare(areas []*pb.SportsObjectDetailed) float64 {
	var sumSquare float64 = 0
	for _, area := range areas {
		sumSquare += area.GetSportsAreaSquare()
	}
	return sumSquare
}

// TODO add tests
func UniqueSportsKinds(areas []*pb.SportsObjectDetailed) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, area := range areas {
		kind := area.SportKind
		if _, ok := keys[kind]; !ok {
			keys[kind] = true
			list = append(list, kind)
		}
	}

	return list
}

// TODO add tests
func UniqueAreaTypes(areas []*pb.SportsObjectDetailed) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, area := range areas {
		kind := area.SportsAreaType
		if _, ok := keys[kind]; !ok {
			keys[kind] = true
			list = append(list, kind)
		}
	}

	return list
}

func FormPolygonIntersectsParkQuery(polygon *pb.Polygon) string {
	polygonQuery := FormGeometryPolygon(polygon)
	return fmt.Sprintf(`
		ST_Intersects(
			%s,
			polygon
		)`,
		polygonQuery,
	)
}

func FormPolygonContainsPointQuery(point *pb.Point) string {
	return fmt.Sprintf(`
		ST_Intersects(
			polygon,
			ST_SetSRID(ST_Point(%f, %f), 4326)
		)`,
		point.Lng,
		point.Lat,
	)
}

func CalculatePolygonCenter(poly *pb.Polygon) *pb.Point {
	pointsLen := len(poly.Points)
	var lat float32 = 0.
	var lng float32 = 0.
	for _, point := range poly.Points {
		lat += point.Lat
		lng += point.Lng
	}
	return &pb.Point{
		Lat: lat / float32(pointsLen),
		Lng: lng / float32(pointsLen),
	}
}

func CalculatePolygonSquareQuery(polygon *pb.Polygon) string {
	polygonQuery := FormGeometryPolygon(polygon)
	return fmt.Sprintf(`
		SELECT ST_Area(%s::geography)/1000000`,
		polygonQuery,
	)
}

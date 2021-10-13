package analytics

import (
	"errors"
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
)

// TODO add tests
func ValidatePolygon(polygon *pb.Polygon) error {
	polyLen := len(polygon.Points)

	if polyLen < 5 {
		return errors.New("polygon must consist of at least 4 different points")
	}

	if polygon.Points[0].Lat != polygon.Points[polyLen-1].Lat &&
		polygon.Points[0].Lng != polygon.Points[polyLen-1].Lng {
		return errors.New("polygon is not closed: first and last point should be the same")
	}

	return nil
}

// TODO add tests
func FormPolygonContainsQuery(polygon *pb.Polygon) string {
	polygonQuery := formGeometryPolygon(polygon)
	return fmt.Sprintf(`
		ST_Contains(
			%s,
			ST_Point(object_point_lng, object_point_lat)
		);`,
		polygonQuery,
	)
}

// TODO add tests
func formGeometryPolygon(polygon *pb.Polygon) string {
	// "POLYGON((37.568742 55.776336,37.563248 55.715140,37.679978 55.719791,37.668305 55.799550,37.568742 55.776336))"
	points := ""
	for _, point := range polygon.Points {
		points += fmt.Sprintf(",%f %f", point.Lng, point.Lat)
	}
	poly := "POLYGON((" + points[1:] + "))"
	return fmt.Sprintf("ST_GeomFromText('%s')", poly)
}

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

	if polygon.Points[0] != polygon.Points[polyLen-1] {
		return errors.New("polygon is not closed: first and last point should be the same")
	}

	return nil
}

// TODO add tests
func FormPolygonEntranceQuery(polygon *pb.Polygon) string {
	polygonQuery := formGeometryPolygon(polygon)
	return fmt.Sprintf(`
		SELECT count(*) FROM public.objects\n
		WHERE ST_Contains(
			ST_GeomFromText('%s'),
			ST_Point(object_point_lng, object_point_lat)
		);`,
		polygonQuery,
	)
}

// TODO add tests
func formGeometryPolygon(polygon *pb.Polygon) string {
	// "POLYGON((37.568742 55.776336,37.563248 55.715140,37.679978 55.719791,37.668305 55.799550,37.568742 55.776336))"
	head := "POLYGON(("
	tail := "))"
	points := ""
	for _, point := range polygon.Points {
		points += fmt.Sprintf(",%f %f", point.Lng, point.Lat)
	}
	return head + points[1:] + tail
}

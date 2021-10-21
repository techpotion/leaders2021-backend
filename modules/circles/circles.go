package circles

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
)

func FormIntersectionsQuery(availability uint32, polygon *pb.Polygon) string {
	polyQuery := FormPolygonContainsQuery(polygon)
	return fmt.Sprintf(`
		SELECT ST_AsGeoJSON(ST_Difference(ST_Union(circle), ST_MakeValid(ST_Collect(circle)))) AS geojson, SUM(object_sum_square) as square FROM (
			SELECT *, ST_ClusterDBSCAN(circle, 0, 1) OVER() AS clst FROM objects
			WHERE availability = %d AND %s
		) q
		GROUP BY clst`,
		availability,
		polyQuery,
	)
}

func FormUnionsQuery(availability uint32, polygon *pb.Polygon) string {
	polyQuery := FormPolygonContainsQuery(polygon)
	return fmt.Sprintf(`
		SELECT ST_AsGeoJSON(ST_Union(circle)) AS geojson, SUM(object_sum_square) as square FROM (
			SELECT *, ST_ClusterDBSCAN(circle, 0, 1) OVER() AS clst FROM objects
			WHERE availability = %d AND %s
		) q
		GROUP BY clst`,
		availability,
		polyQuery,
	)
}

package circles

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
)

// FormIntersectionsQuery forms SQL query for getting geojsons of
// circles intersections - http://postgis.net/workshops/postgis-intro/geometry_returning.html
func FormIntersectionsQuery(availability uint32, polygon *pb.Polygon) string {
	polyQuery := FormPolygonContainsQuery(polygon)
	return fmt.Sprintf(`
		WITH t1 AS (SELECT object_id, circle, SUM(object_sum_square) AS square FROM (
			SELECT *, ST_ClusterDBSCAN(circle, 0, 1) OVER() AS clst FROM objects
			WHERE availability = %d AND %s
		) q
		GROUP BY clst, circle, object_id)

		SELECT (ST_AsGeoJSON(ST_intersection(a.circle, b.circle))) as geojson, a.square + b.square as square FROM t1 a, t1 b
		WHERE ST_Intersects (a.circle,b.circle) AND a.object_id < b.object_id`,
		availability,
		polyQuery,
	)
}

// FormIntersectionsQuery forms SQL query for getting geojsons of
// united circles - http://postgis.net/workshops/postgis-intro/geometry_returning.html
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

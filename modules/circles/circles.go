package circles

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
)

func FormIntersectionsQuery(availability uint32, polygon *pb.Polygon) string {
	polyQuery := FormPolygonContainsQuery(polygon)
	return fmt.Sprintf(`
		SELECT c.object_id, main_object.object_sum_square as object_square, array_agg(c.neighbor_object_id) AS intersections FROM circles AS c

		LEFT JOIN objects AS main_object
		ON main_object.object_id = c.object_id

		LEFT JOIN objects AS neighbor_object
		ON neighbor_object.object_id = c.neighbor_object_id

		WHERE c.availability = %d AND %s

		GROUP BY c.object_id, main_object.object_sum_square`,
		availability,
		polyQuery,
	)
}

package circles

import (
	"fmt"
)

func FormIntersectionsQuery(availability uint32) string {
	return fmt.Sprintf(`
		SELECT c.object_id, array_agg(c.neighbor_object_id) AS intersections, sum(neighbor_object.object_sum_square)+main_object.object_sum_square AS square FROM circles AS c

		LEFT JOIN objects AS main_object
		ON main_object.object_id = c.object_id

		LEFT JOIN objects AS neighbor_object
		ON neighbor_object.object_id = c.neighbor_object_id

		WHERE c.availability = %d

		GROUP BY c.object_id, main_object.object_sum_square, main_object.position`,
		availability,
	)
}

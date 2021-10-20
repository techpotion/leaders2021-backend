package circles

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
)

// TODO add tests
func FormPolygonContainsQuery(polygon *pb.Polygon) string {
	polygonQuery := analytics.FormGeometryPolygon(polygon)
	return fmt.Sprintf(`
		ST_Contains(
			%s,
			main_object.position
		)`,
		polygonQuery,
	)
}

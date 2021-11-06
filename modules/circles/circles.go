package circles

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
)

// FormIntersectionsQuery forms SQL query for getting geojsons of
// circles intersections - http://postgis.net/workshops/postgis-intro/geometry_returning.html
func FormIntersectionsQuery(polygon *pb.Polygon,
	availability uint32,
	sportKinds []string,
	departmentalOrganizationNames []string,
	sportsAreaNames []string,
	sportsAreaTypes []string,
) string {
	sportKindsQuery := formSportKindsQuery(sportKinds)
	departmentalOrganizationNamesQuery := formDepartmentalOrganizationNamesQuery(departmentalOrganizationNames)
	sportsAreaNamesQuery := formSportsAreaNamesQuery(sportsAreaNames)
	sportsAreaTypesQuery := formSportsAreaNamesQuery(sportsAreaTypes)
	polyQuery := FormPolygonContainsQuery(polygon)
	return fmt.Sprintf(`
		WITH t1 AS (SELECT object_id, circle, SUM(sports_area_square) AS square FROM (
			SELECT *, ST_ClusterDBSCAN(circle, 0, 1) OVER() AS clst FROM objects_detailed
			WHERE availability = %d AND %s AND %s AND %s AND %s AND %s
		) q
		GROUP BY clst, circle, object_id)

		SELECT (ST_AsGeoJSON(ST_intersection(a.circle, b.circle))) as geojson, a.square + b.square as square FROM t1 a, t1 b
		WHERE ST_Intersects (a.circle,b.circle) AND a.object_id < b.object_id`,
		availability,
		polyQuery,
		sportKindsQuery,
		departmentalOrganizationNamesQuery,
		sportsAreaNamesQuery,
		sportsAreaTypesQuery,
	)
}

// FormIntersectionsQuery forms SQL query for getting geojsons of
// united circles - http://postgis.net/workshops/postgis-intro/geometry_returning.html
func FormUnionsQuery(polygon *pb.Polygon,
	availability uint32,
	sportKinds []string,
	departmentalOrganizationNames []string,
	sportsAreaNames []string,
	sportsAreaTypes []string,
) string {
	sportKindsQuery := formSportKindsQuery(sportKinds)
	departmentalOrganizationNamesQuery := formDepartmentalOrganizationNamesQuery(departmentalOrganizationNames)
	sportsAreaNamesQuery := formSportsAreaNamesQuery(sportsAreaNames)
	sportsAreaTypesQuery := formSportsAreaNamesQuery(sportsAreaTypes)
	polyQuery := FormPolygonContainsQuery(polygon)
	return fmt.Sprintf(`
		WITH t1 AS (SELECT object_id, circle, SUM(sports_area_square) AS square FROM (
			SELECT *, ST_ClusterDBSCAN(circle, 0, 1) OVER() AS clst FROM objects_detailed
			WHERE availability = %d AND %s AND %s AND %s AND %s AND %s
		) q
		GROUP BY clst, circle, object_id)

		SELECT ST_AsGeoJSON(ST_Union(ST_intersection(a.circle, b.circle))) as geojson, SUM(a.square + b.square) as square FROM t1 a, t1 b
		WHERE ST_Intersects (a.circle,b.circle) AND a.object_id < b.object_id`,
		availability,
		polyQuery,
		sportKindsQuery,
		departmentalOrganizationNamesQuery,
		sportsAreaNamesQuery,
		sportsAreaTypesQuery,
	)
}

func formSportKindsQuery(sportKinds []string) string {
	sportKindsQuery := "True"
	if sportKinds != nil || len(sportKinds) > 0 {
		sportKindsQuery = "("
		for _, kind := range sportKinds {
			sportKindsQuery += "'" + kind + "'" + ","
		}
		sportKindsQuery = "sport_kind IN " + sportKindsQuery[:len(sportKindsQuery)-1] + ")"
	}
	return sportKindsQuery
}

func formDepartmentalOrganizationNamesQuery(departmentalOrganizationNames []string) string {
	depOrgQuery := "True"
	if departmentalOrganizationNames != nil || len(departmentalOrganizationNames) > 0 {
		depOrgQuery = "("
		for _, name := range departmentalOrganizationNames {
			depOrgQuery += "'" + name + "'" + ","
		}
		depOrgQuery = "departmental_organization_name IN " + depOrgQuery[:len(depOrgQuery)-1] + ")"
	}
	return depOrgQuery
}

func formSportsAreaNamesQuery(sportsAreaNames []string) string {
	areaNameQuery := "True"
	if sportsAreaNames != nil || len(sportsAreaNames) > 0 {
		areaNameQuery = "("
		for _, name := range sportsAreaNames {
			areaNameQuery += "'" + name + "'" + ","
		}
		areaNameQuery = "sports_area_name IN " + areaNameQuery[:len(areaNameQuery)-1] + ")"
	}
	return areaNameQuery
}

func formSportsAreaTypesQuery(sportsAreaTypes []string) string {
	areaTypeQuery := "True"
	if sportsAreaTypes != nil || len(sportsAreaTypes) > 0 {
		areaTypeQuery = "("
		for _, name := range sportsAreaTypes {
			areaTypeQuery += "'" + name + "'" + ","
		}
		areaTypeQuery = "sports_area_type IN " + areaTypeQuery[:len(areaTypeQuery)-1] + ")"
	}
	return areaTypeQuery
}

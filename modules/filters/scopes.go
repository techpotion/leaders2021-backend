package filters

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"gorm.io/gorm"
)

// ObjectIdsScope is an object_id filter which applies if valid list was provided
func ObjectIdsScope(ids []uint32, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where(fmt.Sprintf("%s.object_id IN ?", tableName), ids)
		}
		return db
	}
}

// ObjectNamesScope is an object_name filter which applies if valid list was provided
func ObjectNamesScope(names []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where(fmt.Sprintf("%s.object_name IN ?", tableName), names)
		}
		return db
	}
}

// DepartmentalOrganizationIdsScope is a departmental_organization_id filter which applies if valid list was provided
func DepartmentalOrganizationIdsScope(ids []uint32, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where(fmt.Sprintf("%s.departmental_organization_id IN ?", tableName), ids)
		}
		return db
	}
}

// SportsAreaNamesScope is a sports_area_name filter which applies if valid list was provided
func SportsAreaNamesScope(names []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where(fmt.Sprintf("%s.sports_area_name IN ?", tableName), names)
		}
		return db
	}
}

// SportsAreaTypeScope is a sports_area_type filter which applies if valid list was provided
func SportsAreaTypeScope(types []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(types) > 0 {
			return db.Where(fmt.Sprintf("%s.sports_area_type IN ?", tableName), types)
		}
		return db
	}
}

// DepartmentalOrganizationNamesScope is a departmental_organization_name filter which applies if valid list was provided
func DepartmentalOrganizationNamesScope(names []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where(fmt.Sprintf("%s.departmental_organization_name IN ?", tableName), names)
		}
		return db
	}
}

// AvailabilitiesScope is an availability filter which applies if valid list was provided
func AvailabilitiesScope(availabilities []pb.Availability, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(availabilities) > 0 {
			return db.Where(fmt.Sprintf("%s.availability IN ?", tableName), availabilities)
		}
		return db
	}
}

// SportKindsScope is a sport_kind filter which applies if valid list was provided
func SportKindsScope(kinds []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(kinds) > 0 {
			return db.Where(fmt.Sprintf("%s.sport_kind IN ?", tableName), kinds)
		}
		return db
	}
}

// PolygonScope is a sport_kind filter which applies if valid polygon was provided
func PolygonScope(polygon *pb.Polygon) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if polygon != nil {
			polyQuery := analytics.FormPolygonContainsQuery(polygon)
			return db.Where(polyQuery)
		}
		return db
	}
}

package filters

import (
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"gorm.io/gorm"
)

func ObjectIdsScope(ids []uint32, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where(fmt.Sprintf("%s.object_id IN ?", tableName), ids)
		}
		return db
	}
}

func ObjectNamesScope(names []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where(fmt.Sprintf("%s.object_name IN ?", tableName), names)
		}
		return db
	}
}

func DepartmentalOrganizationIdsScope(ids []uint32, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where(fmt.Sprintf("%s.departmental_organization_id IN ?", tableName), ids)
		}
		return db
	}
}

func SportsAreaNamesScope(names []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where(fmt.Sprintf("%s.sports_area_name IN ?", tableName), names)
		}
		return db
	}
}

func SportsAreaTypeScope(types []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(types) > 0 {
			return db.Where(fmt.Sprintf("%s.sports_area_type IN ?", tableName), types)
		}
		return db
	}
}

func DepartmentalOrganizationNamesScope(names []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where(fmt.Sprintf("%s.departmental_organization_name IN ?", tableName), names)
		}
		return db
	}
}

func AvailabilitiesScope(availabilities []pb.Availability, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(availabilities) > 0 {
			return db.Where(fmt.Sprintf("%s.availability IN ?", tableName), availabilities)
		}
		return db
	}
}

func SportKindsScope(kinds []string, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(kinds) > 0 {
			return db.Where(fmt.Sprintf("%s.sport_kind IN ?", tableName), kinds)
		}
		return db
	}
}

func PolygonScope(polygon *pb.Polygon) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if polygon != nil {
			polyQuery := analytics.FormPolygonContainsQuery(polygon)
			return db.Where(polyQuery)
		}
		return db
	}
}

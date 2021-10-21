package filters

import (
	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"gorm.io/gorm"
)

func ObjectIdsScope(ids []uint32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where("object_id IN ?", ids)
		}
		return db
	}
}

func ObjectNamesScope(names []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where("object_name IN ?", names)
		}
		return db
	}
}

func DepartmentalOrganizationIdsScope(ids []uint32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) > 0 {
			return db.Where("departmental_organization_id IN ?", ids)
		}
		return db
	}
}

func SportsAreaNamesScope(names []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where("sports_area_name IN ?", names)
		}
		return db
	}
}

func SportsAreaTypeScope(types []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(types) > 0 {
			return db.Where("sports_area_type IN ?", types)
		}
		return db
	}
}

func DepartmentalOrganizationNamesScope(names []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(names) > 0 {
			return db.Where("departmental_organization_name IN ?", names)
		}
		return db
	}
}

func AvailabilitiesScope(availabilities []pb.Availability) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(availabilities) > 0 {
			return db.Where("availability IN ?", availabilities)
		}
		return db
	}
}

func SportKindsScope(kinds []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(kinds) > 0 {
			return db.Where("sport_kind IN ?", kinds)
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

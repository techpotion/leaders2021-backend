package filters

import (
	"github.com/techpotion/leaders2021-backend/gen/pb"
	"gorm.io/gorm"
)

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

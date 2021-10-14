package pb

import (
	"context"
)

type PointORM struct {
	Lat float32
	Lng float32
}

// TableName overrides the default tablename generated by GORM
func (PointORM) TableName() string {
	return "points"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Point) ToORM(ctx context.Context) (PointORM, error) {
	to := PointORM{}
	var err error
	if prehook, ok := interface{}(m).(PointWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Lat = m.Lat
	to.Lng = m.Lng
	if posthook, ok := interface{}(m).(PointWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *PointORM) ToPB(ctx context.Context) (Point, error) {
	to := Point{}
	var err error
	if prehook, ok := interface{}(m).(PointWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Lat = m.Lat
	to.Lng = m.Lng
	if posthook, ok := interface{}(m).(PointWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Point the arg will be the target, the caller the one being converted from

// PointWithBeforeToORM called before default ToORM code
type PointWithBeforeToORM interface {
	BeforeToORM(context.Context, *PointORM) error
}

// PointWithAfterToORM called after default ToORM code
type PointWithAfterToORM interface {
	AfterToORM(context.Context, *PointORM) error
}

// PointWithBeforeToPB called before default ToPB code
type PointWithBeforeToPB interface {
	BeforeToPB(context.Context, *Point) error
}

// PointWithAfterToPB called after default ToPB code
type PointWithAfterToPB interface {
	AfterToPB(context.Context, *Point) error
}

type SportsObjectORM struct {
	ObjectId                     uint32
	ObjectName                   string
	ObjectAddress                string
	DepartmentalOrganizationId   uint32
	DepartmentalOrganizationName string
	Availability                 uint32
	SportKind                    string
	ObjectSumSquare              float64
	ObjectPoint                  *PointORM `gorm:"not null;embedded;embeddedPrefix:object_point_"`
}

// TableName overrides the default tablename generated by GORM
func (SportsObjectORM) TableName() string {
	return "objects"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SportsObject) ToORM(ctx context.Context) (SportsObjectORM, error) {
	to := SportsObjectORM{}
	var err error
	if prehook, ok := interface{}(m).(SportsObjectWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.ObjectId = m.ObjectId
	to.ObjectName = m.ObjectName
	to.ObjectAddress = m.ObjectAddress
	if m.GetObjectPoint() != nil {
		tempObjectPoint, err := m.GetObjectPoint().ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.DepartmentalOrganizationId = m.DepartmentalOrganizationId
	to.DepartmentalOrganizationName = m.DepartmentalOrganizationName
	to.Availability = m.Availability
	to.SportKind = m.SportKind
	to.ObjectSumSquare = m.ObjectSumSquare
	if posthook, ok := interface{}(m).(SportsObjectWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SportsObjectORM) ToPB(ctx context.Context) (SportsObject, error) {
	to := SportsObject{}
	var err error
	if prehook, ok := interface{}(m).(SportsObjectWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.ObjectId = m.ObjectId
	to.ObjectName = m.ObjectName
	to.ObjectAddress = m.ObjectAddress
	if m.ObjectPoint != nil {
		tempObjectPoint, err := m.ObjectPoint.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.DepartmentalOrganizationId = m.DepartmentalOrganizationId
	to.DepartmentalOrganizationName = m.DepartmentalOrganizationName
	to.Availability = m.Availability
	to.SportKind = m.SportKind
	to.ObjectSumSquare = m.ObjectSumSquare
	if posthook, ok := interface{}(m).(SportsObjectWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SportsObject the arg will be the target, the caller the one being converted from

// SportsObjectWithBeforeToORM called before default ToORM code
type SportsObjectWithBeforeToORM interface {
	BeforeToORM(context.Context, *SportsObjectORM) error
}

// SportsObjectWithAfterToORM called after default ToORM code
type SportsObjectWithAfterToORM interface {
	AfterToORM(context.Context, *SportsObjectORM) error
}

// SportsObjectWithBeforeToPB called before default ToPB code
type SportsObjectWithBeforeToPB interface {
	BeforeToPB(context.Context, *SportsObject) error
}

// SportsObjectWithAfterToPB called after default ToPB code
type SportsObjectWithAfterToPB interface {
	AfterToPB(context.Context, *SportsObject) error
}

type SportsObjectDetailedORM struct {
	ObjectId                     uint32
	ObjectName                   string
	SportsAreaAddress            string
	DepartmentalOrganizationId   uint32
	DepartmentalOrganizationName string
	SportsAreaId                 uint32
	SportsAreaName               string
	SportsAreaType               string
	SportsAreaSquare             float64
	Availability                 uint32
	SportKind                    string
	ObjectPoint                  *PointORM `gorm:"not null;embedded;embeddedPrefix:object_point_"`
}

// TableName overrides the default tablename generated by GORM
func (SportsObjectDetailedORM) TableName() string {
	return "objects_detailed"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SportsObjectDetailed) ToORM(ctx context.Context) (SportsObjectDetailedORM, error) {
	to := SportsObjectDetailedORM{}
	var err error
	if prehook, ok := interface{}(m).(SportsObjectDetailedWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.ObjectId = m.ObjectId
	to.ObjectName = m.ObjectName
	to.SportsAreaAddress = m.SportsAreaAddress
	if m.GetObjectPoint() != nil {
		tempObjectPoint, err := m.GetObjectPoint().ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.DepartmentalOrganizationId = m.DepartmentalOrganizationId
	to.DepartmentalOrganizationName = m.DepartmentalOrganizationName
	to.SportsAreaId = m.SportsAreaId
	to.SportsAreaName = m.SportsAreaName
	to.SportsAreaType = m.SportsAreaType
	to.SportsAreaSquare = m.SportsAreaSquare
	to.Availability = m.Availability
	to.SportKind = m.SportKind
	if posthook, ok := interface{}(m).(SportsObjectDetailedWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SportsObjectDetailedORM) ToPB(ctx context.Context) (SportsObjectDetailed, error) {
	to := SportsObjectDetailed{}
	var err error
	if prehook, ok := interface{}(m).(SportsObjectDetailedWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.ObjectId = m.ObjectId
	to.ObjectName = m.ObjectName
	to.SportsAreaAddress = m.SportsAreaAddress
	if m.ObjectPoint != nil {
		tempObjectPoint, err := m.ObjectPoint.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.DepartmentalOrganizationId = m.DepartmentalOrganizationId
	to.DepartmentalOrganizationName = m.DepartmentalOrganizationName
	to.SportsAreaId = m.SportsAreaId
	to.SportsAreaName = m.SportsAreaName
	to.SportsAreaType = m.SportsAreaType
	to.SportsAreaSquare = m.SportsAreaSquare
	to.Availability = m.Availability
	to.SportKind = m.SportKind
	if posthook, ok := interface{}(m).(SportsObjectDetailedWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SportsObjectDetailed the arg will be the target, the caller the one being converted from

// SportsObjectDetailedWithBeforeToORM called before default ToORM code
type SportsObjectDetailedWithBeforeToORM interface {
	BeforeToORM(context.Context, *SportsObjectDetailedORM) error
}

// SportsObjectDetailedWithAfterToORM called after default ToORM code
type SportsObjectDetailedWithAfterToORM interface {
	AfterToORM(context.Context, *SportsObjectDetailedORM) error
}

// SportsObjectDetailedWithBeforeToPB called before default ToPB code
type SportsObjectDetailedWithBeforeToPB interface {
	BeforeToPB(context.Context, *SportsObjectDetailed) error
}

// SportsObjectDetailedWithAfterToPB called after default ToPB code
type SportsObjectDetailedWithAfterToPB interface {
	AfterToPB(context.Context, *SportsObjectDetailed) error
}

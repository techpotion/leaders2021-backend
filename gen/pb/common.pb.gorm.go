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

type ParkORM struct {
	CommonName     string
	AdmArea        string
	District       string
	Location       string
	HasSportground bool
	Square         float64
	ObjectPoint    *PointORM `gorm:"not null;embedded;embeddedPrefix:center_point_"`
}

// TableName overrides the default tablename generated by GORM
func (ParkORM) TableName() string {
	return "parks"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Park) ToORM(ctx context.Context) (ParkORM, error) {
	to := ParkORM{}
	var err error
	if prehook, ok := interface{}(m).(ParkWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.CommonName = m.CommonName
	to.AdmArea = m.AdmArea
	to.District = m.District
	to.Location = m.Location
	to.HasSportground = m.HasSportground
	if m.GetObjectPoint() != nil {
		tempObjectPoint, err := m.GetObjectPoint().ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.Square = m.Square
	if posthook, ok := interface{}(m).(ParkWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ParkORM) ToPB(ctx context.Context) (Park, error) {
	to := Park{}
	var err error
	if prehook, ok := interface{}(m).(ParkWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.CommonName = m.CommonName
	to.AdmArea = m.AdmArea
	to.District = m.District
	to.Location = m.Location
	to.HasSportground = m.HasSportground
	if m.ObjectPoint != nil {
		tempObjectPoint, err := m.ObjectPoint.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.Square = m.Square
	if posthook, ok := interface{}(m).(ParkWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Park the arg will be the target, the caller the one being converted from

// ParkWithBeforeToORM called before default ToORM code
type ParkWithBeforeToORM interface {
	BeforeToORM(context.Context, *ParkORM) error
}

// ParkWithAfterToORM called after default ToORM code
type ParkWithAfterToORM interface {
	AfterToORM(context.Context, *ParkORM) error
}

// ParkWithBeforeToPB called before default ToPB code
type ParkWithBeforeToPB interface {
	BeforeToPB(context.Context, *Park) error
}

// ParkWithAfterToPB called after default ToPB code
type ParkWithAfterToPB interface {
	AfterToPB(context.Context, *Park) error
}

type PollutionORM struct {
	AdmArea     string
	District    string
	Location    string
	IsPolluted  bool
	Results     string
	ObjectPoint *PointORM `gorm:"not null;embedded;embeddedPrefix:center_point_"`
}

// TableName overrides the default tablename generated by GORM
func (PollutionORM) TableName() string {
	return "pollutions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Pollution) ToORM(ctx context.Context) (PollutionORM, error) {
	to := PollutionORM{}
	var err error
	if prehook, ok := interface{}(m).(PollutionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AdmArea = m.AdmArea
	to.District = m.District
	to.Location = m.Location
	to.IsPolluted = m.IsPolluted
	if m.GetObjectPoint() != nil {
		tempObjectPoint, err := m.GetObjectPoint().ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.Results = m.Results
	if posthook, ok := interface{}(m).(PollutionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *PollutionORM) ToPB(ctx context.Context) (Pollution, error) {
	to := Pollution{}
	var err error
	if prehook, ok := interface{}(m).(PollutionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AdmArea = m.AdmArea
	to.District = m.District
	to.Location = m.Location
	to.IsPolluted = m.IsPolluted
	if m.ObjectPoint != nil {
		tempObjectPoint, err := m.ObjectPoint.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.ObjectPoint = &tempObjectPoint
	}
	to.Results = m.Results
	if posthook, ok := interface{}(m).(PollutionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Pollution the arg will be the target, the caller the one being converted from

// PollutionWithBeforeToORM called before default ToORM code
type PollutionWithBeforeToORM interface {
	BeforeToORM(context.Context, *PollutionORM) error
}

// PollutionWithAfterToORM called after default ToORM code
type PollutionWithAfterToORM interface {
	AfterToORM(context.Context, *PollutionORM) error
}

// PollutionWithBeforeToPB called before default ToPB code
type PollutionWithBeforeToPB interface {
	BeforeToPB(context.Context, *Pollution) error
}

// PollutionWithAfterToPB called after default ToPB code
type PollutionWithAfterToPB interface {
	AfterToPB(context.Context, *Pollution) error
}

type CirclesClusterORM struct {
	Geojson string
	Square  float64
}

// TableName overrides the default tablename generated by GORM
func (CirclesClusterORM) TableName() string {
	return "circles_clusters"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *CirclesCluster) ToORM(ctx context.Context) (CirclesClusterORM, error) {
	to := CirclesClusterORM{}
	var err error
	if prehook, ok := interface{}(m).(CirclesClusterWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Geojson = m.Geojson
	to.Square = m.Square
	if posthook, ok := interface{}(m).(CirclesClusterWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CirclesClusterORM) ToPB(ctx context.Context) (CirclesCluster, error) {
	to := CirclesCluster{}
	var err error
	if prehook, ok := interface{}(m).(CirclesClusterWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Geojson = m.Geojson
	to.Square = m.Square
	if posthook, ok := interface{}(m).(CirclesClusterWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type CirclesCluster the arg will be the target, the caller the one being converted from

// CirclesClusterWithBeforeToORM called before default ToORM code
type CirclesClusterWithBeforeToORM interface {
	BeforeToORM(context.Context, *CirclesClusterORM) error
}

// CirclesClusterWithAfterToORM called after default ToORM code
type CirclesClusterWithAfterToORM interface {
	AfterToORM(context.Context, *CirclesClusterORM) error
}

// CirclesClusterWithBeforeToPB called before default ToPB code
type CirclesClusterWithBeforeToPB interface {
	BeforeToPB(context.Context, *CirclesCluster) error
}

// CirclesClusterWithAfterToPB called after default ToPB code
type CirclesClusterWithAfterToPB interface {
	AfterToPB(context.Context, *CirclesCluster) error
}

type RegionORM struct {
	Region     string
	Square     float64
	Density    float64
	Population float64
}

// TableName overrides the default tablename generated by GORM
func (RegionORM) TableName() string {
	return "regions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Region) ToORM(ctx context.Context) (RegionORM, error) {
	to := RegionORM{}
	var err error
	if prehook, ok := interface{}(m).(RegionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Region = m.Region
	to.Square = m.Square
	to.Density = m.Density
	to.Population = m.Population
	if posthook, ok := interface{}(m).(RegionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *RegionORM) ToPB(ctx context.Context) (Region, error) {
	to := Region{}
	var err error
	if prehook, ok := interface{}(m).(RegionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Region = m.Region
	to.Square = m.Square
	to.Density = m.Density
	to.Population = m.Population
	if posthook, ok := interface{}(m).(RegionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Region the arg will be the target, the caller the one being converted from

// RegionWithBeforeToORM called before default ToORM code
type RegionWithBeforeToORM interface {
	BeforeToORM(context.Context, *RegionORM) error
}

// RegionWithAfterToORM called after default ToORM code
type RegionWithAfterToORM interface {
	AfterToORM(context.Context, *RegionORM) error
}

// RegionWithBeforeToPB called before default ToPB code
type RegionWithBeforeToPB interface {
	BeforeToPB(context.Context, *Region) error
}

// RegionWithAfterToPB called after default ToPB code
type RegionWithAfterToPB interface {
	AfterToPB(context.Context, *Region) error
}

type SubwayORM struct {
	Name                string
	LineColor           string
	DistanceFromPolygon float64
	Point               *PointORM `gorm:"not null;embedded;embeddedPrefix:point_"`
}

// TableName overrides the default tablename generated by GORM
func (SubwayORM) TableName() string {
	return "subway_stations"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Subway) ToORM(ctx context.Context) (SubwayORM, error) {
	to := SubwayORM{}
	var err error
	if prehook, ok := interface{}(m).(SubwayWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Name = m.Name
	to.LineColor = m.LineColor
	if m.GetPoint() != nil {
		tempPoint, err := m.GetPoint().ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Point = &tempPoint
	}
	to.DistanceFromPolygon = m.DistanceFromPolygon
	if posthook, ok := interface{}(m).(SubwayWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SubwayORM) ToPB(ctx context.Context) (Subway, error) {
	to := Subway{}
	var err error
	if prehook, ok := interface{}(m).(SubwayWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Name = m.Name
	to.LineColor = m.LineColor
	if m.Point != nil {
		tempPoint, err := m.Point.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Point = &tempPoint
	}
	to.DistanceFromPolygon = m.DistanceFromPolygon
	if posthook, ok := interface{}(m).(SubwayWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Subway the arg will be the target, the caller the one being converted from

// SubwayWithBeforeToORM called before default ToORM code
type SubwayWithBeforeToORM interface {
	BeforeToORM(context.Context, *SubwayORM) error
}

// SubwayWithAfterToORM called after default ToORM code
type SubwayWithAfterToORM interface {
	AfterToORM(context.Context, *SubwayORM) error
}

// SubwayWithBeforeToPB called before default ToPB code
type SubwayWithBeforeToPB interface {
	BeforeToPB(context.Context, *Subway) error
}

// SubwayWithAfterToPB called after default ToPB code
type SubwayWithAfterToPB interface {
	AfterToPB(context.Context, *Subway) error
}

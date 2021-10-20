package circles

import (
	"github.com/lib/pq"
	"github.com/techpotion/leaders2021-backend/gen/pb"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CircleIntersectionsORM struct {
	gorm.Model
	ObjectId      uint32
	Intersections pq.Int32Array `gorm:"type:bigint[]"`
	Square        float64
}

func (c *CircleIntersectionsORM) ToPB(ctx context.Context) (pb.CircleIntersections, error) {
	return pb.CircleIntersections{
		ObjectId:      c.ObjectId,
		Intersections: c.Intersections,
		Square:        c.Square,
	}, nil
}

package server

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/controllers/sports_objects_detailed"
)

// implementing api service interface generated from proto
type ApiServiceServer struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) ListSportsObjectsDetailed(ctx context.Context, in *pb.SportsObjectsDetailed_ListRequest) (*pb.SportsObjectsDetailed_ListResponse, error) {
	return sports_objects_detailed.List(ctx, in)
}

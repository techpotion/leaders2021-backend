package server

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/controllers/sports_objects"
)

// implementing api service interface generated from proto
type ApiServiceServer struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) SportsObjects(ctx context.Context, in *pb.SportsObjects_ListRequest) (*pb.SportsObjects_ListResponse, error) {
	return sports_objects.List(ctx, in)
}

package server

import (
	"context"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	sportsobjects "github.com/techpotion/leaders2021-backend/modules/controllers/sports_objects"
	sportsobjectsdetailed "github.com/techpotion/leaders2021-backend/modules/controllers/sports_objects_detailed"
)

// implementing api service interface generated from proto
type ApiServiceServer struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) ListSportsObjects(ctx context.Context, in *pb.SportsObjects_ListRequest) (*pb.SportsObjects_ListResponse, error) {
	return sportsobjects.List(ctx, in)
}

func (s *ApiServiceServer) ListSportsObjectsDetailed(ctx context.Context, in *pb.SportsObjectsDetailed_ListRequest) (*pb.SportsObjectsDetailed_ListResponse, error) {
	return sportsobjectsdetailed.List(ctx, in)
}

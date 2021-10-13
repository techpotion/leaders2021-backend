package analytics

import (
	"context"
	"fmt"

	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/analytics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetPolygonAnalytics(ctx context.Context, in *pb.PolygonAnalytics_GetRequest) (*pb.PolygonAnalytics_GetResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := analytics.ValidatePolygon(in.Polygon); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	aksdjh := analytics.FormPolygonContainsQuery(in.Polygon)
	fmt.Println(aksdjh)

	return nil, nil
}

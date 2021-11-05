package exports

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/gen/pb"
	"github.com/techpotion/leaders2021-backend/modules/utils"
	"google.golang.org/grpc"
)

var (
	client pb.AnalyticsExportServiceClient
)

func Init() {
	var err error

	host := ""
	if host, err = utils.GetEnvStr("EXPORTS_HOST"); err != nil {
		host = viper.GetString("microservices.exports.host")
	}

	port := 0
	if port, err = utils.GetEnvInt("EXPORTS_PORT"); err != nil {
		port = viper.GetInt("microservices.exports.port")
	}

	target := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(target)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Panicln("failed to connect to analytics microservice")
	}

	client = pb.NewAnalyticsExportServiceClient(conn)
}

func Get(ctx context.Context, in *pb.PolygonAnalyticsDashboard_Response) (*pb.Exports_GetResponse, error) {
	return client.GetExport(ctx, in)
}

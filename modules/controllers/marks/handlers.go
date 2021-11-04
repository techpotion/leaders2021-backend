package marks

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
	client pb.MicroserviceClient
)

func Init() {
	var err error

	host := ""
	if host, err = utils.GetEnvStr("MARKS_HOST"); err != nil {
		host = viper.GetString("microservices.marks.host")
	}

	port := 0
	if port, err = utils.GetEnvInt("MARKS_PORT"); err != nil {
		port = viper.GetInt("microservices.marks.port")
	}

	target := fmt.Sprintf("%s:%d", host, port)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Panicln("failed to connect to marks microservice")
	}

	client = pb.NewMicroserviceClient(conn)
}

func GetMark(ctx context.Context, in *pb.Marks_GetRequest) (*pb.Marks_GetResponse, error) {
	return client.GetMark(ctx, in)
}

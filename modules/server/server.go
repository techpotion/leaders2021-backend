package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/techpotion/leaders2021-backend/gen/pb"
	"google.golang.org/grpc"
)

// Run starts new server
func Run() {
	grpcInterface := viper.GetString("server.grpc.interface")
	grpcPort := viper.GetInt("server.grpc.port")
	grpcConnectionString := fmt.Sprintf("%s:%d", grpcInterface, grpcPort)

	logrus.Infoln("serving gRPC on " + grpcConnectionString)
	lis, err := net.Listen("tcp", grpcConnectionString)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln("failed to listen")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterApiServiceServer(grpcServer, &ApiServiceServer{})

	// starting new grpc server in a goroutine
	go func() {
		if grpcServer.Serve(lis) != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Fatalln("failed to start server")
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		grpcConnectionString,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln("failed to dial server")
	}

	httpInterface := viper.GetString("server.http.interface")
	httpPort := viper.GetInt("server.http.port")
	httpConnectionString := fmt.Sprintf("%s:%d", httpInterface, httpPort)

	gwMux := runtime.NewServeMux()
	err = pb.RegisterApiServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln("failed to register gateway")
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)

	mux.HandleFunc("/swagger.json", serveSwagger)

	logrus.Infoln("serving HTTP gRPC-Gateway on " + httpConnectionString)
	if err := http.ListenAndServe(httpConnectionString, allowCORS(mux)); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln("gateway listener failed")
	}
}

// allowCORS sets up cors settings
func allowCORS(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"*"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(h)
}

package main

import (
	"go-microservice-example/api"
	"go-microservice-example/internal/audit"
	"go-microservice-example/internal/database"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
}

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	database.Migrate()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		zap.S().Fatal(err)
	}

	srv := grpc.NewServer()

	api.RegisterAuditServiceServer(srv, &audit.Server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		zap.S().Fatal(err)
	}
}

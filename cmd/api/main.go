package main

import (
	"context"
	"go-microservice-example/internal/database"
	"go-microservice-example/pkg/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/spf13/viper"
)

type server struct {
}

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
	api.RegisterAuditServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		zap.S().Fatal(err)
	}
}

func (s *server) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {
	zap.S().Info(
		request.GetUserId(),
		request.GetOrgId(),
		request.GetIpAddress(),
		request.GetTarget(),
		request.GetAction(),
		request.GetActionType(),
		request.GetEventName(),
	)

	return &api.CreateResponse{Response: "HELLO RESPONSE"}, nil
}

func (s *server) ReadUser(context.Context, *api.ReadUserRequest) (*api.ReadUserResponse, error) {
	panic("implement me")
}

func (s *server) ReadOrg(context.Context, *api.ReadOrgRequest) (*api.ReadOrgResponse, error) {
	panic("implement me")
}

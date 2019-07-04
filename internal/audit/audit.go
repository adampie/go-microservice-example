package audit

import (
	"context"
	"go-microservice-example/pkg/api"

	"go.uber.org/zap"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {
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

func (s *Server) ReadUser(context.Context, *api.ReadUserRequest) (*api.ReadUserResponse, error) {
	panic("implement me")
}

func (s *Server) ReadOrg(context.Context, *api.ReadOrgRequest) (*api.ReadOrgResponse, error) {
	panic("implement me")
}

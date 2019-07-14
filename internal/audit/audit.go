package audit

import (
	"context"
	"go-microservice-example/api"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	req := Create{
		id:         uuid.New(),
		userId:     request.GetUserId(),
		orgId:      request.GetOrgId(),
		ipAddress:  request.GetIpAddress(),
		target:     request.GetTarget(),
		action:     request.GetAction(),
		actionType: request.GetActionType(),
		eventName:  request.GetEventName(),
	}

	zap.S().Info(req)

	return &api.CreateResponse{Response: "HELLO RESPONSE"}, nil
}

func (s *Server) ReadUser(context.Context, *api.ReadUserRequest) (*api.ReadUserResponse, error) {
	panic("implement me")
}

func (s *Server) ReadOrg(context.Context, *api.ReadOrgRequest) (*api.ReadOrgResponse, error) {
	panic("implement me")
}

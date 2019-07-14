package audit

import (
	"context"
	"go-microservice-example/api"
	"go-microservice-example/internal/db"
	"go-microservice-example/internal/model"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	req := model.Create{
		Id:         uuid.New(),
		UserId:     request.GetUserId(),
		OrgId:      request.GetOrgId(),
		IpAddress:  request.GetIpAddress(),
		Target:     request.GetTarget(),
		Action:     request.GetAction(),
		ActionType: request.GetActionType(),
		EventName:  request.GetEventName(),
	}

	zap.S().Info(req)

	db.CreateAudit(req)

	return &api.CreateResponse{Response: "HELLO RESPONSE"}, nil
}

func (s *Server) ReadUser(context.Context, *api.ReadUserRequest) (*api.ReadUserResponse, error) {
	panic("implement me")
}

func (s *Server) ReadOrg(context.Context, *api.ReadOrgRequest) (*api.ReadOrgResponse, error) {
	panic("implement me")
}

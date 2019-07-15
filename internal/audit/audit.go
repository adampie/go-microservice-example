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

	return &api.CreateResponse{Response: "CREATE RESPONSE"}, nil
}

func (s *Server) ReadUser(ctx context.Context, request *api.ReadUserRequest) (*api.ReadUserResponse, error) {

	req := model.User{
		UserId: request.GetUserId(),
	}

	zap.S().Info(req)

	db.ReadUser(req)

	return &api.ReadUserResponse{Response: "READ USER RESPONSE"}, nil
}

func (s *Server) ReadOrg(ctx context.Context, request *api.ReadOrgRequest) (*api.ReadOrgResponse, error) {

	req := model.Org{
		OrgId: request.GetOrgId(),
	}

	zap.S().Info(req)

	db.ReadOrg(req)

	return &api.ReadOrgResponse{Response: "READ ORG RESPONSE"}, nil
}

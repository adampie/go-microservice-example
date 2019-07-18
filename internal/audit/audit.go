package audit

import (
	"context"
	"go-microservice-example/api"
	"go-microservice-example/internal/db"
	"go-microservice-example/internal/model"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

// Server gRPC struct
type Server struct {
}

// Create a audit log
func (s *Server) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	req := model.Create{
		ID:         uuid.New(),
		UserID:     request.GetUserId(),
		OrgID:      request.GetOrgId(),
		IPAddress:  request.GetIpAddress(),
		Target:     request.GetTarget(),
		Action:     request.GetAction(),
		ActionType: request.GetActionType(),
		EventName:  request.GetEventName(),
	}

	zap.S().Info(req)

	dbc := db.GetDB()

	db.CreateAudit(dbc, req)

	return &api.CreateResponse{Response: "CREATE RESPONSE"}, nil
}

// ReadUser gets all the audit relevant to the user
func (s *Server) ReadUser(ctx context.Context, request *api.ReadUserRequest) (*api.ReadUserResponse, error) {

	req := model.User{
		UserID: request.GetUserId(),
	}

	zap.S().Info(req)

	dbc := db.GetDB()

	db.ReadUser(dbc, req)

	return &api.ReadUserResponse{Response: "READ USER RESPONSE"}, nil
}

// ReadOrg gets all the audit relevant to the organisation
func (s *Server) ReadOrg(ctx context.Context, request *api.ReadOrgRequest) (*api.ReadOrgResponse, error) {

	req := model.Org{
		OrgID: request.GetOrgId(),
	}

	zap.S().Info(req)

	dbc := db.GetDB()

	db.ReadOrg(dbc, req)

	return &api.ReadOrgResponse{Response: "READ ORG RESPONSE"}, nil
}

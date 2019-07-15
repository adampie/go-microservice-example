package audit

import (
	"context"
	"go-microservice-example/api"
	"reflect"
	"testing"
)

func TestServer_Create(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *api.CreateRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *api.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Create(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ReadUser(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *api.ReadUserRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *api.ReadUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ReadUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ReadUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ReadUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ReadOrg(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *api.ReadOrgRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *api.ReadOrgResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ReadOrg(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ReadOrg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ReadOrg() = %v, want %v", got, tt.want)
			}
		})
	}
}

package db

import (
	"go-microservice-example/internal/model"
	"testing"
)

func TestCreateAudit(t *testing.T) {
	type args struct {
		c model.Create
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAudit(tt.args.c)
		})
	}
}

func TestReadUser(t *testing.T) {
	type args struct {
		u model.User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadUser(tt.args.u)
		})
	}
}

func TestReadOrg(t *testing.T) {
	type args struct {
		o model.Org
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadOrg(tt.args.o)
		})
	}
}

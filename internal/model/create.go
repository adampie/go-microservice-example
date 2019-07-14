package model

import (
	"github.com/google/uuid"
)

type Create struct {
	Id         uuid.UUID
	UserId     string
	OrgId      string
	IpAddress  string
	Target     string
	Action     string
	ActionType string
	EventName  string
}

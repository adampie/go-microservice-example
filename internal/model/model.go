package model

import (
	"github.com/google/uuid"
)

// Create audit struct
type Create struct {
	ID         uuid.UUID
	UserID     string
	OrgID      string
	IPAddress  string
	Target     string
	Action     string
	ActionType string
	EventName  string
}

// Org struct
type Org struct {
	OrgID string
}

// User struct
type User struct {
	UserID string
}

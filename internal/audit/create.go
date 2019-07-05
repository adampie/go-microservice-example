package audit

import (
	"github.com/google/uuid"
)

type Create struct {
	id         uuid.UUID
	userId     string
	orgId      string
	ipAddress  string
	target     string
	action     string
	actionType string
	eventName  string
}

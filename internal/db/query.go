package db

import (
	"go-microservice-example/internal/model"

	"go.uber.org/zap"
)

// CreateAudit create a new audit request
func CreateAudit(c model.Create) {
	db := GetDB()

	id := c.ID
	userID := c.UserID
	orgID := c.OrgID
	ipAddress := c.IPAddress
	target := c.Target
	action := c.Action
	actionType := c.ActionType
	eventName := c.EventName

	_, err := db.Exec("INSERT INTO audit VALUES($1, $2, $3, $4, $5, $6, $7, $8)", id, userID, orgID, ipAddress, target, action, actionType, eventName)
	if err != nil {
		zap.S().Error(err)
	}
}

// ReadUser returns all the audit data relevant to the user
func ReadUser(u model.User) {
	db := GetDB()

	userID := u.UserID

	q, err := db.Exec("SELECT * FROM audit WHERE user_id == $1", userID)
	if err != nil {
		zap.S().Error(err)
	}

	zap.S().Info(q)
}

// ReadOrg returns all the audit data relevant to the Org
func ReadOrg(o model.Org) {
	db := GetDB()

	orgID := o.OrgID

	q, err := db.Exec("SELECT * FROM audit WHERE org_id == $1", orgID)
	if err != nil {
		zap.S().Error(err)
	}

	zap.S().Info(q)
}

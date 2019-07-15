package db

import (
	"go-microservice-example/internal/model"

	"go.uber.org/zap"
)

func CreateAudit(c model.Create) {
	db := GetDB()

	id := c.Id
	userId := c.UserId
	orgId := c.OrgId
	ipAddress := c.IpAddress
	target := c.Target
	action := c.Action
	actionType := c.ActionType
	eventName := c.EventName

	_, err := db.Exec("INSERT INTO audit VALUES($1, $2, $3, $4, $5, $6, $7, $8)", id, userId, orgId, ipAddress, target, action, actionType, eventName)
	if err != nil {
		zap.S().Error(err)
	}
}

func ReadUser(u model.User) {
	db := GetDB()

	userId := u.UserId

	q, err := db.Exec("SELECT * FROM audit WHERE user_id == $1", id, userId)
	if err != nil {
		zap.S().Error(err)
	}

	zap.S().Info(q)
}

func ReadOrg(o model.Org) {
	db := GetDB()

	orgId := o.OrgId

	q, err := db.Exec("SELECT * FROM audit WHERE org_id == $1", orgId)
	if err != nil {
		zap.S().Error(err)
	}

	zap.S().Info(q)
}

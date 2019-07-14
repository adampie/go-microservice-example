package db

import (
	"go-microservice-example/internal/model"

	"go.uber.org/zap"
)

func CreateAudit(c model.Create) {
	db := GetDB()

	id := c.Id
	userId := c.UserId
	orderId := c.OrgId
	ipAddress := c.IpAddress
	target := c.Target
	action := c.Action
	actionType := c.ActionType
	eventName := c.EventName

	_, err := db.Exec("INSERT INTO audit VALUES($1, $2, $3, $4, $5, $6, $7, $8)", id, userId, orderId, ipAddress, target, action, actionType, eventName)
	if err != nil {
		zap.S().Error(err)
	}
}

package main

import (
	"go.uber.org/zap"
	"os"
)

func main() {

	if os.Getenv("ENV") == "PRODUCTION" {
		logger, _ := zap.NewProduction()
		zap.ReplaceGlobals(logger)
	} else {
		logger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
	}

}

package services

import (
	"go.uber.org/zap"
)

func initLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return sugar
}

var Logger = initLogger()

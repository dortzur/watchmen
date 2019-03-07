package services

import (
	"fmt"
	"go.uber.org/zap"
)

func initLogger() *zap.SugaredLogger {
	y := 42
	fmt.Println(y)

	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	sugar.Info("hello")
	return sugar
}

var Logger = initLogger()

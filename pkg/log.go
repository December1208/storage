package pkg

import (
	"go.uber.org/zap"
)

func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

var Logger = GetLogger()

package app

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func InitAppLogger(config AppConfig) error {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zapLogger = zapLogger
	return nil
}

func GetAppLogger() *zap.Logger {
	return zapLogger
}

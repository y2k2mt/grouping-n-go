package app

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func InitLogger(config AppConfig) error {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zapLogger = zapLogger
	return nil
}

func GetZapLogger() *zap.Logger {
	return zapLogger
}

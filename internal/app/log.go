package app

import (
	"fmt"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func InitAppLogger(config AppConfig) error {
	var (
		logger *zap.Logger
		err    error
	)
	switch config.Env {
	case Development:
		logger, err = zap.NewDevelopment()
	case Staging:
		logger, err = zap.NewProduction()
	case Production:
		logger, err = zap.NewProduction()
	default:
		return fmt.Errorf("an unrecognized env %v", config.Env)
	}
	if err != nil {
		return err
	}
	zapLogger = logger
	return nil
}

func GetAppLogger() *zap.Logger {
	return zapLogger
}

package app

import (
	"go.uber.org/zap"
	"log"
)

var zapLogger *zap.Logger

func InitLogger(config Config) error {
	zapLogger, err := zap.NewProduction()
	if err != nil {
    return err
	}
	zapLogger = *zapLogger
}

func GetZapLogger() *zap.Logger {
  return zapLogger
}

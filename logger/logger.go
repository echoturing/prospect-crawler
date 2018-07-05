package logger

import (
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			print(err.Error())
			os.Exit(1)
		}
		return logger
	}
	return logger
}

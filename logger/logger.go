package logger

import (
	"os"

	"go.uber.org/zap"
)

var log *zap.Logger

func GetLogger() *zap.Logger {
	if log == nil {
		log, err := zap.NewProduction()
		if err != nil {
			print(err.Error())
			os.Exit(1)
		}
		return log
	}
	return log
}

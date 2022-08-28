package logger

import (
	"log"

	"go.uber.org/zap"
)

var Log *zap.Logger

func InitialLog() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("logger initial fail %v", err)
	}
	defer logger.Sync()
	Log = logger
}

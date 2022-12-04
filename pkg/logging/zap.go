package logging

import (
	"go.uber.org/zap"
	"os"
)

var log *zap.Logger

type Logger struct {
	*zap.Logger
}

func GetLogger() *Logger {
	return &Logger{
		log,
	}
}

func init() {

	newLogger, _ := zap.NewProduction(
		zap.ErrorOutput(os.Stdout),
		//zap.ErrorOutput(someFile),
		zap.Development(),
		zap.AddCaller(),
	)

	log = newLogger

}

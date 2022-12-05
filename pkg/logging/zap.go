package logging

import (
	"go.uber.org/zap"
	"os"
	"sync"
)

var once sync.Once
var log *zap.Logger

type Logger struct {
	*zap.Logger
}

func GetLogger() *Logger {

	once.Do(func() {
		log, _ = zap.NewProduction(
			zap.ErrorOutput(os.Stdout),
			//zap.ErrorOutput(someFile),
			zap.Development(),
			zap.AddCaller(),
		)
	})

	return &Logger{
		log,
	}

}

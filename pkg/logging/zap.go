package logging

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

type zapLogger struct {
	*zap.Logger
}

func (z zapLogger) Infof(s string, v ...any) {
	//TODO implement me
	panic("implement me")
}

func (z zapLogger) Errorf(s string, v ...any) {
	//TODO implement me
	panic("implement me")
}

func GetLogger() (*zapLogger, error) {

	log, err := zap.NewProduction(
		zap.ErrorOutput(os.Stdout),
		zap.ErrorOutput(os.Stderr),
		zap.Development(),
		zap.AddCaller(),
		zap.AddStacktrace(zap.InfoLevel),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Logger")
	}

	l := &zapLogger{log}

	log.Info("logger got successfully")
	return l, nil

}

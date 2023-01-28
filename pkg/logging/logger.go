package logging

type Logger interface {
	Infof(format string, v ...any)
	Errorf(format string, v ...any)
}

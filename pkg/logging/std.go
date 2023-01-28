package logging

import (
	"io"
	"log"
	"os"
)

type stdLogger struct {
	info  *log.Logger
	error *log.Logger
}

func (l stdLogger) Infof(format string, args ...any) {
	l.info.Printf(format+"\n", args...)
}

func (l stdLogger) Errorf(format string, args ...any) {
	l.info.Printf(format+"\n", args...)
	l.error.Printf(format+"\n", args...)
}

func GetStdLogger() Logger {

	// info level
	//err := os.MkdirAll("logs", 0644)
	//if err != nil {
	//	panic(err)
	//}

	//f, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	//if err != nil {
	//	panic(err)
	//}

	//f, err := os.Create("info.log")
	//if err != nil {
	//	panic(err)
	//}

	multiWriter := io.MultiWriter(os.Stdout)
	infoLevelLogger := log.New(multiWriter, "info:  ", log.LstdFlags)

	// error level
	errorLevelLogger := log.New(os.Stdout, "error: ", log.LstdFlags)

	return &stdLogger{
		info:  infoLevelLogger,
		error: errorLevelLogger,
	}
}

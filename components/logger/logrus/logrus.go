package logruslogger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"instagram/components/logger"
	"runtime"
	"strings"
)

type logrusLogger struct {
	*logrus.Entry
}

func NewLogrusLogger() logger.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	return &logrusLogger{logrus.NewEntry(log)}
}

func (l *logrusLogger) debugSrc() *logrus.Entry {

	if _, ok := l.Entry.Data["source"]; ok {
		return l.Entry
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	return l.Entry.WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.debugSrc().Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.debugSrc().Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.debugSrc().Warn(args...)
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.debugSrc().Print(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.debugSrc().Error(args...)
}

//
//func (l *logrusLogger) Fatal(args ...interface{}) {
//	l.debugSrc().Fatal(args...)
//}
//
//func (l *logrusLogger) Panic(args ...interface{}) {
//	l.debugSrc().Panic(args...)
//}

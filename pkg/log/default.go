package log

import "go.uber.org/zap"

var DefaultLogger *zap.Logger

func init() {
	DefaultLogger = newZap()
	//zap.ReplaceGlobals(DefaultLogger)
}

func Debug(msg string) {
	DefaultLogger.Debug(msg)
}

func Info(msg string) {
	DefaultLogger.Info(msg)
}

func Error(msg string) {
	DefaultLogger.Error(msg)
}

func Panic(msg string) {
	DefaultLogger.Panic(msg)
}

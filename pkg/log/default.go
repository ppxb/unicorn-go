package log

import "go.uber.org/zap"

var DefaultLogger *zap.Logger

func init() {
	DefaultLogger = newZap()
	//zap.ReplaceGlobals(DefaultLogger)
}

func Info(msg string) {
	DefaultLogger.Info(msg)
}

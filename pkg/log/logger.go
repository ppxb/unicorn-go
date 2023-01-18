package log

import (
	"gorm.io/gorm/logger"
	"regexp"
	"runtime"
)

var (
	logDir     = ""
	helperDier = ""
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	logDir = regexp.MustCompile(`logger\.go`).ReplaceAllString(file, "")
	helperDier = regexp.MustCompile(`pkg.log.logger\.go`).ReplaceAllString(file, "")
}

type Interface interface {
	Options() Options
	WithFields(fields map[string]interface{}) Interface
	Log(level Level, v ...interface{})
	Logf(level Level, format string, v ...interface{})
}

type Config struct {
	ops  Options
	gorm logger.Config
}

func New(options ...func(*Options)) Interface {
	ops := getOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	l := newZap(ops)
	return l
}

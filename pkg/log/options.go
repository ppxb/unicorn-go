package log

import (
	"go.uber.org/zap"
	"io"
)

type FileWithLineNumOptions struct {
	skipGorm   bool
	skipHelper bool
}

func WithSkipGorm(flag bool) func(*FileWithLineNumOptions) {
	return func(options *FileWithLineNumOptions) {
		getFileWithLineNumOptionsOrSetDefault(options).skipGorm = flag
	}
}

func WithSkipHelper(flag bool) func(*FileWithLineNumOptions) {
	return func(options *FileWithLineNumOptions) {
		getFileWithLineNumOptionsOrSetDefault(options).skipHelper = flag
	}
}

func getFileWithLineNumOptionsOrSetDefault(options *FileWithLineNumOptions) *FileWithLineNumOptions {
	if options == nil {
		return &FileWithLineNumOptions{}
	}
	return options
}

type Options struct {
	level          Level
	output         io.Writer
	json           bool
	lineNum        bool
	lineNumPrefix  string
	lineNumLevel   int
	lineNumSource  bool
	lineNumVersion bool
}

func getOptionsOrSetDefault(options *Options) *Options {
	if options == nil {
		return &Options{
			level:          Level(zap.DebugLevel),
			lineNum:        true,
			lineNumLevel:   1,
			lineNumVersion: true,
		}
	}
	return options
}

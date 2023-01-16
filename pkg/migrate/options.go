package migrate

import (
	"context"
	"embed"
	"github.com/ppxb/unicorn/pkg/utils"
)

type Options struct {
	ctx         context.Context
	driver      string
	uri         string
	lockName    string
	before      func(ctx context.Context) error
	changeTable string
	fs          embed.FS
	fsRoot      string
}

// WithContext WithCtx set datasource global context
func WithContext(ctx context.Context) func(*Options) {
	return func(options *Options) {
		if !utils.InterfaceIsNil(ctx) {
			getOptionsOrSetDefault(options).ctx = ctx
		}
	}
}

// WithDriver set datasource driver
func WithDriver(s string) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).driver = s
	}
}

// WithFs set database go embed file
func WithFs(fs embed.FS) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).fs = fs
	}
}

func WithFsRoot(s string) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).fsRoot = s
	}
}

// WithUri set datasource uri
func WithUri(s string) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).uri = s
	}
}

// WithBefore set database migrate before hook
func WithBefore(f func(ctx context.Context) error) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).before = f
	}
}

func getOptionsOrSetDefault(options *Options) *Options {
	if options == nil {
		return &Options{
			driver:      "mysql",
			lockName:    "MigrationLock",
			changeTable: "schema_migrations",
		}
	}
	return options
}

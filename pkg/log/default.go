package log

import "context"

var DefaultWrapper *Wrapper

func init() {
	DefaultWrapper = &Wrapper{
		log:    New(),
		fields: map[string]interface{}{},
	}
}

func NewDefaultWrapper() *Wrapper {
	return DefaultWrapper
}

func Trace(args ...interface{}) {
	DefaultWrapper.Trace(args...)
}

func WithContext(ctx context.Context) *Wrapper {
	return DefaultWrapper.WithContext(ctx)
}

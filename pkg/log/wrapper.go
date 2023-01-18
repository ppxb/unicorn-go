package log

import (
	"context"
	"fmt"
	"github.com/ppxb/unicorn/pkg/constant"
	"os"
)

type Wrapper struct {
	log    Interface
	fields map[string]interface{}
}

// NewWrapper gen new log wrapper
func NewWrapper(l Interface) *Wrapper {
	return &Wrapper{
		log:    l,
		fields: map[string]interface{}{},
	}
}

func (w *Wrapper) Trace(args ...interface{}) {
	if !w.log.Options().level.Enabled(TraceLevel) {
		return
	}
	ns := copyFields(w.fields)
	//show log's line number
	if w.log.Options().lineNum {
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(DebugLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(TraceLevel, args...)
}

func (w *Wrapper) Debug(args ...interface{}) {
	if !w.log.Options().level.Enabled(DebugLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(DebugLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(DebugLevel, args...)
}

func (w *Wrapper) Info(args ...interface{}) {
	fmt.Println(args)
	if !w.log.Options().level.Enabled(InfoLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(InfoLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(InfoLevel, args...)
}

func (w *Wrapper) Warn(args ...interface{}) {
	if !w.log.Options().level.Enabled(WarnLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(WarnLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(WarnLevel, args...)
}

func (w *Wrapper) Error(args ...interface{}) {
	if !w.log.Options().level.Enabled(ErrorLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(ErrorLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(ErrorLevel, args...)
}

func (w *Wrapper) Fatal(args ...interface{}) {
	if !w.log.Options().level.Enabled(FatalLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {
	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).Logf(FatalLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(FatalLevel, args...)
	os.Exit(1)
}

func (w *Wrapper) WithContext(ctx context.Context) *Wrapper {
	ns := copyFields(w.fields)
	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func (w *Wrapper) WithError(err error) *Wrapper {
	ns := copyFields(w.fields)
	ns[constant.LogErrorKey] = err
	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func (w *Wrapper) WithFields(fields map[string]interface{}) *Wrapper {
	ns := copyFields(fields)
	for k, v := range w.fields {
		ns[k] = v
	}
	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

package log

import "golang.org/x/net/context"

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

func (w *Wrapper) WithContext(ctx context.Context) *Wrapper {
	ns := copyFields(w.fields)
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

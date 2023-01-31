package http

import (
	"context"
	"github.com/ppxb/unicorn/pkg/utils"
	"net/http"
)

type HttpOptions struct {
	ctx       context.Context
	host      string
	port      int
	urlPrefix string
	handler   http.Handler
	exit      func()
}

func WithHttpCtx(ctx context.Context) func(*HttpOptions) {
	return func(options *HttpOptions) {
		if !utils.InterfaceIsNil(ctx) {
			getHttpOptionsOrSetDefault(options).ctx = ctx
		}
	}
}

func WithHttpHost(host string) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptionsOrSetDefault(options).host = host
	}
}

func WithHttpPort(port int) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptionsOrSetDefault(options).port = port
	}
}

func WithHttpUrlPrefix(prefix string) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptionsOrSetDefault(options).urlPrefix = prefix
	}
}

func WithHttpHandler(handler http.Handler) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptionsOrSetDefault(options).handler = handler
	}
}

func WithHttpExit(f func()) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptionsOrSetDefault(options).exit = f
	}
}

func getHttpOptionsOrSetDefault(options *HttpOptions) *HttpOptions {
	if options == nil {
		return &HttpOptions{
			ctx:       context.Background(),
			host:      "0.0.0.0",
			port:      9527,
			urlPrefix: "api",
		}
	}

	return options
}

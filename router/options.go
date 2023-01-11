package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ppxb/unicorn/pkg/middleware"
)

type Options struct {
	redis       redis.UniversalClient
	redisBinlog bool
	group       *gin.RouterGroup
	jwt         bool
	jwtOps      []func(*middleware.JwtOptions)
	casbin      bool
	casbinOps   []func(*middleware.CasbinOptions)
}

func WithGroup(group *gin.RouterGroup) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).group = group
	}
}

func WithJwt(flag bool) func(*Options) {
	return func(options *Options) {
		getOptionsOrSetDefault(options).jwt = flag
	}
}

func getOptionsOrSetDefault(options *Options) *Options {
	if options == nil {
		return &Options{
			redisBinlog: false,
			jwt:         true,
			casbin:      true,
		}
	}
	return options
}

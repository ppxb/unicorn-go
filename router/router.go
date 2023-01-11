package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ppxb/unicorn/api"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/middleware"
)

type Router struct {
	ops Options
}

func newRouter(options ...func(*Options)) *Router {
	ops := getOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	if ops.group == nil {
		panic("api group为空")
	}

	r := &Router{
		ops: *ops,
	}
	return r
}

func InitRouter(ctx context.Context) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Cors())

	apiGroup := r.Group(global.Config.Server.ApiPrefix)
	apiGroup.GET("/ping", api.Ping)

	v1Group := apiGroup.Group(global.Config.Server.ApiVersion)

	nr := newRouter(
		WithGroup(v1Group),
		WithJwt(true),
	)

	// register routes
	InitUserRouter(nr)

	fmt.Println("初始化router成功")

	return r
}
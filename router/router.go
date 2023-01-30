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

	InitUserRouter(nr)

	fmt.Println("[初始化] Router成功")
	return r
}

func (r Router) Group(path string) gin.IRoutes {
	router := r.ops.group.Group(path)
	return router
}

func (r Router) Casbin(path string) gin.IRoutes {
	router := r.Group(path)
	if r.ops.jwt {
		router.Use()
	}
	if r.ops.casbin {
		router.Use(middleware.Casbin(r.ops.casbinOps...))
	}
	return router
}

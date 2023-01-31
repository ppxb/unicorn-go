package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ppxb/unicorn/api"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/middleware"
)

func Register(s *server.Hertz) {
	apiGroup := s.Group(global.Config.Server.ApiPrefix)
	//apiGroup.GET("/ping", api.Ping)

	middleware.InitJwt()

	testGroup := apiGroup.Group("test")
	testGroup.Use(middleware.JwtMiddleware.MiddlewareFunc())
	testGroup.GET("/ping", api.Ping)
	v1Group := apiGroup.Group(global.Config.Server.ApiVersion)

	InitBaseRouter(v1Group)
}

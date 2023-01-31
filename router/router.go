package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ppxb/unicorn/api"
	"github.com/ppxb/unicorn/pkg/global"
)

func Register(s *server.Hertz) {
	apiGroup := s.Group(global.Config.Server.ApiPrefix)
	apiGroup.GET("/ping", api.Ping)

	v1Group := apiGroup.Group(global.Config.Server.ApiVersion)

	InitBaseRouter(v1Group)
}

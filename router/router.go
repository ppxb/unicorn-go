package router

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	"github.com/ppxb/unicorn/api"
	_ "github.com/ppxb/unicorn/docs"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
)

func Register(s *server.Hertz) {
	apiGroup := s.Group(global.Config.Server.ApiPrefix)
	apiGroup.GET("/ping", api.Ping)

	middleware.InitJwt()

	s.Use(recovery.Recovery())
	s.GET("/swagger/*any",
		swagger.WrapHandler(swaggerFiles.Handler,
			swagger.URL("http://localhost:8848/swagger/doc.json"),
			swagger.DocExpansion("none"),
		),
	)

	v1Group := apiGroup.Group(global.Config.Server.ApiVersion)
	v1Group.
		Use(middleware.JwtMiddleware.MiddlewareFunc()).
		Use(middleware.CasbinHandler())

	InitBaseRouter(apiGroup)
	InitUserRouter(v1Group)
}

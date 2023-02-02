package router

import (
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
	//apiGroup.GET("/ping", api.Ping)

	middleware.InitJwt()

	s.GET("/swagger/*any",
		swagger.WrapHandler(swaggerFiles.Handler,
			swagger.URL("http://localhost:8848/swagger/doc.json"),
			swagger.DocExpansion("none"),
		),
	)

	testGroup := apiGroup.Group("test")
	testGroup.Use(middleware.JwtMiddleware.MiddlewareFunc()).Use(middleware.CasbinHandler())
	testGroup.GET("/ping", api.Ping)

	v1Group := apiGroup.Group(global.Config.Server.ApiVersion)

	InitBaseRouter(v1Group)
}

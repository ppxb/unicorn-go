package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/ppxb/unicorn/pkg/middleware"
)

func InitBaseRouter(r *route.RouterGroup) {
	router := r.Group("base")
	router.POST("/login", middleware.JwtMiddleware.LoginHandler)
}

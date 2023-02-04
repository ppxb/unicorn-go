package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/ppxb/unicorn/pkg/middleware"
)

func InitBaseRouter(r *route.RouterGroup) {
	router := r.Group("base")
	router.POST("/login", middleware.JwtMiddleware.LoginHandler)
	// refresh token should with signed token,the max refresh time is expire time add refresh time
	// e.g. the expire-time is 1 hour and max refresh-time is 1 hour so the final refresh-time is 2 hour
	router.GET("/refresh_token", middleware.JwtMiddleware.RefreshHandler)
}

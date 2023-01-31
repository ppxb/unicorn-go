package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/ppxb/unicorn/api/v1"
)

func InitBaseRouter(r *route.RouterGroup) {
	router := r.Group("")
	router.POST("login", v1.Login)
}

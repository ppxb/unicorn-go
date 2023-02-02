package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/ppxb/unicorn/api/v1"
)

func InitUserRouter(r *route.RouterGroup) {
	router := r.Group("user")
	router.POST("/create", v1.CreateUser)
}

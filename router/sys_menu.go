package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	v1 "github.com/ppxb/unicorn/api/v1"
)

func InitMenuRouter(r *route.RouterGroup) {
	router := r.Group("menu")
	router.GET("/tree", v1.GetMenuTree)
}

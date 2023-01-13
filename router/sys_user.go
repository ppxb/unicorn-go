package router

import v1 "github.com/ppxb/unicorn/api/v1"

func InitUserRouter(r *Router) {
	router := r.ops.group.Group("/user")
	router.POST("/create", v1.CreateUser)
}

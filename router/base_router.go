package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ppxb/unicorn/api/v1"
)

func InitBaseRouter(r *gin.RouterGroup) {
	baseRouter := r.Group("base")
	baseRouter.POST("/login", v1.Login)
}

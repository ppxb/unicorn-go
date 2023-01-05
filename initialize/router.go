package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/unicorn/api"
	"github.com/ppxb/unicorn/router"
)

func Router() *gin.Engine {
	r := gin.Default()

	publicGroup := r.Group("")
	publicGroup.GET("/ping", api.Ping)

	router.InitBaseRouter(publicGroup)

	//privateGroup := router.Group("")

	return r
}

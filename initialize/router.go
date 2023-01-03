package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.Default()

	publicGroup := router.Group("")
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]string{
				"status": "ok",
			})
		})
	}

	return router
}

package dto

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShouldBind(c *gin.Context, r interface{}) {
	err := c.ShouldBind(r)
	if err != nil {
		fmt.Println(err.Error())
	}
}

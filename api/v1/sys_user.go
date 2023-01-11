package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/unicorn/pkg/dto"
)

func CreateUser(c *gin.Context) {
	var r dto.CreateUser
	dto.ShouldBind(c, &r)
}

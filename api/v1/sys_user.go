package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ppxb/unicorn/pkg/dto"
	"github.com/ppxb/unicorn/pkg/models"
	"github.com/ppxb/unicorn/pkg/service"
	"github.com/ppxb/unicorn/pkg/utils"
)

func CreateUser(c *gin.Context) {
	var r dto.CreateUser
	dto.ShouldBind(c, &r)
	my := service.New(c)
	r.Password = utils.GenPwd(r.Password)
	err := my.Q.Create(r, new(models.SysUser))
	if err != nil {
		fmt.Println(err)
	}
}

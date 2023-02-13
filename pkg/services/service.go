package services

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/repository"
)

func GetCurrentUser(c *app.RequestContext) (nu models.SysUser) {
	tu, exists := c.Get(global.Config.Jwt.IdentityKey)
	if !exists {
		return
	}
	nu = repository.GetUserByUUID(tu.(*models.SysUser).UUID)
	return
}

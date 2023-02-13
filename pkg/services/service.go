package services

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
)

func GetCurrentUser(c *app.RequestContext) string {
	user, _ := c.Get(global.Config.Jwt.IdentityKey)
	return user.(*models.SysUser).UUID
}

package services

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/repository"
)

type IUserService interface {
	GetUserInfo(c *app.RequestContext)
}

type UserServiceImpl struct {
}

func (service UserServiceImpl) GetUserInfo(c *app.RequestContext) models.SysUser {
	user, _ := c.Get(global.Config.Jwt.IdentityKey)
	uuid := user.(*models.SysUser).UUID
	return repository.GetUserByUUID(uuid)
}

package services

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/repository"
)

type IUserService interface {
	GetUserInfo(c *app.RequestContext)
}

type UserServiceImpl struct {
}

func (service *UserServiceImpl) GetUserInfo(c *app.RequestContext) models.SysUser {
	uuid := GetCurrentUser(c)
	return repository.GetUserByUUID(uuid)
}

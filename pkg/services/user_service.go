package services

import (
	"github.com/cloudwego/hertz/pkg/app"
)

type IUserService interface {
	GetUserInfo(c *app.RequestContext)
}

type UserServiceImpl struct {
}

package services

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/repository"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

type IBaseService interface {
	Login(req request.Login) (models.SysUser, error)
	GetCurrentUser(c *app.RequestContext)
}

type BaseServiceImpl struct {
	Ctx context.Context
}

func (service *BaseServiceImpl) Login(req request.Login) (u models.SysUser, err error) {
	u, err = repository.GetUserByMobile(req.Mobile)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, errors.New("用户不存在")
	}
	if err = utils.ComparePwd(u.Password, req.Password); err != nil {
		return u, errors.New("密码错误")
	}
	return u, err
}

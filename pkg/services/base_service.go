package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/repository"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

type IBaseService interface {
	Login(req request.Login) (*models.SysUser, error)
}

type BaseServiceImpl struct {
	Ctx context.Context
}

func (service *BaseServiceImpl) Login(req request.Login) (*models.SysUser, error) {
	user, err := repository.GetUserByMobile(req.Mobile)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	if err = utils.ComparePwd(user.Password, req.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

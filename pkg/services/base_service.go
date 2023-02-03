package services

import (
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

type BaseService interface {
	Login(req request.Login) (*models.SysUser, error)
}

func Login(req request.Login) (*models.SysUser, error) {
	var u *models.SysUser
	if errors.Is(global.Mysql.Where("mobile = ?", req.Mobile).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	if err := utils.ComparePwd(u.Password, req.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	return u, nil
}

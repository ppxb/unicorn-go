package service

import (
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

func Login(req request.Login) (u *models.SysUser, err error) {
	if errors.Is(global.Mysql.Where("mobile = ?", req.Mobile).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	if ok := utils.ComparePwd(req.Password, u.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return u, nil
}

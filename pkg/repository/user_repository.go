package repository

import (
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
)

func GetUserByMobile(mobile string) (user *models.SysUser, err error) {
	if err = global.Mysql.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

package repository

import (
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
)

func GetUserByMobile(mobile string) (u models.SysUser, err error) {
	err = global.Mysql.
		Where("mobile = ?", mobile).
		Where("status = ?", models.SysUserStatusEnable).
		First(&u).Error
	return u, err
}

func GetUserByUUID(uuid string) (u models.SysUser) {
	var nu models.SysUser
	err := global.Mysql.
		Preload("Role").
		Where("uuid = ?", uuid).
		Where("status = ?", models.SysUserStatusEnable).
		First(&u).Error
	if err != nil {
		return nu
	}
	return u
}

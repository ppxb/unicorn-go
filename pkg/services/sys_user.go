package services

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

func CreateUser(r request.CreateUser) (err error) {
	var u models.SysUser
	if errors.Is(global.Mysql.Where("mobile = ?", r.Mobile).First(&u).Error, gorm.ErrRecordNotFound) {
		utils.Struct2StructByJson(r, &u)
		u.Password = utils.GenPwd(r.Password)
		u.UUID = uuid.NewString()
		return global.Mysql.Create(&u).Error
	}
	return errors.New("用户已存在")
}

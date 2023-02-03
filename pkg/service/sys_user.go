package service

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
	err = global.Mysql.Where("mobile = ?", r.Mobile).First(&u).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Struct2StructByJson(r, &u)
		u.Password = utils.GenPwd(r.Password)
		u.UUID = uuid.NewString()
		err = global.Mysql.Create(&u).Error
		return
	}
	return errors.New("用户已存在")
}

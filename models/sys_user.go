package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

type SysUser struct {
	M
	UUID     string    `gorm:"index;comment:用户UUID" json:"uuid"`
	Mobile   string    `gorm:"index:idx_mobile,unique;comment:用户手机号" json:"mobile"`
	Password string    `gorm:"comment:用户密码" json:"password"`
	RoleId   uint      `gorm:"comment:角色ID" json:"roleId"`
	Role     SysRole   `gorm:"foreignKey:RoleId;references:RoleId;comment:角色" json:"role"`
	Roles    []SysRole `gorm:"many2many:sys_user_role" json:"roles"`
}

func InitUsers(db *gorm.DB) {
	users := []SysUser{
		{
			UUID:     uuid.NewString(),
			Mobile:   "110",
			Password: utils.GenPwd("123"),
			RoleId:   1001,
		},
	}
	if err := db.Create(&users); err != nil {
		fmt.Println(err)
	}

	if err := db.Model(&users[0]).Association("Roles").Replace(
		[]*SysRole{
			{RoleId: 1001},
		}); err != nil {
		fmt.Println("初始化关联失败")
	}
}

package models

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type SysRole struct {
	CreatedAt  carbon.DateTime `gorm:"comment:创建时间" json:"createdAt"`
	UpdatedAt  carbon.DateTime `gorm:"comment:更新时间" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt  `gorm:"index" json:"-"`
	RoleId     uint            `gorm:"not null;unique;primaryKey;comment:角色ID" json:"roleId"`
	RoleName   string          `gorm:"comment:角色名" json:"roleName"`
	DataRoleId []*SysRole      `gorm:"many2many:sys_data_role_id" json:"dataRoleId"`
	ParentId   int             `gorm:"comment:父角色ID" json:"parentId"`
	Children   []SysRole       `gorm:"-" json:"children"`
	Users      []SysUser       `gorm:"many2many:sys_user_role" json:"-"`
}

func InitRoles(db *gorm.DB) {
	roles := []SysRole{
		{
			RoleId:   1001,
			RoleName: "测试用户",
			ParentId: 0,
		},
		{
			RoleId:   10011,
			RoleName: "测试用户子用户",
			ParentId: 1001,
		},
	}
	if err := db.Create(&roles); err != nil {
		fmt.Println(err.Error)
	}
}
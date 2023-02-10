package models

import "github.com/golang-module/carbon/v2"

type SysUser struct {
	M
	UUID       string          `gorm:"index:idx_uuid;unique;comment:用户UUID" json:"uuid"`
	Mobile     string          `gorm:"index:idx_mobile,unique;comment:用户手机号" json:"mobile"`
	Password   string          `gorm:"comment:用户密码" json:"password"`
	Name       string          `gorm:"comment:用户姓名" json:"name"`
	Avatar     string          `gorm:"comment:用户头像" json:"avatar"`
	Status     uint            `gorm:"type:tinyint(1);default:0;comment:用户状态(0：未禁用；1：已禁用)" json:"status"`
	RoleId     uint            `gorm:"comment:角色ID" json:"roleId"`
	Role       SysRole         `gorm:"foreignKey:RoleId" json:"role"`
	LastLogin  carbon.DateTime `gorm:"comment:最后登录时间" json:"lastLogin"`
	Locked     uint            `gorm:"type:tinyint(1);default:0;comment:锁定状态(0：未锁定；1：已锁定)" json:"locked"`
	LockExpire int64           `gorm:"comment:锁定时间" json:"lockExpire"`
	PassWrong  int             `gorm:"comment:密码错误次数" json:"passWrong"`
}

type UserInfoResp struct {
	UUID      string          `json:"uuid"`
	Mobile    string          `json:"mobile"`
	Name      string          `json:"name"`
	Avatar    string          `json:"avatar"`
	LastLogin carbon.DateTime `json:"lastLogin"`
}

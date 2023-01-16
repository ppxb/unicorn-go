package models

type SysUser struct {
	Mobile   string `gorm:"index:idx_mobile,unique;comment:用户手机号" json:"mobile"`
	Password string `gorm:"comment:用户密码" json:"password"`
}

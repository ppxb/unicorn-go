package models

type SysRole struct {
	M
	Name    string    `gorm:"comment:角色名称" json:"name"`
	Keyword string    `gorm:"index:idx_keyword;unique;comment:关键字(唯一)" json:"keyword"`
	Desc    string    `gorm:"comment:角色描述" json:"desc"`
	Status  *uint     `gorm:"type:tinyint(1);default:1;comment:角色状态(0：禁用；1：启用)" json:"status"`
	Sort    *uint     `gorm:"default:1;comment:排序>=0，值越小权限越高。0为超级管理员" json:"sort"`
	Menus   []uint    `gorm:"-" json:"menus"`
	Users   []SysUser `gorm:"foreignKey:RoleId" json:"users"`
}

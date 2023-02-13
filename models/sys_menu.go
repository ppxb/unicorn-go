package models

type SysMenu struct {
	M
	Name     string    `gorm:"comment:名称" json:"name"`
	Title    string    `gorm:"标题" json:"title"`
	Icon     string    `gorm:"图标" json:"icon"`
	Path     string    `gorm:"url路径" json:"path"`
	Sort     *uint     `gorm:"type:int unsigned;comment:排序应>=0" json:"sort"`
	Status   *uint     `gorm:"type:tinyint(1);default:0;comment:状态(0：未禁用；1：已禁用)" json:"status"`
	Visible  *uint     `gorm:"type:tinyint(1);default:0;comment:显示(0：正常；1：隐藏)" json:"visible"`
	ParentId uint      `gorm:"default:0;comment:父级菜单ID" json:"parentId"`
	Children []SysMenu `gorm:"-" json:"children"`
	RoleIds  []uint    `gorm:"-" json:"roleIds"`
}

type SysMenuRoleRelation struct {
	MenuId uint `json:"menuId"`
	RoleId uint `json:"roleId"`
}

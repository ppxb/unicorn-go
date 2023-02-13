package repository

import (
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/constant"
	"github.com/ppxb/unicorn/pkg/global"
)

func GetAllMenu() (menus []models.SysMenu) {
	global.Mysql.Model(&models.SysMenu{}).Find(&menus)
	return menus
}

func GetMenuByRoleId(roleId, roleSort uint) (rp []models.SysMenu) {
	menuIds := make([]uint, 0)
	global.Mysql.
		Model(&models.SysMenuRoleRelation{}).
		Where("role_id = ?", roleId).
		Pluck("menu_id", &menuIds)
	if len(menuIds) > 0 {
		q := global.Mysql.
			Model(&models.SysMenu{}).
			Where("id IN (?)", menuIds)
		if roleSort != 0 {
			q.Where("status = ?", 0)
		}
		q.Order("sort").Find(&rp)
	}
	return
}

func FindMenuByRoleId(roleId, roleSort uint) []models.SysMenu {
	menuIds := make([]uint, 0)
	global.Mysql.
		Model(&models.SysMenuRoleRelation{}).
		Where("role_id = ?", roleId).
		Pluck("menu_id", &menuIds)
	rp := make([]models.SysMenu, 0)
	if len(menuIds) > 0 {
		q := global.Mysql.
			Model(models.SysMenu{}).
			Where("id IN (?)", menuIds)
		if roleSort != constant.Zero {
			q.Where("status = ?", constant.Zero)
		}
		q.Order("sort").Find(&rp)
	}
	return rp
}

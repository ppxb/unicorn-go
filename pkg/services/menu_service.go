package services

import (
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/constant"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/repository"
	"github.com/ppxb/unicorn/pkg/utils"
)

type IMenuService interface {
	GetMenuTree(roleId, roleSort uint) []models.SysMenu
	FindMenu(currentRoleId, currentRoleSort uint) []models.SysMenu
}

type MenuServiceImpl struct {
}

func (service *MenuServiceImpl) GetMenuTree(roleId, roleSort uint) []models.SysMenu {
	allMenu := repository.GetAllMenu()
	roleMenu := repository.GetMenuByRoleId(roleId, roleSort)
	_, newMenus := addParentMenu(roleMenu, allMenu)

	return GenMenuTree(0, newMenus)
}

func GenMenuTree(parentId uint, roleMenus []models.SysMenu) (tree []models.SysMenu) {
	roleMenuIds := make([]uint, 0)
	allMenu := make([]models.SysMenu, 0)
	global.Mysql.Model(&models.SysMenu{}).Find(&allMenu)
	_, newRoleMenus := addParentMenu(roleMenus, allMenu)
	for _, menu := range newRoleMenus {
		if !utils.ContainsUint(roleMenuIds, menu.Id) {
			roleMenuIds = append(roleMenuIds, menu.Id)
		}
	}
	tree = genMenuTree(parentId, roleMenuIds, allMenu)
	return
}

func (service *MenuServiceImpl) FindMenu(currenRoleId, currentRoleSort uint) []models.SysMenu {
	menus := findMenuByCurrentRole(currenRoleId, currentRoleSort)
	return GenMenuTree(0, menus)
}

func findMenuByCurrentRole(currentRoleId, currentRoleSort uint) []models.SysMenu {
	menus := make([]models.SysMenu, 0)
	if currentRoleSort != constant.Zero {
		menus = repository.FindMenuByRoleId(currentRoleId, currentRoleSort)
	} else {
		global.Mysql.Order("sort").Find(&menus)
	}
	return menus
}

func genMenuTree(parentId uint, roleMenuIds []uint, allMenu []models.SysMenu) (tree []models.SysMenu) {
	tree = make([]models.SysMenu, 0)
	for _, menu := range allMenu {
		if !utils.ContainsUint(roleMenuIds, menu.Id) {
			continue
		}
		if menu.ParentId == parentId {
			menu.Children = genMenuTree(menu.Id, roleMenuIds, allMenu)
			tree = append(tree, menu)
		}
	}
	return
}

func addParentMenu(menus, all []models.SysMenu) (newMenuIds []uint, newMenus []models.SysMenu) {
	parentIds := make([]uint, 0)
	menuIds := make([]uint, 0)
	for _, menu := range menus {
		if menu.ParentId > 0 {
			parentIds = append(parentIds, menu.ParentId)
			parentMenuIds := findParentMenuId(menu.ParentId, all)
			if len(parentMenuIds) > 0 {
				parentIds = append(parentIds, parentMenuIds...)
			}
		}
		menuIds = append(menuIds, menu.Id)
	}

	if len(parentIds) > 0 {
		menuIds = append(menuIds, parentIds...)
	}
	newMenuIds = make([]uint, 0)
	newMenus = make([]models.SysMenu, 0)
	for _, menu := range all {
		for _, id := range menuIds {
			if id == menu.Id && !utils.ContainsUint(newMenuIds, id) {
				newMenus = append(newMenus, menu)
				newMenuIds = append(newMenuIds, id)
			}
		}
	}
	return
}

func findParentMenuId(menuId uint, all []models.SysMenu) (parentIds []uint) {
	var currentMenu models.SysMenu
	parentIds = make([]uint, 0)
	for _, menu := range all {
		if menuId == menu.Id {
			currentMenu = menu
			break
		}
	}
	if currentMenu.ParentId == 0 {
		return parentIds
	}
	parentIds = append(parentIds, currentMenu.ParentId)
	newParentIds := findParentMenuId(currentMenu.ParentId, all)
	if len(newParentIds) > 0 {
		parentIds = append(parentIds, newParentIds...)
	}
	return
}

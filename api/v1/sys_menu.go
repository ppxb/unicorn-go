package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/resp"
	"github.com/ppxb/unicorn/pkg/services"
)

var menuService = &services.MenuServiceImpl{}

func GetMenuTree(ctx context.Context, c *app.RequestContext) {
	u := services.GetCurrentUser(c)
	menus := menuService.GetMenuTree(u.RoleId, *u.Role.Sort)
	resp.SuccessWithData(menus, c)
}

func FindMenu(ctx context.Context, c *app.RequestContext) {
	u := services.GetCurrentUser(c)
	menus := menuService.FindMenu(u.RoleId, *u.Role.Sort)
	resp.SuccessWithData(menus, c)
}

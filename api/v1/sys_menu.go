package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/services"
)

var menuService = &services.MenuServiceImpl{}

func GetMenuTree(ctx context.Context, c *app.RequestContext) {
	menuService.GetMenuTree(c)
}

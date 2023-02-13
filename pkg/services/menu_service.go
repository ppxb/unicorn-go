package services

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

type IMenuService interface {
	GetMenuTree(c *app.RequestContext)
}

type MenuServiceImpl struct {
}

func (service *MenuServiceImpl) GetMenuTree(c *app.RequestContext) {
	fmt.Println(GetCurrentUser(c))
}

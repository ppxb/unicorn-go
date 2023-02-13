package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/log"
	"github.com/ppxb/unicorn/pkg/resp"
	"github.com/ppxb/unicorn/pkg/services"
	"sync"
)

var userService = &services.UserServiceImpl{}

func CasbinHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u := userService.GetUserInfo(c)
		obj := string(c.Request.URI().Path())
		act := string(c.Request.Method())
		sub := u.Role.Keyword
		if !check(sub, obj, act) {
			resp.SuccessWithMsg("没有该权限", c)
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

var checkLock sync.Mutex

func check(sub, obj, act string) bool {
	checkLock.Lock()
	defer checkLock.Unlock()

	err := global.CasbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error(fmt.Sprintf("Casbin 读取策略失败：%s", err.Error()))
	}
	pass, _ := global.CasbinEnforcer.Enforce(sub, obj, act)
	return pass
}

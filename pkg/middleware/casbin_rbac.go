package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/response"
)

func CasbinHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		user, _ := c.Get(IdentityKey)
		obj := string(c.Request.URI().Path())
		act := string(c.Request.Method())
		sub := user.(*User).Username
		if ok, _ := global.CasbinEnforcer.Enforce(sub, obj, act); !ok {
			response.SuccessWithMsg("权限不足", c)
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

func Casbin(options ...func(*CasbinOptions)) gin.HandlerFunc {
	ops := getCasbinOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	if ops.Enforcer == nil {
		panic("casbin enforcer is empty")
	}
	return func(c *gin.Context) {
		sub := ops.getCurrentUser(c)
		obj := strings.Replace(c.Request.URL.Path, "/"+ops.urlPrefix, "", 1)
		act := c.Request.Method
		if !check(sub.RoleKeyword, obj, act, *ops) {
			ops.failWithCode(400)
			return
		}
		c.Next()
	}
}

var checkLock sync.Mutex

func check(sub, obj, act string, ops CasbinOptions) bool {
	checkLock.Lock()
	defer checkLock.Unlock()
	pass, _ := ops.Enforcer.Enforce(sub, obj, act)
	return pass
}

package main

import (
	"context"
	"embed"
	"github.com/ppxb/unicorn/initialize"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/server"
	"github.com/ppxb/unicorn/router"
	"github.com/zeromicro/go-zero/core/logx"
	"runtime"
	"runtime/debug"
	"strings"
)

//go:embed conf
var conf embed.FS

var ctx = context.Background()

func main() {
	defer func() {
		if err := recover(); err != nil {
			logx.WithContext(ctx).Errorf("[服务器] 启动失败，堆栈信息：%s", string(debug.Stack()))
		}
	}()

	_, file, _, _ := runtime.Caller(0)
	global.RuntimeRoot = strings.TrimSuffix(file, "main.go")

	initialize.Config(ctx, conf)
	initialize.Redis()
	initialize.Mysql()
	initialize.CasbinEnforcer()

	server.Listen(
		server.WithHttpCtx(ctx),
		server.WithHttpHost(global.Config.Server.Host),
		server.WithHttpPort(global.Config.Server.Port),
		server.WithHttpHandler(router.InitRouter(ctx)),
	)
}

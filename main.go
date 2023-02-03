package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ppxb/unicorn/initialize"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/router"
	"time"
)

//go:embed conf
var conf embed.FS

var ctx = context.Background()

// @title unicorn Framework
// @version 1.0.0
// @description 一个使用Go开发的RBAC基础框架
// @termsOfService  http://swagger.io/terms/

// @contact.name ppxb
// @contact.url http://github.com/ppxb

// @license.name MIT
// @license.url https://github.com/ppxb/unicorn-go/blob/master/LICENCE

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	initialize.Config(ctx, conf)
	initialize.Mysql()
	initialize.Casbin()

	s := server.New(
		server.WithHostPorts(fmt.Sprintf(":%d", global.Config.Server.Port)),
		server.WithExitWaitTime(5*time.Second),
	)

	router.Register(s)
	s.Spin()
}

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

// @title unicorn framework
// @version 1.0.0
// @description A simple RBAC base framework built by go.
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

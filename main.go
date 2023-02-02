package main

import (
	"context"
	"embed"
	"github.com/ppxb/unicorn/initialize"
	"github.com/ppxb/unicorn/pkg/http"
)

//go:embed conf
var conf embed.FS

var ctx = context.Background()

// @title unicorn framework
// @version 1.0.0
// @description A simple RBAC base framework built by go.
// @licence.name MIT
// @licence.url https://github.com/ppxb/unicorn-go/blob/master/LICENCE
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	initialize.Config(ctx, conf)
	initialize.Mysql()
	initialize.Casbin()

	http.Listen()
}

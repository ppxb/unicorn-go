package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/ppxb/unicorn/initialize"
	"github.com/ppxb/unicorn/pkg/global"
	"runtime"
	"strings"
)

//go:embed conf
var conf embed.FS

var ctx = context.Background()

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	_, file, _, _ := runtime.Caller(0)
	global.RuntimeRoot = strings.TrimSuffix(file, "main.go")

	initialize.Config(ctx, conf)
	initialize.Redis()
	initialize.Mysql()
	initialize.CasbinEnforcer()
}

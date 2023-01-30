package main

import (
	"context"
	"embed"
	"github.com/ppxb/unicorn/initialize"
	"github.com/ppxb/unicorn/pkg/log"
)

//go:embed conf
var conf embed.FS

var ctx = context.Background()

func main() {
	initialize.Config(ctx, conf)
	log.Info("初始化成功")
}

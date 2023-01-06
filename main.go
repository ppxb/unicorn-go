package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/ppxb/unicorn/initialize"
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

	initialize.Config(ctx, conf)
	initialize.Mysql()
}

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

func main() {
	initialize.Config(ctx, conf)
	initialize.Mysql()
	initialize.Casbin()

	http.Listen()
}

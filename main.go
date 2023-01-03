package main

import (
	"fmt"
	"github.com/ppxb/unicorn/core"
	"github.com/ppxb/unicorn/initialize"
	"github.com/ppxb/unicorn/pkg/server"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	core.InitViper()
	initialize.Gorm()
	router := initialize.Router()

	server.Http(router)
}

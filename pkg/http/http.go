package http

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/router"
	"time"
)

func Listen() {
	s := server.New(
		server.WithHostPorts(fmt.Sprintf(":%d", global.Config.Server.Port)),
		server.WithExitWaitTime(5*time.Second),
	)
	s.Use(recovery.Recovery())

	router.Register(s)
	s.Spin()
}

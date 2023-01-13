package server

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Listen(options ...func(*HttpOptions)) {
	ops := getHttpOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}

	host := ops.host
	port := ops.port
	ctx := ops.ctx
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: ops.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("[服务器] 启动失败")
		}
	}()

	fmt.Printf("[服务器] 启动成功，监听在：http://%s:%d/%s \n", host, port, ops.urlPrefix)

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if ops.exit != nil {
		ops.exit()
	}
	fmt.Println("[服务器] 关闭中")

	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ops.ctx); err != nil {
		fmt.Println(errors.Wrap(err, "[服务器] 关闭失败"))
	}

	fmt.Println("[服务器] 已关闭")
}

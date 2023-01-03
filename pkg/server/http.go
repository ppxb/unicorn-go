package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ppxb/unicorn/core"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Http(r *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", core.AppConfig.Server.Host, core.AppConfig.Server.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
	}()

	fmt.Printf("server is running at %s:%d \n", core.AppConfig.Server.Host, core.AppConfig.Server.Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("server is shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("server shutdown failed...")
	}

	fmt.Println("server exited...")
}

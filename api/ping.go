package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, map[string]string{
		"ping": "pong",
	})
}

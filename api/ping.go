package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/response"
)

// Ping check server is running
// @Security Bearer
// @Accept json
// @Produce json
// @Success 20001 {object} response.Response "ok"
// @Tags Common
// @Description Ping
// @Router /api/test/ping [GET]
func Ping(ctx context.Context, c *app.RequestContext) {
	response.SuccessWithData(map[string]interface{}{
		"ping": "pong",
	}, c)
}

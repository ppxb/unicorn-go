package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/resp"
)

// Ping check server is running
// @Security Bearer
// @Accept json
// @Produce json
// @Success 20001 {object} resp.Response "ok"
// @Tags 其他接口
// @Description Ping
// @Router /api/ping [GET]
func Ping(ctx context.Context, c *app.RequestContext) {
	resp.SuccessWithData(map[string]interface{}{
		"ping": "pong",
	}, c)
}

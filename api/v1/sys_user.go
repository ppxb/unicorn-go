package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateUser 创建用户
// @Security Bearer
// @Accept json
// @Produce json
// @Success 20001 {object} resp.Response "ok"
// @Tags 用户
// @Description 创建用户
// @Param params body request.CreateUser true "params"
// @Router /api/v1/user/create [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {

}

package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/resp"
	"github.com/ppxb/unicorn/pkg/service"
)

// CreateUser 创建用户
// @Security Bearer
// @Accept json
// @Produce json
// @Success 20001 {object} resp.Response "ok"
// @Tags 用户接口
// @Description 创建用户
// @Param params body request.CreateUser true "params"
// @Router /api/v1/user/create [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var r request.CreateUser
	err := c.BindAndValidate(&r)
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	err = service.CreateUser(r)
	if err != nil {
		resp.CheckError(err, c)
		return
	}
	resp.Success(c)
}

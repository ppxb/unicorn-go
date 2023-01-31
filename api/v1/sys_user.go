package v1

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ppxb/unicorn/pkg/response"
)

type LoginReq struct {
	Username string `json:"username" vd:"len($)>1"`
	Password string `json:"password" vd:"len($)>1"`
}

func Login(ctx context.Context, c *app.RequestContext) {
	var req LoginReq
	err := c.BindAndValidate(&req)
	if err != nil {
		fmt.Println(err.Error())
	}
	response.Success(c)
}

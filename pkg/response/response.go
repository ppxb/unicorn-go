package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
)

const (
	OkCode    = 200001
	ErrorCode = 400001
)

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func Result(code int, msg string, data interface{}, c *app.RequestContext) {
	c.JSON(http.StatusOK, Response{
		Code:      code,
		Msg:       msg,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

func Success(c *app.RequestContext) {
	Result(OkCode, "success", nil, c)
}

func SuccessWithMsg(msg string, c *app.RequestContext) {
	Result(OkCode, msg, nil, c)
}

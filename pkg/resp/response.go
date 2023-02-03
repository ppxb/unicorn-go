package resp

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
)

const (
	OkCode   = 200001
	FailCode = 400001
)

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type Page struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
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

func SuccessWithData(data interface{}, c *app.RequestContext) {
	Result(OkCode, "success", data, c)
}

func Fail(c *app.RequestContext) {
	Result(FailCode, "failed", nil, c)
}

func FailWithMsg(msg string, c *app.RequestContext) {
	Result(FailCode, msg, nil, c)
}

func FailWithData(data interface{}, c *app.RequestContext) {
	Result(FailCode, "failed", data, c)
}

func CheckError(format interface{}, c *app.RequestContext) {
	var f string
	switch format.(type) {
	case string:
		f = format.(string)
	case error:
		f = fmt.Sprintf("%v", format.(error))
	}
	if f != "" {
		FailWithMsg(f, c)
	}
}

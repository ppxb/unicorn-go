package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/log"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/resp"
	"github.com/ppxb/unicorn/pkg/services"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       2 * time.Hour,
		MaxRefresh:    2 * time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			resp.SuccessWithData(map[string]interface{}{
				"token":  token,
				"expire": expire,
			}, c)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req request.Login
			if err := c.BindAndValidate(&req); err != nil {
				return nil, err
			}
			return services.Login(req)
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &models.SysUser{
				Mobile: claims[IdentityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.SysUser); ok {
				return jwt.MapClaims{
					IdentityKey: v.Mobile,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			log.Error(fmt.Sprintf("jwt 错误:%+v", e.Error()))
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			resp.FailWithMsg(message, c)
		},
	})
	if err != nil {
		log.Panic(err.Error())
	}
}

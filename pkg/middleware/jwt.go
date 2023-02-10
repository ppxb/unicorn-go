package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/log"
	"github.com/ppxb/unicorn/pkg/request"
	"github.com/ppxb/unicorn/pkg/resp"
	"github.com/ppxb/unicorn/pkg/services"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = global.Config.Jwt.IdentityKey
	BaseService   = &services.BaseServiceImpl{
		Ctx: context.Background(),
	}
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         global.Config.Jwt.Realm,
		Key:           []byte(global.Config.Jwt.SecretKey),
		Timeout:       global.Config.Jwt.Expire * time.Hour,
		MaxRefresh:    global.Config.Jwt.Refresh * time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			resp.SuccessWithData(map[string]interface{}{
				"token":  token,
				"expire": expire,
			}, c)
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			resp.SuccessWithData(map[string]interface{}{
				"token":  token,
				"expire": expire,
			}, c)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req request.Login

			if err = c.BindAndValidate(&req); err != nil {
				return nil, err
			}

			user, _ := BaseService.Login(req)
			fmt.Println(user)
			return BaseService.Login(req)
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
			log.Error(fmt.Sprintf("JWT 解析错误: %+v", e.Error()))
			if errors.Is(e, jwt.ErrEmptyCookieToken) {
				return "请求未携带token"
			}
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			resp.FailWithMsg(message, c)
		},
	})
	if err != nil {
		log.Error(fmt.Sprintf("JWT 解析错误: %+v", err.Error()))
	}
}

package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/constant"
)

type CorsOptions struct {
	origin     string
	header     string
	expose     string
	method     string
	credential string
}

func getCorsOptionsOrSetDefault(options *CorsOptions) *CorsOptions {
	if options == nil {
		return &CorsOptions{
			origin:     constant.MiddlewareCorsOrigin,
			header:     constant.MiddlewareCorsHeader,
			expose:     constant.MiddlewareCorsExpose,
			method:     constant.MiddlewareCorsMethods,
			credential: constant.MiddlewareCorsCredentials,
		}
	}
	return options
}

type JwtOptions struct {
	realm              string
	key                string
	timeout            int
	maxRefresh         int
	tokenLookup        string
	tokenHeaderName    string
	sendCookie         bool
	cookieName         string
	privateBytes       []byte
	success            func()
	successWithData    func(...interface{})
	failWithMsg        func(format interface{}, a ...interface{})
	failWithCodeAndMsg func(code int, format interface{}, a ...interface{})
	loginPwdCheck      func(c *gin.Context, r interface{}) (userId int64, err error)
}

func getJwtOptionsOrSetDefault(options *JwtOptions) *JwtOptions {
	if options == nil {
		return &JwtOptions{
			realm:           "jwt",
			key:             "test secret",
			timeout:         24,
			maxRefresh:      168,
			tokenLookup:     "header: Authorization, query: token, cookie: jwt",
			tokenHeaderName: "Bearer",
			success: func() {

			},
			successWithData: func(i ...interface{}) {

			},
			failWithMsg: func(format interface{}, a ...interface{}) {

			},
			failWithCodeAndMsg: func(code int, format interface{}, a ...interface{}) {

			},
			loginPwdCheck: func(c *gin.Context, r interface{}) (userId int64, err error) {
				return 0, errors.Errorf("")
			},
		}
	}
	return options
}

type CasbinOptions struct {
	urlPrefix      string
	getCurrentUser func(c *gin.Context) interface{}
	Enforcer       *casbin.Enforcer
	failWithCode   func(code int)
}

func ParseCasbinOptions(options ...func(*CasbinOptions)) *CasbinOptions {
	ops := getCasbinOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	return ops
}

func getCasbinOptionsOrSetDefault(options *CasbinOptions) *CasbinOptions {
	if options == nil {
		return &CasbinOptions{
			urlPrefix: constant.MiddlewareUrlPrefix,
			failWithCode: func(code int) {

			},
		}
	}
	return options
}

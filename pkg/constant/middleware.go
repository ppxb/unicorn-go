package constant

const (
	MiddlewareUrlPrefix = "api"
	// cors
	MiddlewareCorsOrigin      = "*"
	MiddlewareCorsHeader      = "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Sign-Token,api-idempotence-token"
	MiddlewareCorsMethods     = "OPTIONS,GET,POST,PUT,PATCH,DELETE"
	MiddlewareCorsExpose      = "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type"
	MiddlewareCorsCredentials = "true"
)

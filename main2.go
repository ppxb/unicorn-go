package main

////go:embed conf
//var conf embed.FS
//
//var ctx = context.Background()
//
//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			logx.WithContext(ctx).Errorf("[服务器] 启动失败，堆栈信息：%s", string(debug.Stack()))
//		}
//	}()
//
//	_, file, _, _ := runtime.Caller(0)
//	global.RuntimeRoot = strings.TrimSuffix(file, "main2.go")
//
//	initialize.Config(ctx, conf)
//	initialize.Redis()
//	initialize.Mysql()
//	initialize.CasbinEnforcer()
//
//	server.Listen(
//		server.WithHttpCtx(ctx),
//		server.WithHttpHost(global.Config.Server.Host),
//		server.WithHttpPort(global.Config.Server.Port),
//		server.WithHttpHandler(router.InitRouter(ctx)),
//	)
//}

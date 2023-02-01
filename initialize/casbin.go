package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/zeromicro/go-zero/core/logx"
)

func Casbin() {
	e := mysqlCasbin()
	global.CasbinEnforcer = e
	logx.WithContext(ctx).Info("[初始化] Casbin初始化成功")
}

func mysqlCasbin() *casbin.CachedEnforcer {
	adapter, err := gormadapter.NewAdapterByDBUseTableName(
		global.Mysql.WithContext(ctx),
		global.Config.Mysql.TablePrefix,
		"sys_casbin",
	)
	if err != nil {
		panic(errors.Wrap(err, "[初始化] Casbin Enforcer失败"))
	}

	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`
	m, _ := model.NewModelFromString(text)
	enforcer, err := casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		panic(errors.Wrap(err, "[初始化] Casbin Enforcer失败"))
	}
	enforcer.SetExpireTime(60 * 60)
	err = enforcer.LoadPolicy()
	if err != nil {
		panic(errors.Wrap(err, "[初始化] Casbin Enforcer失败"))
	}
	return enforcer
}

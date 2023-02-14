package initialize

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/log"
)

func Casbin() {
	e := mysqlCasbin()
	global.CasbinEnforcer = e
	log.Info("Casbin 初始化成功...")
}

func mysqlCasbin() (enforcer *casbin.Enforcer) {
	adapter, err := gormadapter.NewAdapterByDBUseTableName(
		global.Mysql.WithContext(ctx),
		global.Config.Mysql.TablePrefix,
		"sys_casbin",
	)
	if err != nil {
		log.Panic(fmt.Sprintf("Casbin 初始化失败：%s", err.Error()))
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
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Panic(fmt.Sprintf("Casbin 初始化失败：%s", err.Error()))
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Panic(fmt.Sprintf("Casbin 初始化失败：%s", err.Error()))
	}
	return
}

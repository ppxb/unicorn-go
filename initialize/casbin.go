package initialize

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/global"
)

func CasbinEnforcer() {
	e, err := mysqlCasbin()
	if err != nil {
		panic(errors.Wrap(err, "初始化Casbin Enforcer失败"))
	}
	global.CasbinEnforcer = e
	fmt.Println("初始化Casbin成功")
}

func mysqlCasbin() (e *casbin.Enforcer, err error) {
	adapter, err := gormadapter.NewAdapterByDBUseTableName(
		global.Mysql.WithContext(ctx),
		global.Config.Mysql.TablePrefix,
		"sys_casbin",
	)
	if err != nil {
		return
	}

	config := global.ConfBox.Get(global.Config.Server.CasbinModelPath)
	casbinModel := model.NewModel()
	err = casbinModel.LoadModelFromText(string(config))
	if err != nil {
		return
	}

	e, err = casbin.NewEnforcer(casbinModel, adapter)
	if err != nil {
		return
	}

	err = e.LoadPolicy()
	if err != nil {
		return
	}

	return
}

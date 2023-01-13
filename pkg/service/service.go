package service

import (
	"context"
	"fmt"
	"github.com/ppxb/unicorn/pkg/constant"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/query"
)

type MysqlService struct {
	Q query.Mysql
}

func New(ctx context.Context) MysqlService {
	ops := []func(*query.MysqlOptions){
		query.WithMysqlCtx(ctx),
		query.WithMysqlDb(global.Mysql),
		query.WithMysqlCasbinEnforcer(global.CasbinEnforcer),
		query.WithMysqlCachePrefix(fmt.Sprintf("%s_%s", global.Config.Mysql.DSN.DBName, constant.QueryCachePrefix)),
	}

	if global.Config.Redis.Uri != "" {
		ops = append(ops, query.WithMysqlRedis(global.Redis))
	}
	my := MysqlService{
		Q: query.NewMysql(ops...),
	}
	return my
}

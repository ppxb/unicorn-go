package query

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/ppxb/unicorn/pkg/constant"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
)

type MysqlOptions struct {
	ctx         context.Context
	db          *gorm.DB
	redis       redis.UniversalClient
	cachePrefix string
	enforcer    *casbin.Enforcer
}

func WithMysqlCtx(ctx context.Context) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		if !utils.InterfaceIsNil(ctx) {
			getMysqlOptionsOrSetDefault(options).ctx = ctx
		}
	}
}

func WithMysqlDb(db *gorm.DB) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		if db != nil {
			getMysqlOptionsOrSetDefault(options).db = db
		}
	}
}

func WithMysqlCasbinEnforcer(en *casbin.Enforcer) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		if en != nil {
			getMysqlOptionsOrSetDefault(options).enforcer = en
		}
	}
}

func WithMysqlCachePrefix(prefix string) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		getMysqlOptionsOrSetDefault(options).cachePrefix = prefix
	}
}

func WithMysqlRedis(redis redis.UniversalClient) func(*MysqlOptions) {
	return func(options *MysqlOptions) {
		getMysqlOptionsOrSetDefault(options).redis = redis
	}
}

func getMysqlOptionsOrSetDefault(options *MysqlOptions) *MysqlOptions {
	if options == nil {
		return &MysqlOptions{
			ctx:         context.Background(),
			cachePrefix: constant.QueryCachePrefix,
		}
	}
	return options
}

package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/ppxb/unicorn/pkg/models"
	"gorm.io/gorm"
)

var (
	Mode           string
	RuntimeRoot    string
	Config         Configuration
	ConfBox        models.ConfBox
	Mysql          *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	Redis          redis.UniversalClient
)

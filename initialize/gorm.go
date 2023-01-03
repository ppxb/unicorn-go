package initialize

import (
	"github.com/ppxb/unicorn/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch core.AppConfig.Database.Type {
	case "mysql":
		return newGormMysql()
	case "postgres":
		return newGormPostgres()
	default:
		return newGormMysql()
	}
}

func newGormMysql() *gorm.DB {
	m := core.AppConfig.Database
	if m.DbName == "" {
		return nil
	}

	config := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
	}

	if db, err := gorm.Open(mysql.New(config)); err != nil {
		return nil
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(10)
		sqlDb.SetMaxOpenConns(100)
		return db
	}
}

func newGormPostgres() *gorm.DB {
	return nil
}

package initialize

import (
	"context"
	"embed"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/log"
	"github.com/ppxb/unicorn/pkg/migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

//go:embed db/*.sql
var sqlFs embed.FS

func Mysql() {
	cfg, err := m.ParseDSN(global.Config.Mysql.Uri)
	if err != nil {
		log.Panic(fmt.Sprintf("Mysql 初始化失败：%s", err.Error()))
	}
	global.Config.Mysql.DSN = *cfg

	err = migrate.WithHooks(
		migrate.WithContext(ctx),
		migrate.WithUri(global.Config.Mysql.Uri),
		migrate.WithFs(sqlFs),
		migrate.WithFsRoot("db"),
		migrate.WithBefore(beforeMigrate),
	)
	if err != nil {
		log.Panic(fmt.Sprintf("Mysql 初始化失败：%s", err.Error()))
	}

	log.Info("Mysql 初始化成功...")
}

func beforeMigrate(ctx context.Context) {
	var cancel context.CancelFunc

	init := false
	ctx, cancel = context.WithTimeout(ctx, time.Duration(global.Config.Server.ConnectTimeout)*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					log.Panic(fmt.Sprintf("Mysql 初始化失败，连接超时(%ds)", global.Config.Server.ConnectTimeout))
				}
				// avoid goroutine deadlock
				return
			}
		}
	}()

	db, err := gorm.Open(mysql.Open(global.Config.Mysql.DSN.FormatDSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Config.Mysql.TablePrefix + "_",
			SingularTable: true,
		},
		QueryFields: true,
	})
	if err != nil {
		log.Panic(fmt.Sprintf("Mysql 初始化失败：%s", err.Error()))
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)

	init = true
	global.Mysql = db
}

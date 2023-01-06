package initialize

import (
	"context"
	"embed"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var sqlFs embed.FS

func Mysql() {
	cfg, err := m.ParseDSN(global.Config.Mysql.Uri)
	if err != nil {
		panic(errors.Wrap(err, "初始化数据库失败"))
	}
	global.Config.Mysql.DSN = *cfg

	err = migrate.WithHooks(
		migrate.WithCtx(ctx),
		migrate.WithUri(global.Config.Mysql.Uri),
		migrate.WithBefore(beforeMigrate),
	)
	if err != nil {
		panic(errors.Wrap(err, "初始化数据库失败"))
	}

	fmt.Println("初始化数据库成功")
}

func beforeMigrate(ctx context.Context) (err error) {
	var cancel context.CancelFunc

	init := false
	ctx, cancel = context.WithTimeout(ctx, time.Duration(global.Config.Server.ConnectTimeout)*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("initialize mysql failed: connect timeout(%ds", global.Config.Server.ConnectTimeout))
				}
				// avoid goroutine deadlock
				return
			}
		}
	}()

	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(global.Config.Mysql.DSN.FormatDSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Config.Mysql.TablePrefix + "_",
			SingularTable: true,
		},
		QueryFields: true,
	})
	if err != nil {
		return
	}

	init = true
	global.Mysql = db
	return
}

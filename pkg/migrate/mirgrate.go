package migrate

import (
	"database/sql"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func WithHooks(options ...func(*Options)) (err error) {
	ops := getOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}

	err = database(ops)
	if err != nil {
		return
	}

	_, err = sql.Open(ops.driver, ops.uri)
	if err != nil {
		fmt.Println(errors.Wrap(err, "打开数据库连接失败"))
		return
	}

	if ops.before != nil {
		err = ops.before(ops.ctx)
		if err != nil {
			fmt.Println(errors.Wrap(err, "执行before hook失败"))
			return
		}
	}
	return
}

func database(ops *Options) (err error) {
	var cfg *m.Config
	var db *sql.DB

	cfg, err = m.ParseDSN(ops.uri)
	if err != nil {
		fmt.Println(errors.Wrap(err, "无效的数据库uri"))
		return
	}

	dbname := cfg.DBName
	cfg.DBName = ""
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname))
	if err != nil {
		fmt.Println(errors.Wrap(err, "创建数据库失败"))
	}
	return
}

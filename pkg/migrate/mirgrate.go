package migrate

import (
	"database/sql"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
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

	var db *sql.DB
	db, err = sql.Open(ops.driver, ops.uri)
	if err != nil {
		fmt.Println(errors.Wrap(err, "[数据库] 打开数据库连接失败"))
		return
	}

	defer func() {
		releaseErr := releaseLock(ops, db)
		if releaseErr != nil && err != nil {
			err = releaseErr
		}
	}()

	var lockAcquired bool
	for {
		lockAcquired, err = acquireLock(ops, db)
		if err != nil {
			return
		}
		if lockAcquired {
			break
		} else {
			fmt.Println("[数据库] 无法获得锁，正在重试")
		}
	}

	if ops.before != nil {
		err = ops.before(ops.ctx)
		if err != nil {
			fmt.Println(errors.Wrap(err, "[数据库] 执行before hook失败"))
			return
		}
	}

	migrate.SetTable(ops.changeTable)
	source := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: ops.fs,
		Root:       ops.fsRoot,
	}
	err = status(ops, db, source)
	if err != nil {
		fmt.Println("[数据库] 显示迁移状态失败")
		return
	}

	_, err = migrate.Exec(db, ops.driver, source, migrate.Up)
	if err != nil {
		fmt.Println("[数据库] 迁移失败")
		return
	}

	fmt.Println("[数据库] 迁移成功")
	return
}

func database(ops *Options) (err error) {
	var cfg *m.Config
	var db *sql.DB

	cfg, err = m.ParseDSN(ops.uri)
	if err != nil {
		fmt.Println(errors.Wrap(err, "[数据库] 无效的数据库uri"))
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
		fmt.Println(errors.Wrap(err, "[数据库] 创建数据库失败"))
	}
	return
}

func releaseLock(ops *Options, db *sql.DB) (err error) {
	q := fmt.Sprintf("SELECT RELEASE_LOCK('%v')", ops.lockName)
	_, err = db.Exec(q)

	if err != nil {
		fmt.Println("[数据库] 锁释放失败")
		return
	}

	fmt.Println("[数据库] 锁释放成功")
	return
}

func acquireLock(ops *Options, db *sql.DB) (f bool, err error) {
	q := fmt.Sprintf("SELECT GET_LOCK('%v', 5)", ops.lockName)
	err = db.QueryRow(q).Scan(&f)

	if err != nil {
		fmt.Println("[数据库] 申请互斥锁失败")
	}
	fmt.Println("[数据库] 申请互斥锁成功")
	return
}

func status(ops *Options, db *sql.DB, source *migrate.EmbedFileSystemMigrationSource) (err error) {
	var migrations []*migrate.Migration
	migrations, err = source.FindMigrations()
	if err != nil {
		fmt.Println("[数据库] 没有找到数据库迁移文件")
		return
	}

	var records []*migrate.MigrationRecord
	records, err = migrate.GetMigrationRecords(db, ops.driver)
	if err != nil {
		fmt.Println("[数据库] 没有找到数据库迁移历史")
		return
	}
	rows := make(map[string]bool)
	pending := make([]string, 0)
	applied := make([]string, 0)
	for _, item := range migrations {
		rows[item.Id] = false
	}

	for _, item := range records {
		rows[item.Id] = true
	}

	for i, l := 0, len(migrations); i < l; i++ {
		if !rows[migrations[i].Id] {
			pending = append(pending, migrations[i].Id)
		} else {
			applied = append(applied, migrations[i].Id)
		}
	}

	fmt.Printf("[数据库] 迁移状态：等待中 %d，已接收 %d \n", len(pending), len(applied))
	return
}

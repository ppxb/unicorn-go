package migrate

import (
	"database/sql"
	"fmt"
	m "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/log"
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
		fmt.Println(errors.Wrap(err, "open mysql failed"))
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
			log.Error("mysql can't acquire lock,retrying...")
		}
	}

	if ops.before != nil {
		err = ops.before(ops.ctx)
		if err != nil {
			fmt.Println(errors.Wrap(err, "execute mysql before hook failed"))
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
		return
	}

	_, err = migrate.Exec(db, ops.driver, source, migrate.Up)
	if err != nil {
		log.Error("mysql migrate failed")
		return
	}
	return
}

func database(ops *Options) (err error) {
	var cfg *m.Config
	var db *sql.DB

	cfg, err = m.ParseDSN(ops.uri)
	if err != nil {
		fmt.Println(errors.Wrap(err, "invalid mysql uri"))
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
		fmt.Println(errors.Wrap(err, "create mysql database failed"))
	}
	return
}

func releaseLock(ops *Options, db *sql.DB) (err error) {
	q := fmt.Sprintf("SELECT RELEASE_LOCK('%v')", ops.lockName)
	_, err = db.Exec(q)

	if err != nil {
		log.Error("release mysql lock failed")
		return
	}
	return
}

func acquireLock(ops *Options, db *sql.DB) (f bool, err error) {
	q := fmt.Sprintf("SELECT GET_LOCK('%v', 5)", ops.lockName)
	err = db.QueryRow(q).Scan(&f)

	if err != nil {
		log.Error("mysql acquire lock failed")
	}
	return
}

func status(ops *Options, db *sql.DB, source *migrate.EmbedFileSystemMigrationSource) (err error) {
	var migrations []*migrate.Migration
	migrations, err = source.FindMigrations()
	if err != nil {
		log.Error("can't find migrate files")
		return
	}

	var records []*migrate.MigrationRecord
	records, err = migrate.GetMigrationRecords(db, ops.driver)
	if err != nil {
		log.Error("can't find migrate history")
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
	return
}

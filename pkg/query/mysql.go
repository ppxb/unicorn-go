package query

import (
	"context"
	"github.com/ppxb/unicorn/pkg/utils"
	"gorm.io/gorm"
	"reflect"
)

type Mysql struct {
	ops MysqlOptions
	Ctx context.Context
	Tx  *gorm.DB
	Db  *gorm.DB
}

func NewMysql(options ...func(*MysqlOptions)) Mysql {
	ops := getMysqlOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	if ops.db == nil {
		panic("mysql 数据库不存在")
	}
	my := Mysql{
		Db:  ops.db.WithContext(context.Background()),
		Tx:  getTx(ops.db, *ops),
		ops: *ops,
	}
	return my
}

func getTx(db *gorm.DB, ops MysqlOptions) *gorm.DB {
	tx := db
	return tx
}

func (m Mysql) Create(r interface{}, model interface{}) (err error) {
	i := model
	v := reflect.ValueOf(r)
	if v.Kind() == reflect.Slice {
		mv := reflect.Indirect(reflect.ValueOf(model))
		if mv.Kind() == reflect.Struct {
			slice := reflect.MakeSlice(reflect.SliceOf(mv.Type()), 0, 0)
			arr := reflect.New(slice.Type())
			i = arr.Interface()
		} else if mv.Kind() == reflect.Slice {
			slice := reflect.MakeSlice(mv.Type(), 0, 0)
			arr := reflect.New(slice.Type())
			i = arr.Interface()
		}
	}
	utils.Struct2StructByJson(r, i)
	err = m.Tx.Create(i).Error
	return
}

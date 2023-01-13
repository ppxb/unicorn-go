package query

import (
	"context"
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
		ops: *ops,
	}
	return my
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
	return
}

package global

import (
	"github.com/go-sql-driver/mysql"
)

type Configuration struct {
	Server ServerConfiguration `mapstructure:"server" json:"server"`
	Mysql  MysqlConfiguration  `mapstructure:"mysql" json:"mysql"`
	Redis  RedisConfiguration  `mapstructure:"redis" json:"redis"`
}

type ServerConfiguration struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	ApiVersion      string `mapstructure:"api-version" json:"apiVersion"`
	ConnectTimeout  int    `mapstructure:"connect-timeout" json:"connectTimeout"`
	CasbinModelPath string `mapstructure:"casbin-model-path" json:"casbinModelPath"`
}

type MysqlConfiguration struct {
	Uri         string       `mapstructure:"uri" json:"uri"`
	TablePrefix string       `mapstructure:"table-prefix" json:"tablePrefix"`
	ShowSql     bool         `mapstructure:"show-sql" json:"showSql"`
	DSN         mysql.Config `json:"-"`
}

type RedisConfiguration struct {
	Uri string `mapstructure:"uri" json:"uri"`
}

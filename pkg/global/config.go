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
	Port           int    `mapstructure:"port" json:"port"`
	ApiPrefix      string `mapstructure:"api-prefix" json:"apiPrefix"`
	ApiVersion     string `mapstructure:"api-version" json:"apiVersion"`
	ConnectTimeout int    `mapstructure:"connect-timeout" json:"connectTimeout"`
}

type MysqlConfiguration struct {
	Uri          string       `mapstructure:"uri" json:"uri"`
	TablePrefix  string       `mapstructure:"table-prefix" json:"tablePrefix"`
	ShowSql      bool         `mapstructure:"show-sql" json:"showSql"`
	MaxIdleConns int          `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int          `mapstructure:"max-open-conns" json:"maxOpenConns"`
	DSN          mysql.Config `json:"-"`
}

type RedisConfiguration struct {
	Uri string `mapstructure:"uri" json:"uri"`
}

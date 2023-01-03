package config

import "fmt"

type Database struct {
	Type     string `mapstructure:"type" json:"type"`
	Host     string `mapstructure:"path" json:"path"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	DbName   string `mapstructure:"db-name" json:"db-name"`
	Config   string `mapstructure:"config" json:"config"`
}

func (m *Database) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DbName, m.Config)
}

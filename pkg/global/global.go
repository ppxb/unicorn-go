package global

import "gorm.io/gorm"

var (
	Mode   string
	Config Configuration
	Mysql  *gorm.DB
)

package models

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type M struct {
	Id        uint            `gorm:"primaryKey" json:"id"`
	CreatedAt carbon.DateTime `gorm:"comment:创建时间" json:"createdAt"`
	UpdatedAt carbon.DateTime `gorm:"comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}

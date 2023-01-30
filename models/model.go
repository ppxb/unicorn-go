package models

import "github.com/golang-module/carbon/v2"

type M struct {
	Id        uint            `gorm:"primaryKey;comment:auto increment id" json:"id"`
	CreatedAt carbon.DateTime `gorm:"comment:create time" json:"createdAt"`
	UpdatedAt carbon.DateTime `gorm:"comment:update time" json:"updatedAt"`
}

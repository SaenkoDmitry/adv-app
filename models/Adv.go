package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)


type Adv struct {
	gorm.Model
	Name string `gorm:"type:varchar(1000)"`
	Price float32 `gorm:"type:decimal(10,0)"`
	Photos postgres.Jsonb `gorm:"type:jsonb"`
}

func (Adv) TableName() string {
	return "adv.adv"
}

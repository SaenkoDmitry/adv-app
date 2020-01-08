package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Adv struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(200)"`
	Price       float32        `gorm:"type:decimal(10,0)"`
	Avatar      string         `gorm:"type:text"`
	Photos      postgres.Jsonb `gorm:"type:jsonb"`
	Description string         `gorm:"type:varchar(1000)"`
}

func (Adv) TableName() string {
	return "adv.adv"
}

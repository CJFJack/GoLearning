package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(32);not null;default:''"`
	Password   string    `gorm:"type:varchar(1024);not null;default:''"`
	Birthday   time.Time `gorm:"type:date;not null"`
	Sex        bool      `gorm:"not null;default:false"`
	Tel        string    `gorm:"type varchar(16);not null;default:''"`
	Addr       string    `gorm:"type varchar(512);not null;default:''"`
	Desc       string    `gorm:"column:description;type:text;not null;default:''"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

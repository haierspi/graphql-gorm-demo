package db

import "github.com/jinzhu/gorm"

type tag struct {
	gorm.Model
	Name string `gorm:"not null;unique;index"`
}

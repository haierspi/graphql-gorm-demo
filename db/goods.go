package db

import "github.com/jinzhu/gorm"

type goods struct {
	gorm.Model
	Name        string `gorm:"not null;unique;index"`
	Description string `gorm:"not null;default:''"`
	Category    string `gorm:"not null;default:''"`
}

package db

import "time"

type goodsTag struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	GoodsID    uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
}

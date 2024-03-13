package models

import "time"

type Product struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null; type:varchar(191)"`
	Brand     string `gorm:"not null; type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
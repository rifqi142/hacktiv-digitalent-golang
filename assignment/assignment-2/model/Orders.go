package model

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	Orders_id      int       `json:"order_id" gorm:"primaryKey"`
	Customer_name  string    `json:"customer_name"`
	Ordered_at     time.Time `json:"ordered_at"`
	Items          []Items   `json:"items"`
}

func (o *Orders) BeforeCreate(tx *gorm.DB) error {
	o.Ordered_at = time.Now() 

	return nil
}

func (o *Orders) BeforeUpdate(tx *gorm.DB) error {
	o.Ordered_at = time.Now() 

	return nil
}
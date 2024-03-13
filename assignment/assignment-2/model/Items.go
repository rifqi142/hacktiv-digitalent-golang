package model

type Items struct {
	Item_id     int    `json:"item_id" gorm:"primaryKey"`
	Item_code   int    `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Orders_id   int    `json:"order_id"`
}
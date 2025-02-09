package models

type Order struct {
    OrderID   int     `json:"order_id" gorm:"primaryKey"`
    ProductID int     `json:"product_id"`
    Quantity  int     `json:"quantity"`
    OrderDate string  `json:"order_date"`
}
package models

type Inventory struct {
    ProductID int  `json:"product_id" gorm:"primaryKey"`
    Quantity  int  `json:"quantity"`
    Location  string `json:"location"`
}
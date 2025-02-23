package models

type Product struct {
    ID          int     `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Category    string  `json:"category"`
    ImagePath   string  `json:"image_path"`
}
package db

import (
    "inventory-management/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func InitDB() {
    dsn := "user:password@tcp(127.0.0.1:3306)/inventory_management?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})
}

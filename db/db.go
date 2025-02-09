package db

import (
    "dibimbing_golang_day_23_assignment/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func InitDB() {
    dsn := "root:@tcp(localhost:3306)/golang_inventory_management?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})
}

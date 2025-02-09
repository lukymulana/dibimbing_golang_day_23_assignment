package handlers

import (
    "inventory-management/db"
    "inventory-management/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&order)
    c.JSON(http.StatusCreated, order)
}

func GetOrderByID(c *gin.Context) {
    orderID := c.Param("orderID")
    var order models.Order
    if err := db.DB.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}
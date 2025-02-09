package handlers

import (
    "inventory-management/db"
    "inventory-management/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetInventoryByProductID(c *gin.Context) {
    productID := c.Param("productID")
    var inventory models.Inventory
    if err := db.DB.First(&inventory, productID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
        return
    }
    c.JSON(http.StatusOK, inventory)
}

func UpdateInventory(c *gin.Context) {
    productID := c.Param("productID")
    var inventory models.Inventory
    if err := db.DB.First(&inventory, productID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
        return
    }
    if err := c.ShouldBindJSON(&inventory); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&inventory)
    c.JSON(http.StatusOK, inventory)
}
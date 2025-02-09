package handlers

import (
    "dibimbing_golang_day_23_assignment/db"
    "dibimbing_golang_day_23_assignment/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&product)
    c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
    var products []models.Product
    db.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := db.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func GetProductsByCategory(c *gin.Context) {
    category := c.Param("category")
    var products []models.Product
    db.DB.Where("category = ?", category).Find(&products)
    c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := db.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := db.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    db.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
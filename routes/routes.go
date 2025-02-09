package routes

import (
    "inventory-management/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // Product routes
    r.POST("/products", handlers.CreateProduct)
    r.GET("/products", handlers.GetProducts)
    r.GET("/products/:id", handlers.GetProductByID)
    r.GET("/products/category/:category", handlers.GetProductsByCategory)
    r.PUT("/products/:id", handlers.UpdateProduct)
    r.DELETE("/products/:id", handlers.DeleteProduct)

    // Inventory routes
    r.GET("/inventory/:productID", handlers.GetInventoryByProductID)
    r.PUT("/inventory/:productID", handlers.UpdateInventory)

    // Order routes
    r.POST("/orders", handlers.CreateOrder)
    r.GET("/orders/:orderID", handlers.GetOrderByID)
}
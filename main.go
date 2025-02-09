package main

import (
    "inventory-management/db"
    "inventory-management/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    db.InitDB()

    // Set up the router
    r := gin.Default()

    // Define routes
    routes.SetupRoutes(r)

    // Start the server
    r.Run() // listen and serve on 0.0.0.0:8080
}
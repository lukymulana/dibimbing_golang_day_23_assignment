package main

import (
    "dibimbing_golang_day_23_assignment/db"
    "dibimbing_golang_day_23_assignment/routes"
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
    r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
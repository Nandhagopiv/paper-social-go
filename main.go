package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Nandhagopiv/paper-social-go/routes"
    "github.com/Nandhagopiv/paper-social-go/utils"
)

func main() {
    // Initialize the database connection
    utils.ConnectDatabase()

    // Create a new Gin router
    r := gin.Default()

    // Setup API routes
    routes.SetupRoutes(r)

    // Run the server on port 8080
    r.Run(":8080")
}
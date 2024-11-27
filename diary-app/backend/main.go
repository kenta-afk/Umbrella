package main

import (
    "diary/config"
    "diary/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Initialize the database
    config.InitDB()

    // Register routes
    routes.RegisterRoutes(r)

    // Start the server
    r.Run(":8080")
}
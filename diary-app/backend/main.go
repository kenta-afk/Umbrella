package main

import (
    "diary/config"
    "diary/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    
    r.Use(cors.Default())

   
    config.InitDB()

    
    routes.RegisterRoutes(r)

   
    r.Run(":8080")
}
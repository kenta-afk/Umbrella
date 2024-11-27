package routes

import (
    "diary/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/entries", controllers.GetEntries)
    r.POST("/entries", controllers.CreateEntry)
}
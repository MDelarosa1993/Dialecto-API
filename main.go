package main

import (
	"Dialecto-API/db"
	"Dialecto-API/handlers"
	"Dialecto-API/middlewares"
	"Dialecto-API/models"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()

	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration Failed:", err)
	}
	log.Println("Database migration complete!")

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
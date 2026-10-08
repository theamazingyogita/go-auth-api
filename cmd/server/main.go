package main

import (
	"go-auth-api/internal/config"
	"go-auth-api/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load() // load all the settings

	db, err := database.Connect(cfg.DB) // connect to postgres
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	log.Println("connected to the db")

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})
	log.Printf("Server running on port %s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server: %v", err)
	}
}

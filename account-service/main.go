package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rifanamrozi/account-service/handlers"
	config "github.com/rifanamrozi/account-service/pkg/config"
)

func main() {
	// Load environment variables from .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	// Initialize Supabase
	config.InitSupabase()

	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.GET("/accounts", handlers.GetAccounts)
	r.GET("/history", handlers.GetPaymentHistory)

	r.Run(":8080")
}

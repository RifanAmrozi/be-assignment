package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rifanamrozi/payment-service/handlers"
	"github.com/rifanamrozi/payment-service/pkg/config"
)

func main() {
	// Load environment variables from .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
	// // Handle DB connection
	// db, err := database.ConnectDB()
	// if err != nil {
	//  log.Fatal("Cannot connect to database")
	// }
	// // Defer disconnect until program stops
	// defer db.Client.Disconnect()

	// Initialize Prisma
	config.InitPrisma()
	r := gin.Default()

	r.POST("/send", handlers.Send)
	r.POST("/withdraw", handlers.Withdraw)

	r.Run(":8081")
}

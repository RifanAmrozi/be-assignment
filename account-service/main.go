package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rifanamrozi/account-service/handlers"
)

func main() {
    r := gin.Default()
    
    r.POST("/login", handlers.Login)
    r.POST("/register", handlers.Register)
    r.GET("/accounts", handlers.GetAccounts)
    r.GET("/history", handlers.GetPaymentHistory)

    r.Run(":8080")
}

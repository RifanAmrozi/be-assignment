package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rifanamrozi/payment-service/handlers"
)

func main() {
    r := gin.Default()
    
    r.POST("/send", handlers.Send)
    r.POST("/withdraw", handlers.Withdraw)

    r.Run(":8081")
}

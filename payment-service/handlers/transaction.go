package handlers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

func processTransaction(transaction map[string]interface{}) {
    time.Sleep(30 * time.Second)
    // Update transaction status to processed
}

func Send(c *gin.Context) {
    var transaction map[string]interface{}
    if err := c.BindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    go processTransaction(transaction)
    c.JSON(http.StatusOK, gin.H{"message": "Transaction processing started"})
}

func Withdraw(c *gin.Context) {
    var transaction map[string]interface{}
    if err := c.BindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    go processTransaction(transaction)
    c.JSON(http.StatusOK, gin.H{"message": "Transaction processing started"})
}

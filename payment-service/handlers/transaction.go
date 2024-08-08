package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	prisma "github.com/rifanamrozi/payment-service/db"
	"github.com/rifanamrozi/payment-service/pkg/config"
)

func processTransaction(transaction map[string]interface{}) {
	time.Sleep(30 * time.Second)
	// Update transaction status to processed
	ctx := context.Background()
	config.PrismaClient.Transaction.FindUnique(
		prisma.Transaction.ID.Equals(transaction["id"].(int)),
	).Update(
		prisma.Transaction.Status.Set("processed"),
	).Exec(ctx)
}

func Send(c *gin.Context) {
	var transaction struct {
		Amount    float64 `json:"amount"`
		Currency  string  `json:"currency"`
		ToAddress string  `json:"toAddress"`
		AccountId int     `json:"accountId"`
	}
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := config.PrismaClient.Transaction.CreateOne(
		prisma.Transaction.Amount.Set(transaction.Amount),
		prisma.Transaction.Currency.Set(transaction.Currency),
		prisma.Transaction.ToAddress.Set(transaction.ToAddress),
		prisma.Transaction.Status.Set("processing"),
	).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go processTransaction(map[string]interface{}{
		"id": tx.ID,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Transaction processing started", "transaction": tx})
}

func Withdraw(c *gin.Context) {
	var transaction struct {
		Amount    float64 `json:"amount"`
		Currency  string  `json:"currency"`
		ToAddress string  `json:"toAddress"`
		AccountId int     `json:"accountId"`
	}
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := config.PrismaClient.Transaction.CreateOne(
		prisma.Transaction.Amount.Set(transaction.Amount),
		prisma.Transaction.Currency.Set(transaction.Currency),
		prisma.Transaction.ToAddress.Set(transaction.ToAddress),
		prisma.Transaction.Status.Set("processing"),
	).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go processTransaction(map[string]interface{}{
		"id": tx.ID,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Transaction processing started", "transaction": tx})
}

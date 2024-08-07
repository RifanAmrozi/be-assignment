package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	"github.com/rifanamrozi/account-service/pkg/config"
)

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	auth := config.SupabaseClient.Auth
	user, err := auth.SignIn(c.Request.Context(), supa.UserCredentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func Register(c *gin.Context) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	signupOptions := supa.UserCredentials{
		Email:    user.Email,
		Password: user.Password,
	}

	auth := config.SupabaseClient.Auth
	_, err := auth.SignUp(c.Request.Context(), signupOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func GetAccounts(c *gin.Context) {
	// Implement logic to get all accounts for a user
	c.JSON(http.StatusOK, gin.H{"accounts": []string{}})
}

func GetPaymentHistory(c *gin.Context) {
	// Implement logic to get payment history for a user
	c.JSON(http.StatusOK, gin.H{"history": []string{}})
}

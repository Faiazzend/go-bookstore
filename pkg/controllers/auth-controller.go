package controllers

import (
	"net/http"
	"time"

	"github.com/Faiazzend/go-bookstore/pkg/auth"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configuredUsername := auth.GetEnv("JWT_AUTH_USERNAME", "admin")
	configuredPassword := auth.GetEnv("JWT_AUTH_PASSWORD", "password")

	if request.Username != configuredUsername || request.Password != configuredPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, expiresAt, err := auth.GenerateToken(request.Username, auth.JWTSecret(), 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
		"token_type":   "Bearer",
		"expires_at":   expiresAt.Format(time.RFC3339),
	})
}

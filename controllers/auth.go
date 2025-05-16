package controllers

import (
	"net/http"

	"github.com/dohyeoplim/blog-server/config"
	"github.com/dohyeoplim/blog-server/models"
	"github.com/dohyeoplim/blog-server/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SetupRequest struct {
	Email string `json:"email"`
}

func SetupTOTP(c *gin.Context) {
	var req SetupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	totpSetup, err := services.GenerateTOTP(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate TOTP"})
		return
	}

	user := models.User{
		ID:         uuid.New(),
		Email:      req.Email,
		TOTPSecret: totpSetup.Secret,
		IsVerified: false,
	}

	config.DB.Create(&user)
	c.Data(200, "image/png", totpSetup.QRPNG)
}

type VerifyRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func VerifyTOTP(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user models.User
	result := config.DB.First(&user, "email = ?", req.Email)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if services.ValidateTOTP(user.TOTPSecret, req.Token) {
		user.IsVerified = true
		config.DB.Save(&user)
		c.JSON(http.StatusOK, gin.H{"message": "TOTP verified!"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
	}
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/akashriva/gin_framework/helper"
)

// GenerateToken generates a JWT token for the user
func GenerateToken(c *gin.Context) {
	role := c.Param("role")
	if role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role is required"})
		return
	}

	// Example data; in a real application, you'd get this from the database
	userId := "exampleId"
	email := "example@example.com"

	token, err := helper.GenerateToken(userId, email, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

/*
 * User Register with email and password
 */

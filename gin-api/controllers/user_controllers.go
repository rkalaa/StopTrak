package controllers

import (
	"gin-api/services"
	"gin-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers fetches all users and returns as JSON
func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.JSONResponse(c, gin.H{"error": "Failed to fetch users"}, http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(c, users, http.StatusOK)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user services.UserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, gin.H{"error": "Invalid input"}, http.StatusBadRequest)
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		utils.JSONResponse(c, gin.H{"error": "Failed to create user"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(c, gin.H{"message": "User created successfully"}, http.StatusCreated)
}

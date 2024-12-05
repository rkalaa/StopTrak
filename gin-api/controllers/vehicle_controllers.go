package controllers

import (
	"gin-api/services"
	"gin-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVehicles retrieves all vehicles
func GetVehicles(c *gin.Context) {
	vehicles, err := services.GetAllVehicles()
	if err != nil {
		utils.JSONResponse(c, gin.H{"error": "Failed to fetch vehicles"}, http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(c, vehicles, http.StatusOK)
}

// GetVehicleByID retrieves a vehicle by ID
func GetVehicleByID(c *gin.Context) {
	vehicleID := c.Param("id")
	vehicle, err := services.GetVehicleByID(vehicleID)
	if err != nil {
		utils.JSONResponse(c, gin.H{"error": "Vehicle not found"}, http.StatusNotFound)
		return
	}
	utils.JSONResponse(c, vehicle, http.StatusOK)
}

// CreateVehicle creates a new vehicle
func CreateVehicle(c *gin.Context) {
	var vehicle services.VehicleInput
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		utils.JSONResponse(c, gin.H{"error": "Invalid input"}, http.StatusBadRequest)
		return
	}

	err := services.CreateVehicle(vehicle)
	if err != nil {
		utils.JSONResponse(c, gin.H{"error": "Failed to create vehicle"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(c, gin.H{"message": "Vehicle created successfully"}, http.StatusCreated)
}

// DeleteVehicle deletes a vehicle by ID
func DeleteVehicle(c *gin.Context) {
	vehicleID := c.Param("id")
	err := services.DeleteVehicle(vehicleID)
	if err != nil {
		utils.JSONResponse(c, gin.H{"error": "Failed to delete vehicle"}, http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(c, gin.H{"message": "Vehicle deleted successfully"}, http.StatusOK)
}

package controllers

import (
	"api/config"
	"api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type VisitInput struct {
	HouseID      uint   `json:"house_id" binding:"required"`
	Reason       string `json:"reason" binding:"required"`
	LicensePlate string `json:"license" binding:"required"`
}

func CreateVisit(c *gin.Context) {
	var input model.Visit

	// Attempt to bind the incoming JSON data to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	// Create a new visit record
	visit := model.Visit{
		HouseID:      input.HouseID,
		Reason:       input.Reason,
		Status:       "incoming",
		ArrivalTime:  time.Now(),
		LicensePlate: input.LicensePlate,
	}

	// Attempt to save the visit record to the database
	if err := config.DB.Create(&visit).Error; err != nil {
		// Log the error for debugging
		log.Printf("Error creating visit: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save visit"})
		return
	}

	// Send a successful response back to the client
	c.JSON(http.StatusOK, gin.H{
		"ID":           visit.ID,
		"CreatedAt":    visit.CreatedAt,
		"UpdatedAt":    visit.UpdatedAt,
		"DeletedAt":    visit.DeletedAt,
		"house_id":     visit.HouseID,
		"Reason":       visit.Reason,
		"Status":       visit.Status,
		"ArrivalTime":  visit.ArrivalTime,
		"LicensePlate": visit.LicensePlate,
	})
}

func GetAllVisit(c *gin.Context) {
	var visits []model.Visit

	if err := config.DB.Find(&visits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve visits"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"visits": visits})
}

func UpdateVisitStatus(c *gin.Context) {
	var requestData struct {
		VisitID int    `json:"visit_id"`
		Status  string `json:"status"` // Status can be 'Accepted' or 'Declined'
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	result := config.DB.Model(&model.Visit{}).Where("id = ?", requestData.VisitID).Update("status", requestData.Status)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

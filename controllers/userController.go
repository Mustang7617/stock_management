package controllers

import (
	"api/config"
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVisitByHouseID(c *gin.Context) {
	houseID := c.Query("house_id")
	// loggedInUser := c.MustGet("username").(model.User)

	// if fmt.Sprint(loggedInUser.HouseID) != houseID {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this data"})
	// 	return
	// }

	var visit []model.Visit
	if err := config.DB.Where("house_id = ?", houseID).Find(&visit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve visits"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"visit": visit})

}

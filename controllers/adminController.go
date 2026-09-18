package controllers

import (
	"api/config"
	"api/model"
	"api/utils"
	"crypto/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	HouseID string `json:"house_id" binding:"required"`
	Role    string `json:"role" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser model.User
	if err := config.DB.Where("house_id = ?", input.HouseID).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this house_id already created"})
		return
	}

	username := generateUsername(input.HouseID, input.Role)

	password := generateRandomPassword()

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	houseIDUint, err := strconv.ParseUint(input.HouseID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid house_id format"})
		return
	}

	user := model.User{
		HouseID:  uint(houseIDUint),
		UserName: username,
		Password: hashedPassword,
		Role:     input.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "User created successfully",
		"username": username,
		"password": password,
		"house_id": input.HouseID,
	})
}

func generateUsername(houseID string, role string) string {
	return strings.ToLower(houseID) + "_" + strings.ToLower(role)
}

func generateRandomPassword() string {
	const charset = "0123456789"
	randomBytes := make([]byte, 6)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "123456"
	}

	for i := range randomBytes {
		randomBytes[i] = charset[randomBytes[i]%byte(len(charset))]
	}

	return string(randomBytes)
}

func DeleteUser(c *gin.Context) {
	houseID := c.Param("house_id")

	var user model.User
	if err := config.DB.Where("house_id = ?", houseID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User was deleted"})
}

func GetAllUsers(c *gin.Context) {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

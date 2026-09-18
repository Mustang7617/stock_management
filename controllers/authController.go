package controllers

import (
	"net/http"
	"strconv"
	"time"

	"api/config"
	"api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key") // Use a strong key, ideally from an env variable

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	HouseID  string `json:"house_id"` // Keep HouseID as a string for JWT claims
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := config.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Create a new token
	expirationTime := time.Now().Add(1 * time.Hour) // Set token expiration
	claims := &Claims{
		Username: user.UserName,
		Role:     user.Role,
		HouseID:  strconv.Itoa(int(user.HouseID)), // Convert uint to string
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	// Return the token, role, house_id, and username
	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"role":     user.Role,
		"house_id": claims.HouseID,
		"username": claims.Username,
	})
}

package main

import (
	"api/config"
	// "api/middleware"

	"api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// r.Use(middleware.EnableCORS())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:38959"}, // Allow your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
	}))

	config.ConnectDatabase()
	routes.SetupRouter(r)

	r.Run(":8001")

}

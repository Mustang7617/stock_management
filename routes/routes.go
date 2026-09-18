package routes

import (
	"api/controllers"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func SetupRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	admin := r.Group("/admin")
	{
		admin.POST("/create", controllers.CreateUser)
		admin.GET("/getuser", controllers.GetAllUsers)
		admin.DELETE("/delete/:house_id", controllers.DeleteUser)
	}

	user := r.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Access granted"})
		})
		user.GET("/getvisit/:house_id", controllers.GetVisitByHouseID)

		user.PUT("/updatestatus", controllers.UpdateVisitStatus)
	}

	guard := r.Group("/guard")
	{
		guard.POST("/CreateVisit", controllers.CreateVisit)
		guard.GET("/GetAllvisit", controllers.GetAllVisit)
	}

}

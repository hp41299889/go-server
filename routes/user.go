package routes

import (
	"go-server/controllers"

	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine) {
	router.Group("/api/v0").
		POST("/user", controllers.CreateUser()).
		GET("/users", controllers.ReadUsers())
	// GET("/user/:userId", controllers.ReadUser()).
	// GET("/users", controllers.ReadUsers()).
	// PUT("/user:userId", controllers.UpdateUser()).
	// DELETE("/user:userId", controllers.DeleteUser())
}

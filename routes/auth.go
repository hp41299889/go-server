package routes

import (
	"go-server/controllers"

	"github.com/gin-gonic/gin"
)

func Login(router *gin.Engine) {
	router.Group("/api/v0").
		POST("/login", controllers.Login())
}

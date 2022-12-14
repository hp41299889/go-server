package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-server/docs"
	"go-server/routes"
	"go-server/services"
)

func main() {
	router := gin.Default()

	services.PostgresConnect()
	routes.User(router)
	routes.Login(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}

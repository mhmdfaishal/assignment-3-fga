package router

import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetIndex)

	router.Static("/static/", "./assets")
	
	
	return router
}

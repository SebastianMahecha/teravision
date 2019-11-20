package main

import (	
	"github.com/gin-gonic/gin"	
	"github.com/teravision/controllers"
	"github.com/teravision/config"
)

func main() {
	
	config.InitDB()
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the api
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/user", controllers.SaveUser)		
	
	// Start and run the server
	router.Run(":3000")
	
}
	
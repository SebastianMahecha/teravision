package main

import (	
	"github.com/gin-gonic/gin"	
	"github.com/teravision/controllers"
	"github.com/teravision/models"
	"github.com/teravision/config"
)

func main() {
	
	//open database connection
	config.InitDB()
	//migrate models
	Migrate()
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the api
	api := router.Group("/api")
	// Setup route group for the v1
	v1 := api.Group("/v1")
	
	// endpoints list
	v1.POST("/user", controllers.SaveUser)		
	
	// Start and run the server
	router.Run(":3000")	
}

//Function to run model migration
func Migrate() {
	config.DB.AutoMigrate(models.User{})
}

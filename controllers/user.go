package controllers
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
//Controller for saver user information 
func SaveUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{			
		},
	})

	return
}
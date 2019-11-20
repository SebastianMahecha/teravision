package controllers
import (	
	"net/http"
	"github.com/teravision/httpmodels"
	"github.com/teravision/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)
//Controller for saver user information 
func SaveUser(c *gin.Context)  {
	
	u := &httpmodels.UserRequest{}

	if err := c.ShouldBindBodyWith(u, binding.JSON); err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"data": gin.H{
				"error_code": "incomplete_data",
				"message": "Send complete data please",
			},
		})
		return
	}

	err := helpers.CreateUser(u)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"data": gin.H{
				"error_code": "general_error",
				"message": "User not saved successfully. Error: "+err.Error(),
			},
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{	
				"message": "User saved successfully",		
			},
		})
	}
	
	return
}

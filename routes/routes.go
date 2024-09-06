package routes

import (
	"github.com/Luqky07/user_api/controllers"
	"github.com/gin-gonic/gin"
)

// Setup the api routes
func SetupRoutes(router *gin.Engine) {
	//Client creation
	router.POST("/users", controllers.CreateNewClient)
}

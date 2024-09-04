package routes

import (
	"github.com/Luqky07/user_api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", controllers.Hello)

	router.POST("/users", controllers.CreateNewUser)
}

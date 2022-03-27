package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/src/backend/controllers"
	"github.com/rabbice/ecommerce/src/backend/middleware"
)

func UserRoutes(v1 *gin.Engine) {
	v1.POST("/auth/login", controllers.Login())
	v1.POST("/auth/signup", controllers.SignUp())
	v1.Use(middleware.Authenticate())
	v1.POST("/create/address", controllers.AddAddress())
	v1.DELETE("address/:id", controllers.DeleteAddress())
}

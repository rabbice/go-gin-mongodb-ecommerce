package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
	"github.com/rabbice/ecommerce/middleware"
)

func UserRoutes(v1 *gin.Engine) {
	v1.POST("/auth/login", controllers.Login())
	v1.POST("/auth/signup", controllers.SignUp())
	v1.Use(middleware.Authenticate())
	v1.POST("/create/address", controllers.AddAddress())
	v1.PUT("/address/:id", controllers.EditAddress())
	v1.DELETE("address/:id", controllers.DeleteAddress())
	v1.PATCH("/delivery", controllers.ConfirmDelivery())
}

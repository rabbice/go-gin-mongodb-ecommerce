package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
)

func UserRoutes(v1 *gin.Engine) {
	v1.POST("/auth/login", controllers.Login())
	v1.POST("/auth/signup", controllers.SignUp())
}

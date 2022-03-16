package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/product", controllers.AddProduct())
	incomingRoutes.GET("/products", controllers.GetProducts())
	incomingRoutes.GET("/product/:product_id", controllers.GetProduct())
	incomingRoutes.DELETE("/product/:product_id", controllers.DeleteProduct())
	incomingRoutes.PUT("/product/:product_id", controllers.UpdateProduct())
}

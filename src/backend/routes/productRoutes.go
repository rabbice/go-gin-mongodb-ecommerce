package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/src/backend/controllers"
	"github.com/rabbice/ecommerce/src/backend/middleware"
)

func ProductRoutes(v1 *gin.Engine) {
	v1.GET("/products", controllers.GetProducts())
	v1.GET("/product/:product_id", controllers.GetProduct())
	v1.Use(middleware.Authenticate())
	v1.POST("/product", controllers.AddProduct())
	v1.DELETE("/product/:product_id", controllers.DeleteProduct())
	v1.PUT("/product/:product_id", controllers.UpdateProduct())
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/src/backend/controllers"
)

func ProductRoutes(v1 *gin.Engine) {
	v1.GET("/products", controllers.GetProducts())
	v1.GET("/product/:product_id", controllers.GetProduct())
	v1.GET("/product/search", controllers.SearchProductByQuery())

}

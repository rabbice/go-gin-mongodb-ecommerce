package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
	"github.com/rabbice/ecommerce/middleware"
)

func SellerRoutes(v1 *gin.Engine) {
	v1.Use(middleware.Authenticate())
	v1.POST("/shop", controllers.AddShop())
	v1.DELETE("/shop/:shop_id", controllers.DeleteShop())
	v1.PUT("/shop/:shop_id", controllers.UpdateShop())
	v1.POST("/product", controllers.AddProduct())
	v1.DELETE("/product/:product_id", controllers.DeleteProduct())
	v1.PUT("/product/:product_id", controllers.UpdateProduct())
	v1.POST("/ship", controllers.CreateDelivery())

}

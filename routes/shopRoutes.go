package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
)

func ShopRoutes(v1 *gin.Engine) {
	v1.POST("/shop", controllers.AddShop())
	v1.GET("/shop/:shop_id", controllers.GetShop())
	v1.GET("/shops", controllers.GetShops())
	v1.DELETE("/shop/:shop_id", controllers.DeleteShop())
	v1.PUT("/shop/:shop_id", controllers.UpdateShop())
}

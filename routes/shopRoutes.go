package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
)

func ShopRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/shop", controllers.AddShop())
	incomingRoutes.GET("/shop/:shop_id", controllers.GetShop())
	incomingRoutes.GET("/shops", controllers.GetShops())
	incomingRoutes.DELETE("/shop/:shop_id", controllers.DeleteShop())
	incomingRoutes.PUT("/shop/:shop_id", controllers.UpdateShop())
}

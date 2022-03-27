package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/src/backend/controllers"
)

func ShopRoutes(v1 *gin.Engine) {
	v1.GET("/shops", controllers.GetShops())
	v1.GET("/shop/:shop_id", controllers.GetShop())
}

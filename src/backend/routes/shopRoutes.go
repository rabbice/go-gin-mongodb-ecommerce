package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/src/backend/controllers"
	"github.com/rabbice/ecommerce/src/backend/middleware"
)

func ShopRoutes(v1 *gin.Engine) {
	v1.GET("/shop/:shop_id", controllers.GetShop())
	v1.GET("/shops", controllers.GetShops())
	v1.Use(middleware.Authenticate())
	v1.POST("/shop", controllers.AddShop())
	v1.DELETE("/shop/:shop_id", controllers.DeleteShop())
	v1.PUT("/shop/:shop_id", controllers.UpdateShop())
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	routes.ShopRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":5000")
}

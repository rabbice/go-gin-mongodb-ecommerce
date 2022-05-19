package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/controllers"
	"github.com/rabbice/ecommerce/database"
	"github.com/rabbice/ecommerce/middleware"
	"github.com/rabbice/ecommerce/routes"
)

var cartHandler *controllers.CartHandler

func main() {
	cartHandler = controllers.NewCartHandler(database.OpenCollection(database.Client, "product"), database.OpenCollection(database.Client, "user"))
	store := cookie.NewStore([]byte("secret"))

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(sessions.Sessions("marketplace_api", store))
	router.Use(cors.Default())
	router.Use(middleware.PrometheusMiddleware())

	srv := &http.Server{
		Addr:           ":5000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	routes.AdminRoutes(router)
	routes.ProductRoutes(router)
	routes.ShopRoutes(router)
	routes.UserRoutes(router)

	routes.SellerRoutes(router)

	router.GET("/cart", cartHandler.AddToCart())
	router.GET("/remove/cart", cartHandler.RemoveFromCart())
	router.GET("/order", cartHandler.BuyItem())
	

	log.Println("Starting server on port :5000...")
	srv.ListenAndServe()
}

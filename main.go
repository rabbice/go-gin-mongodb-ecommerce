package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/src/backend/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	srv := &http.Server{
		Addr:           ":5000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	routes.ProductRoutes(router)
	routes.ShopRoutes(router)
	routes.UserRoutes(router)

	routes.SellerRoutes(router)

	log.Println("Starting server on port :5000...")
	srv.ListenAndServe()
}

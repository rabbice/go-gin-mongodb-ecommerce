package main

import (
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

	routes.UserRoutes(router)
	routes.ProductRoutes(router)

	routes.ShopRoutes(router)

	srv.ListenAndServe()
}

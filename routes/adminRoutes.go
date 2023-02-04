package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func AdminRoutes(v1 *gin.Engine) {
	v1.GET("/prometheus", gin.WrapH(promhttp.Handler()))
}
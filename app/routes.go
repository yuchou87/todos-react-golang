package main

import (
	"github.com/gin-gonic/gin"
)

func applyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", pingHandle)
	}
}

func setRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		applyRoutes(api)
	}
}

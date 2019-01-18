package main

import (
	"github.com/gin-gonic/gin"
	"lovehome/gin_demo/apis"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	apiGroup := router.Group("/api/v1.0")

	{
		apiGroup.GET("/areas", apis.GetArea)
	}

	return router
}

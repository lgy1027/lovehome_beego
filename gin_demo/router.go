package main

import (
	"github.com/gin-gonic/gin"
	"lovehome/gin_demo/apis"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/api/v1.0/areas", apis.GetArea)

	return router
}

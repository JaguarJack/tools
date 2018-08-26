package main

import (
	"api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	routers.Route()
	routers.Start(9092)
}
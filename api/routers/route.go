package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"api/controllers"
	"net/http"
	"fmt"
)

var router = gin.Default()
var config = cors.DefaultConfig()


func Route(){
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	config.AllowMethods = []string{"PUC", "POST", "GET", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))
	router.Use(cors.Default())

	//获取主分类
	router.GET("getCategory", controllers.GetCategory)

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
}

func Start(port int) {
	router.Run(fmt.Sprintf(":%d", port))
}
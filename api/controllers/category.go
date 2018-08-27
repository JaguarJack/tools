package controllers

import (
	"github.com/gin-gonic/gin"
	"api/models"
	"net/http"
)

func GetCategory(c *gin.Context) {
	category, err := models.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": err.Error(),
			"data": "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": category,
	})

	return
}

func GetTools(c *gin.Context) {
	name := c.PostForm("name")

	if (name == "all") {
		tools, err := models.GetToolsByCName("")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": err.Error(),
				"data": "",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
			"data": tools,
		})
	} else {
		tools, err := models.GetToolsByCName(name)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": err.Error(),
				"data": "",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
			"data": tools,
		})
	}




}


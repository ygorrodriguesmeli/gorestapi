package controllers

import "github.com/gin-gonic/gin"

func ShowProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

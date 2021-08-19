package controllers

import (
	"strconv"

	"github.com/ygorrodriguesmeli/gorestapi/src/database"
	"github.com/ygorrodriguesmeli/gorestapi/src/models"

	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabase()

	var product models.Product
	err = db.First(&product, parsedId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "product not found: " + err.Error(),
		})
		return
	}

	c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetDatabase()
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json: " + err.Error(),
		})
		return
	}
	createErr := db.Create(&product).Error
	if createErr != nil {
		c.JSON(404, gin.H{
			"error": "cannot create product: " + createErr.Error(),
		})
		return
	}
	c.JSON(200, product)
}

func ListProducts(c *gin.Context) {
	db := database.GetDatabase()
	var products []models.Product
	err := db.Find(&products).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot list products: " + err.Error(),
		})
		return
	}
	c.JSON(200, products)
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/ygorrodriguesmeli/gorestapi/src/database"
	"github.com/ygorrodriguesmeli/gorestapi/src/models"
	"github.com/ygorrodriguesmeli/gorestapi/src/services"

	"github.com/gin-gonic/gin"
)

var (
	ProductController = productController{}
)

type productController struct{}

func (controller *productController) GetProductById(c *gin.Context) {
	product, err := services.ProductService.GetProductById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (controller *productController) CreateProduct(c *gin.Context) {
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

func (controller *productController) ListProducts(c *gin.Context) {
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

func (controller *productController) UpdateProduct(c *gin.Context) {
	db := database.GetDatabase()
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json: " + err.Error(),
		})
		return
	}
	createErr := db.Save(&product).Error
	if createErr != nil {
		c.JSON(404, gin.H{
			"error": "cannot update product: " + createErr.Error(),
		})
		return
	}
	c.JSON(200, product)
}

func (controller *productController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Product{}, parsedId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "product not found: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

package server

import (
	"github.com/ygorrodriguesmeli/gorestapi/src/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/product/:id", controllers.ProductController.GetProductById)
	router.GET("/product", controllers.ProductController.ListProducts)
	router.POST("/product", controllers.ProductController.CreateProduct)
	router.PUT("/product", controllers.ProductController.UpdateProduct)
	router.DELETE("/product/:id", controllers.ProductController.DeleteProduct)
	return router
}

package server

import (
	"github.com/ygorrodriguesmeli/gorestapi/src/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/product/:id", controllers.GetProductById)
	router.GET("/product", controllers.ListProducts)
	router.POST("/product", controllers.CreateProduct)
	router.PUT("/product", controllers.UpdateProduct)
	router.DELETE("/product/:id", controllers.DeleteProduct)
	return router
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ngvanloi/ecommerce-api/controllers"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetAllProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}

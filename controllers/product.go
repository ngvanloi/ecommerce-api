package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngvanloi/ecommerce-api/database"
	"github.com/ngvanloi/ecommerce-api/models"
)

// GetAllProducts lấy tất cả sản phẩm
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct tạo một sản phẩm mới
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Name: input.Name, Description: input.Description, Price: input.Price, Quantity: input.Quantity}
	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductByID lấy sản phẩm theo ID
func GetProductByID(c *gin.Context) {
	var product models.Product
	if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct cập nhật thông tin sản phẩm
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&product).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct xóa sản phẩm
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngvanloi/ecommerce-api/database"
	"github.com/ngvanloi/ecommerce-api/routes"
)

func main() {
	// Kết nối đến cơ sở dữ liệu
	database.ConnectDatabase()

	// Khởi tạo router Gin
	r := gin.Default()

	// Đăng ký routes
	routes.ProductRoutes(r)

	// Chạy server trên cổng 8080
	r.Run(":8080")
}

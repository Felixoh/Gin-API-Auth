package main

import (
	"github.com/Felixoh/handler"
	"github.com/Felixoh/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)

	api := r.Group("/api")

	api.Use(middleware.ValidateToken())
	product := api.Group("/product")
	product.Use(middleware.Authorization([]int{1}))
	product.GET("/", handler.GetAll)
	product.POST("/", middleware.Authorization([]int{4}), handler.AddProduct)

	r.Run("localhost:8080")
}

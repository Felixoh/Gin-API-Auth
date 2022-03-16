package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Id   int
	Name string
}

func GetAll(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusOK, []Products{
		{
			Id:   1,
			Name: "Product 1",
		},
		{
			Id:   2,
			Name: "Product 2",
		},
	})
}

func AddProduct(context *gin.Context) {
	var product Products
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	product.Id = 10
	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product has been added", "product": product})

}

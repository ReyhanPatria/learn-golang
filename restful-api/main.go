package main

import (
	"net/http"
	"reyhan/restful-api/model"

	"github.com/gin-gonic/gin"
)

var products = []model.Product{
	{ID: 1, Name: "Apple", Price: 23000, Description: "Apple fresh from Aomori Prefecture"},
	{ID: 2, Name: "Banana", Price: 25000, Description: "Banana from the Banana Republic"},
	{ID: 3, Name: "Steam", Price: 0, Description: "It's the software, not actual steam"},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func main() {
	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		getProducts(c)
	})

	router.Run("localhost:8080")
}

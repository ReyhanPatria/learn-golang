package main

import (
	"net/http"
	"reyhan/restful-api/model"
	"strconv"

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

func getProductById(c *gin.Context) {
	// Get ID from path parameter
	id, err := strconv.Atoi(c.Param("id"))

	// Check if ID is integer
	// If not then return 400 bad request
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Search for product with matching ID
	for _, p := range products {
		// If found product with matching ID, return 200 ok with product json
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	// If no matching ID found, return 404 not found
	c.AbortWithStatus(http.StatusNotFound)
}

func insertProduct(c *gin.Context) {
	var newProduct model.Product

	err := c.BindJSON(&newProduct)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusOK, newProduct)
}

func main() {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.POST("/product", insertProduct)

	router.GET("/product/:id", getProductById)

	router.Run("localhost:8080")
}

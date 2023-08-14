package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type product struct {
	ID   int
	Name string
}

var products = make(map[int]product)

func SetRoutes() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/* -- Unauthenticated Endpoints -- */
	unauthenticated := router.Group("/api")
	unauthenticated.GET("/register", getRegister)

	/* -- Authenticated Endpoints -- */
	authenticated := router.Group("/api")
	authenticated.Use(APIAuthenticator())
	authenticated.GET("/product/:id", getProduct)
	authenticated.POST("/product", postProduct)
	authenticated.DELETE("/product:id", deleteProduct)

	router.Run()
}

func getRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"apikey": GenerateAPIKey()})
}

func getProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse integer ID."})
	}

	c.JSON(http.StatusOK, products[id])
}

func postProduct(c *gin.Context) {
	var json product
	c.MustBindWith(&json, binding.JSON)
	products[json.ID] = json
	c.JSON(200, gin.H{"message": "Product created."})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse integer ID."})
	}

	delete(products, id)
	c.JSON(200, gin.H{"message": "Product deleted."})
}

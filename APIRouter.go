package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	//authenticated.Use(APIAuthenticator())
	authenticated.GET("/pokemon/:name", getPokemonByName)
	authenticated.DELETE("/pokemon/:name", deletePokemonByName)

	router.Run()
}

func getRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"apikey": GenerateAPIKey()})
}

func getPokemonByName(c *gin.Context) {
	name := c.Param("name")

	pokemon := SelectPokemonByName(name)
	if pokemon == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to select pokemon from database."})
	}
	c.JSON(http.StatusOK, pokemon)
}

func deletePokemonByName(c *gin.Context) {
	name := c.Param("name")
	err := DeletePokemonByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete pokemon from database."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted pokemon from database."})
}

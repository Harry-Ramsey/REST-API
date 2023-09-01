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
	authenticated.Use(APIAuthenticator())
	authenticated.GET("/pokemon/:name", getPokemon)

	router.Run()
}

func getRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"apikey": GenerateAPIKey()})
}

func getPokemon(c *gin.Context) {
	name := c.Param("name")

	pokemon := SelectPokemonByName(name)
	if pokemon == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select pokemon from database."})
	}
	c.JSON(http.StatusOK, pokemon)
}

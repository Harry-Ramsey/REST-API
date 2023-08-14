package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func APIAuthenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		APIKey := c.GetHeader("Authorization")
		if !validateAPIKey(APIKey) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is not a valid API Key."})
			return
		}
		c.Next()
	}
}

func GenerateAPIKey() string {
	buffer := make([]byte, 128)
	_, err := rand.Read(buffer)
	if err != nil {
		fmt.Println("Failed to generate random API key")
		// Add error handling here...
	}

	key := base64.StdEncoding.EncodeToString(buffer)
	rdb.Set(ctx, key, 0, 0)
	return key
}

func validateAPIKey(key string) bool {
	exist, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return exist == 1
}

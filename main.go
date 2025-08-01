package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":    []any{1, "ronak", 3},
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

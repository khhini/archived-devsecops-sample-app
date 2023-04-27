package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Worlds!.",
		})
	})
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!.", name),
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	router.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"check": os.Getenv("DB_USERNAME"),
		})
	})
	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}

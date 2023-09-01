package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		delay, _ := strconv.Atoi(c.Query("delay"))
		time.Sleep(time.Duration(delay) * time.Millisecond)

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Worlds!.",
		})
	})
	router.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!.", name),
		})
	})
	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}

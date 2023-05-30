package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Notifier!.",
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
		ctx := context.Background()
		client, err := secretmanager.NewClient(ctx)
		if err != nil {
			log.Fatalf("faild to setup client: %v", err)
		}
		defer client.Close()

		accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
			Name: os.Getenv("DB_PASSWORD_SECRET_MANAGER_REF"),
		}

		result, err := client.AccessSecretVersion(ctx, accessRequest)
		if err != nil {
			log.Fatalf("failed to access secrets version: %v", err)
		}
		secrets := fmt.Sprintf("%s", result.Payload.Data)
		c.JSON(http.StatusOK, gin.H{
			"secrets": secrets,
		})
	})
	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}

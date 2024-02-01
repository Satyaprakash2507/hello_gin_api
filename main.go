package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the gin package
	// Add blank import for functions-framework-go
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Define your API endpoints
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the server
	r.Run(":8080") // You can specify the port you want to run on
}

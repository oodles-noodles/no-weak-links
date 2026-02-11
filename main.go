package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World! This is a simple Go application.")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})

	port := "8080"
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

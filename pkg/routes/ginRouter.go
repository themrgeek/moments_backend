package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.1"}) // Adjust for your network
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ // Corrected status code
			"message": "Server is running...",
		})
	})
	return r
}

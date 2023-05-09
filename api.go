package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ipAddr := os.Getenv("IP_ADDRESS")
	if ipAddr == "" {
		ipAddr = "127.0.0.1"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.SetTrustedProxies(nil)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("%s:%s", ipAddr, port)) // listen and serve on 127.0.0.1:8080
}

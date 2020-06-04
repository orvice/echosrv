package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-swagger/go-swagger/examples/stream-server/restapi"
	_ "github.com/go-swagger/go-swagger/examples/stream-server/restapi/operations"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

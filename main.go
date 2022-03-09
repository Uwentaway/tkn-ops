package main

import (
	"devops-hello-world/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		sum := &pkg.Sum{
			Num1: 100,
			Num2: 200,
		}
		res := sum.Add()

		c.JSON(http.StatusOK, gin.H{
			"data": res,
			"say":  "Hello World",
		})
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Running",
		})
	})
	r.Run(":8080")
}

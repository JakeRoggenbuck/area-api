package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"math"
	"strconv"
	"fmt"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Group("/api/v1")
	{
		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to Area API.",
			})
		})

		r.GET("/circle/:r", func(c *gin.Context) {
			r := c.Param("r")
			i, err := strconv.ParseFloat(r, 32)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Could not calculate.",
				})
			}
			
			c.JSON(http.StatusOK, gin.H{
				"message": math.Pow(2 * math.Pi, i),
			})
		})
	}

	port := ":8080"
	fmt.Println("http://localhost" + port)
	r.Run(port)
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
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
			radius, err := strconv.ParseFloat(r, 32)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Could not calculate.",
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"message": math.Pow(radius, 2.0) * math.Pi,
			})
		})

		r.GET("/square/:w", func(c *gin.Context) {
			w := c.Param("w")
			width, err := strconv.ParseFloat(w, 32)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Could not calculate.",
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"message": width * width,
			})
		})

		r.GET("/rectangle/:h/:w", func(c *gin.Context) {
			h := c.Param("h")
			w := c.Param("w")
			width, err := strconv.ParseFloat(w, 32)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Could not calculate.",
				})
			}
			height, err := strconv.ParseFloat(h, 32)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Could not calculate.",
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"message": height * width,
			})
		})
	}

	port := ":8080"
	fmt.Println("http://localhost" + port)
	r.Run(port)
}

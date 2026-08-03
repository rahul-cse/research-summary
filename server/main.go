package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Upload endpoint working",
		})
	})

	router.Run(":7070")
}

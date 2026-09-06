package main

import (
	"os"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"github.com/rahul-cse/research-summary/pdf"

	"github.com/rahul-cse/research-summary/text"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No file uploaded",
			})
			return
		}

		uploadedFile, err := file.Open()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not open uploaded file",
			})
			return
		}

		defer uploadedFile.Close()

		rawText, err := pdf.ExtractText(uploadedFile)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not extract PDF text",
			})
			return
		}

		cleanedText := text.Clean(rawText)

		c.JSON(http.StatusOK, gin.H{
			"message":  "File uploaded successfully",
			"filename": file.Filename,
			"text":     cleanedText,
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "7070"
	}
	router.Run("0.0.0.0:" + port)
}

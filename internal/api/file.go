package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(router *gin.RouterGroup) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
			return
		}

		userDir := filepath.Join("uploads", username.(string))
		if _, err := os.Stat(userDir); os.IsNotExist(err) {
			err := os.MkdirAll(userDir, os.ModePerm)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user directory"})
				return
			}
		}

		filePath := filepath.Join(userDir, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	}

	router.POST("/upload", handler)
}

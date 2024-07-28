package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"phto/internal/config"
	"phto/internal/entity"
	"phto/internal/syncbyte"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadPhoto(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
			return
		}

		albumIDStr := c.PostForm("album_id")
		if albumIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Album ID is required"})
			return
		}
		albumID, err := strconv.Atoi(albumIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
			return
		}

		album, err := entity.FindAlbumByID(uint(albumID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
			return
		}

		filename := generateUniqueFilename(file.Filename)

		userDir := filepath.Join(conf.StoragePath, "uploads", username.(string), album.Name)
		if _, err := os.Stat(userDir); os.IsNotExist(err) {
			err := os.MkdirAll(userDir, os.ModePerm)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user directory"})
				return
			}
		}

		if err := c.SaveUploadedFile(file, filepath.Join(userDir, filename)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		photo := &entity.Photo{
			Name:     filepath.Base(file.Filename),
			FileName: filepath.Base(filename),
			FileSize: file.Size,
			FileType: strings.Split(file.Header.Get("Content-Type"), ";")[0],
		}

		pho, err := photo.Create(uint(albumID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Photo %d create successfully", pho.ID)})
	}

	router.POST("/photo/upload", handler)
}

func generateUniqueFilename(originalFilename string) string {
	return originalFilename + "_" + time.Now().Format("2006-01-02_15-04-05")
}

func ImportPhoto(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
			return
		}

		if err := syncbyte.ImportOriginals(username.(string), conf); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to import photo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "import successfully"})
	}

	router.POST("/photo/import", handler)
}

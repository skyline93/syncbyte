package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"phto/internal/config"
	"phto/internal/entity"
	"phto/internal/syncbyte"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadPhoto(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
			return
		}

		user := entity.FindUser(username.(string))

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

		var album *entity.Album
		for _, alb := range user.Albums {
			if alb.ID == uint(albumID) {
				album = &alb
				break
			}
		}

		if album == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
			return
		}

		uniqueFileName := syncbyte.GenerateUniqueFilename(file.Filename)

		photo := &entity.Photo{
			Name:     filepath.Base(file.Filename),
			FileName: filepath.Base(uniqueFileName),
			FileSize: file.Size,
			FileType: strings.Split(file.Header.Get("Content-Type"), ";")[0],
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "open src file failed"})
			return
		}
		defer src.Close()

		if err := syncbyte.UploadPhoto(username.(string), album.Name, uniqueFileName, src, conf); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
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

		if err := syncbyte.ImportOriginalsFromWebDAV(username.(string), conf); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to import photo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "import successfully"})
	}

	router.POST("/photo/import", handler)
}

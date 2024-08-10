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

// UploadPhoto godoc
//
//	@Summary		Upload photo
//	@Description	upload photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Router			/api/v1/photo/upload [post]
//	@Param			album_id	formData	int		true	"album id"
//	@Param			file		formData	file	true	"the file to upload"
//	@Success		200			{object}	Response
func UploadPhoto(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, Error(400, "User context not found"))
			return
		}

		user := entity.FindUser(username.(string))

		albumIDStr := c.PostForm("album_id")
		if albumIDStr == "" {
			c.JSON(http.StatusBadRequest, Error(400, "Album ID is required"))
			return
		}
		albumID, err := strconv.Atoi(albumIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error(400, "Invalid album ID"))
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
			c.JSON(http.StatusBadRequest, Error(400, "Invalid album ID"))
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, Error(400, "No file provided"))
			return
		}

		uniqueFileName := syncbyte.GenerateUniqueFilename(file.Filename)

		photo := &entity.Photo{
			Name:     filepath.Base(file.Filename),
			FileName: filepath.Base(uniqueFileName),
			FileSize: file.Size,
			FileType: strings.Split(file.Header.Get("Content-Type"), ";")[0],
			AlbumID:  album.ID,
			UserID:   user.ID,
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error(400, "open src file failed"))
			return
		}
		defer src.Close()

		if err := syncbyte.UploadPhoto(username.(string), album.Name, uniqueFileName, src, conf); err != nil {
			c.JSON(http.StatusInternalServerError, Error(400, "Failed to save file"))
			return
		}

		pho, err := photo.Create(uint(albumID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error(400, "Failed to create photo"))
			return
		}

		c.JSON(http.StatusOK, Success(fmt.Sprintf("Photo %d create successfully", pho.ID)))
	}

	router.POST("/photo/upload", handler)
}

// ImportPhoto godoc
//
//	@Summary		Import photo
//	@Description	import photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Router			/api/v1/photo/import [post]
//	@Success		200	{object}	Response
func ImportPhoto(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, Error(400, "User context not found"))
			return
		}

		if err := syncbyte.ImportOriginals(username.(string), conf); err != nil {
			c.JSON(http.StatusInternalServerError, Error(400, "Failed to import photo"))
			return
		}

		if err := syncbyte.ImportOriginalsFromWebDAV(username.(string), conf); err != nil {
			c.JSON(http.StatusInternalServerError, Error(400, "Failed to import photo"))
			return
		}

		c.JSON(http.StatusOK, Success("import successfully"))
	}

	router.POST("/photo/import", handler)
}

// GetPhotos godoc
//
//	@Summary		Get photos
//	@Description	get photos
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Router			/api/v1/photo [get]
//	@Param			album_id	query		int	true	"album id"
//	@Success		200			{object}	Response
func ListPhotos(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, Error(400, "User context not found"))
			return
		}

		user := entity.FindUser(username.(string))

		albumID, _ := strconv.Atoi(c.Query("album_id"))

		var album *entity.Album
		for _, alb := range user.Albums {
			if alb.ID == uint(albumID) {
				album = &alb
				break
			}
		}

		if album == nil {
			c.JSON(http.StatusBadRequest, Error(400, "Invalid album ID"))
			return
		}

		photos, err := entity.ListPhotos(username.(string), album.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error(400, "Failed to import photo"))
			return
		}

		c.JSON(http.StatusOK, Success(photos))
	}

	router.GET("/photo", handler)
}

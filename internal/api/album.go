package api

import (
	"fmt"
	"net/http"
	"phto/internal/config"
	"phto/internal/entity"

	"github.com/gin-gonic/gin"
)

func CreateAlbum(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
			return
		}

		user := entity.FindUser(username.(string))
		if user == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		var json struct {
			AlbumName string `json:"album_name" binding:"required"`
		}

		if c.Bind(&json) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if entity.ExistsAlbum(json.AlbumName, user.Name) {
			c.JSON(http.StatusConflict, gin.H{"error": "Album already exists"})
			return
		}

		album := entity.Album{Name: json.AlbumName}
		alb, err := album.Create(json.AlbumName, user.ID)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Album create failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("album %d created successfully", alb.ID)})
	}

	router.POST("/albums", handler)
}

func GetAlbums(router *gin.RouterGroup, conf *config.Config) {
	handler := func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
			return
		}

		user := entity.FindUser(username.(string))
		if user == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		albums, err := entity.ListAlbumsByUserID(user.ID)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Albums get failed"})
			return
		}

		c.JSON(http.StatusOK, albums)
	}

	router.GET("/albums", handler)
}

package server

import (
	"encoding/base64"
	"net/http"
	"path/filepath"
	"phto/internal/config"
	"phto/internal/entity"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
)

const (
	MethodHead      = "HEAD"
	MethodGet       = "GET"
	MethodPut       = "PUT"
	MethodPost      = "POST"
	MethodPatch     = "PATCH"
	MethodDelete    = "DELETE"
	MethodOptions   = "OPTIONS"
	MethodMkcol     = "MKCOL"
	MethodCopy      = "COPY"
	MethodMove      = "MOVE"
	MethodLock      = "LOCK"
	MethodUnlock    = "UNLOCK"
	MethodPropfind  = "PROPFIND"
	MethodProppatch = "PROPPATCH"
)

var WebDAVHandler = func(c *gin.Context, router *gin.RouterGroup, srv *webdav.Handler) {
	srv.ServeHTTP(c.Writer, c.Request)
}

func WebDAVOriginals(conf *config.Config, router *gin.RouterGroup) {
	srv := &webdav.Handler{
		Prefix:     router.BasePath(),
		FileSystem: webdav.Dir(filepath.Join(conf.StoragePath, "originals")),
		LockSystem: webdav.NewMemLS(),
	}

	handlerFunc := func(c *gin.Context) {
		WebDAVHandler(c, router, srv)
	}

	handleRead := func(h func(*gin.Context)) {
		router.Handle(MethodHead, "/*path", h)
		router.Handle(MethodGet, "/*path", h)
		router.Handle(MethodOptions, "/*path", h)
		router.Handle(MethodLock, "/*path", h)
		router.Handle(MethodUnlock, "/*path", h)
		router.Handle(MethodPropfind, "/*path", h)
	}

	handleRead(handlerFunc)
}

func WebDAVUser(conf *config.Config, router *gin.RouterGroup) {
	srv := &webdav.Handler{
		Prefix:     router.BasePath(),
		FileSystem: webdav.Dir(filepath.Join(conf.StoragePath, "user")),
		LockSystem: webdav.NewMemLS(),
	}

	handlerFunc := func(c *gin.Context) {
		WebDAVHandler(c, router, srv)
	}

	handleRead := func(h func(*gin.Context)) {
		router.Handle(MethodHead, "/*path", h)
		router.Handle(MethodGet, "/*path", h)
		router.Handle(MethodOptions, "/*path", h)
		router.Handle(MethodLock, "/*path", h)
		router.Handle(MethodUnlock, "/*path", h)
		router.Handle(MethodPropfind, "/*path", h)
	}

	handleWrite := func(h func(*gin.Context)) {
		router.Handle(MethodPut, "/*path", h)
		router.Handle(MethodPost, "/*path", h)
		router.Handle(MethodPatch, "/*path", h)
		router.Handle(MethodDelete, "/*path", h)
		router.Handle(MethodMkcol, "/*path", h)
		router.Handle(MethodCopy, "/*path", h)
		router.Handle(MethodMove, "/*path", h)
		router.Handle(MethodProppatch, "/*path", h)
	}

	handleRead(handlerFunc)

	handleWrite(handlerFunc)
}

func WebDAVAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !checkAuth(auth) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func checkAuth(auth string) bool {
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		return false
	}

	payload, _ := base64.StdEncoding.DecodeString(parts[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return false
	}

	user := entity.FindUser(pair[0])
	if user == nil {
		logger.Infof("user %s not found", pair[0])
		return false
	}

	return user.InvalidPassword(pair[1])
}

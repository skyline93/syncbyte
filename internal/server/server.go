package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"phto/internal/api"
	"phto/internal/auth"
	"phto/internal/config"
	"phto/internal/log"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	BaseUri = "/api/v1"
)

var (
	logger *logrus.Logger
	APIv1  *gin.RouterGroup
)

func init() {
	logger = log.NewLogger("server.log")
}

func registerRoutes(router *gin.Engine, conf *config.Config) {
	WebDAV(conf, router.Group("/originals", WebDAVAuth()))

	api.Login(router)
	api.Register(router)
	api.DeleteUser(router)

	api.Ping(APIv1)
	api.UploadFile(APIv1, conf)
	api.UploadPhoto(APIv1, conf)
	api.ImportPhoto(APIv1, conf)

	api.CreateAlbum(APIv1, conf)
}

func StartHttp(ctx context.Context, conf *config.Config) {
	router := gin.Default()

	APIv1 = router.Group(BaseUri, auth.AuthMiddleware())

	registerRoutes(router, conf)

	tcpSocket := fmt.Sprintf("%s:%d", conf.HttpHost, conf.HttpPort)

	listener, err := net.Listen("tcp", tcpSocket)
	if err != nil {
		logger.Errorf("server: listener %s", err)
		return
	}

	logger.Infof("server listen at %s", tcpSocket)

	server := &http.Server{Addr: tcpSocket, Handler: router}

	go func() {
		if err := server.ServeTLS(listener, conf.TLSCert, conf.TLSKey); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.Infof("server: shutdown complete")
			} else {
				logger.Infof("server: %s", err)
			}
		}
	}()

	<-ctx.Done()
	logger.Info("server: shutting down")

	if err = server.Close(); err != nil {
		logger.Errorf("server: shutdown failed (%s)", err)
	}
}

func Start(conf *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("start server")
	go StartHttp(ctx, conf)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	sig := <-quit

	logger.Info("shutting down...")
	cancel()

	time.Sleep(2 * time.Second)

	if sig == syscall.SIGUSR1 {
		os.Exit(1)
	}
}

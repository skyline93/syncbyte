package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = NewLogger("app.log")
}

func StartServer(ctx context.Context) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	tcpSocket := fmt.Sprintf("%s:%d", "0.0.0.0", 5000)

	listener, err := net.Listen("tcp", tcpSocket)
	if err != nil {
		logger.Errorf("server: listener %s", err)
		return
	}

	logger.Infof("server listen at %s", tcpSocket)

	server := &http.Server{Addr: tcpSocket, Handler: router}

	go func() {
		if err := server.Serve(listener); err != nil {
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("start server")
	go StartServer(ctx)

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

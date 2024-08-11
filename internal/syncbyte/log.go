package syncbyte

import (
	"phto/internal/log"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = log.NewLogger("server.log")
}

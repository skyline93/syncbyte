package main

import (
	"os"
	"phto/internal/commands"
	"phto/internal/log"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var logger *logrus.Logger

func init() {
	logger = log.NewLogger("app.log")
}

func main() {
	app := cli.NewApp()
	app.Name = "syncbyte"
	app.Usage = "syncbyte"
	app.Commands = commands.Syncbyte

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
	}
}

package commands

import (
	"phto/internal/config"
	"phto/internal/server"

	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts the Web server",
	Flags:   startFlags,
	Action:  startAction,
}

var startFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "host, H",
		Usage: "server host",
	},
	cli.IntFlag{
		Name:  "port, p",
		Usage: "server port",
		Value: 8000,
	},
	cli.StringFlag{
		Name:  "tls-cert",
		Usage: "tls cert file path",
		Value: "server.crt",
	},
	cli.StringFlag{
		Name:  "tls-key",
		Usage: "tls key file path",
		Value: "server.key",
	},
}

func startAction(ctx *cli.Context) error {
	conf := config.Config{
		HttpHost: ctx.String("host"),
		HttpPort: ctx.Int("port"),
		TLSCert:  ctx.String("tls-cert"),
		TLSKey:   ctx.String("tls-key"),
	}

	server.Start(&conf)

	return nil
}

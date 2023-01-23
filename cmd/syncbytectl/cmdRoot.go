package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/cmd/syncbytectl/config"
	"github.com/skyline93/syncbyte-go/internal/pkg/logging"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use:   "syncbytectl",
	Short: "syncbytectl is syncbyte command tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetLevel(logging.LogrusLevel(config.Conf.Core.LogLevel))
		log.AddHook(&logging.FormatterHook{
			Writer:    os.Stdout,
			LogLevels: log.AllLevels,
			Formatter: &log.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
				ForceColors:     true,
			},
		})
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}

func main() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

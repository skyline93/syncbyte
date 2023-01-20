package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/agent"
	"github.com/skyline93/syncbyte-go/internal/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/natefinch/lumberjack.v2"
)

var cmdRoot = &cobra.Command{
	Use:   "syncbyte-agent",
	Short: "syncbyte agent is a backup agent",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		hook := lumberjack.Logger{
			Filename:  filepath.Join(agent.Conf.Core.LogPath, "agent.log"),
			MaxSize:   1024,
			MaxAge:    365,
			Compress:  true,
			LocalTime: true,
		}

		log.SetOutput(&hook)
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(logging.LogrusLevel(agent.Conf.Core.LogLevel))

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

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run server of syncbyte-agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := agent.RunServer(); err != nil {
			fmt.Printf("run server failed, err: %v", err)
		}
	},
}

func init() {
	cobra.OnInitialize(agent.InitConfig)
	cmdRoot.AddCommand(cmdRun)
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

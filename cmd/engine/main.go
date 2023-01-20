package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/natefinch/lumberjack.v2"
)

var cmdRoot = &cobra.Command{
	Use:   "syncbyte-engine",
	Short: "syncbyte engine is a backup engine",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		hook := lumberjack.Logger{
			Filename:  filepath.Join(config.Conf.Core.LogPath, "engine.log"),
			MaxSize:   1024,
			MaxAge:    365,
			Compress:  true,
			LocalTime: true,
		}

		log.SetOutput(&hook)
		log.SetFormatter(&log.JSONFormatter{})
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

var cmdBackup = &cobra.Command{
	Use:   "backup",
	Short: "backup",
	Run: func(cmd *cobra.Command, args []string) {
		if err := backup.Backup(options.SourcePath); err != nil {
			fmt.Printf("backup failed, err: %v", err)
		}
	},
}

type Options struct {
	Host string
	Port int

	SourcePath string
}

var options Options

func init() {
	cobra.OnInitialize(config.InitConfig, repository.InitRepository)
	cmdRoot.AddCommand(cmdBackup)

	f := cmdBackup.Flags()
	f.StringVarP(&options.Host, "host", "H", "127.0.0.1", "server host")
	f.IntVarP(&options.Port, "port", "p", 50051, "server port")
	f.StringVarP(&options.SourcePath, "source", "s", "", "source path")

	cmdBackup.MarkFlagRequired("source")
	cmdBackup.MarkFlagRequired("dest")
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

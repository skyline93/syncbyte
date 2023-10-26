package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/apiserver"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/job"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/pkg/logging"
	"github.com/skyline93/syncbyte-go/internal/pkg/scheduler"
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

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run",
	Run: func(cmd *cobra.Command, args []string) {
		if runOptions.WithScheduler {
			go runScheduler()
		}

		if err := apiserver.Run(); err != nil {
			log.Errorf("apiserver run error, err : %v", err)
			os.Exit(1)
		}
	},
}

var cmdScheduler = &cobra.Command{
	Use:   "scheduler",
	Short: "scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		runScheduler()
	},
}

func runScheduler() {
	sch := scheduler.New(context.TODO(), config.Conf.Core.MaxConcurrentNum)
	scheduler.S = sch
	sch.AddPeriodicalJob(job.NewJobScheduler(sch, config.Conf.Core.JobSchedulerCron))
	sch.Start()
}

type RunOptions struct {
	WithScheduler bool
}

var runOptions RunOptions

func init() {
	cobra.OnInitialize(config.InitConfig, repository.InitRepository)

	cmdRoot.AddCommand(cmdRun)
	cmdRun.AddCommand(cmdScheduler)

	f := cmdRun.Flags()
	f.BoolVarP(&runOptions.WithScheduler, "scheduler", "S", false, "run with scheduler")
}

func main() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

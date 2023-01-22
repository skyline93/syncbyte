package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/job"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/engine/scheduling"
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

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "create",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var cmdScheduler = &cobra.Command{
	Use:   "scheduler",
	Short: "scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		sch := scheduler.New(context.TODO(), config.Conf.Core.MaxConcurrentNum)
		scheduler.S = sch
		sch.AddPeriodicalJob(job.NewJobScheduler(sch, config.Conf.Core.JobSchedulerCron))
		sch.Start()
	},
}

var cmdScheduledJob = &cobra.Command{
	Use:   "scheduledjob",
	Short: "scheduledjob",
	Run: func(cmd *cobra.Command, args []string) {
		jobid, err := scheduling.ScheduleBackup(uint(scheduledJobOptions.BackupPolicyID))
		if err != nil {
			log.Errorf("schedule backup job failed, err: %v", err)
			os.Exit(1)
		}
		log.Infof("schedule backup job successed, jobid: %d", jobid)
	},
}

var cmdBackupPolicy = &cobra.Command{
	Use:   "backuppolicy",
	Short: "backuppolicy",
	Run: func(cmd *cobra.Command, args []string) {
		func() {
			var args interface{}
			if plOptions.Resource.Type == string(backup.NAS) {
				args = backup.NasResourceArgs{
					Dir: plOptions.Resource.Dir,
				}
			}

			res := backup.Resource{
				Name: plOptions.Resource.Name,
				Type: plOptions.Resource.Type,
				Args: args,
			}
			plid, err := backup.CreatePolicy(res, plOptions.Retention)
			if err != nil {
				log.Infof("create backup policy failed, err: %v", err)
				os.Exit(1)
			}
			log.Infof("create backup policy successed, policy id is %d", plid)
		}()
	},
}

type Resource struct {
	Name string
	Type string
	Dir  string
}

type BackupPolicyOptions struct {
	Resource  Resource
	Retention int
}

type ScheduledJobOptions struct {
	BackupPolicyID int
}

var plOptions BackupPolicyOptions
var scheduledJobOptions ScheduledJobOptions

func init() {
	cobra.OnInitialize(config.InitConfig, repository.InitRepository)
	cmdRoot.AddCommand(cmdCreate)
	cmdRoot.AddCommand(cmdRun)
	cmdCreate.AddCommand(cmdBackupPolicy)
	cmdRun.AddCommand(cmdScheduledJob)
	cmdRun.AddCommand(cmdScheduler)

	fcpl := cmdBackupPolicy.Flags()
	fcpl.IntVarP(&plOptions.Retention, "retention", "r", 7, "backup set save retention")
	fcpl.StringVarP(&plOptions.Resource.Type, "type", "t", "", "resource type")
	fcpl.StringVarP(&plOptions.Resource.Name, "name", "n", "", "resource name")
	fcpl.StringVarP(&plOptions.Resource.Dir, "dir", "d", "", "backup dir, only nas resource")

	fsj := cmdScheduledJob.Flags()
	fsj.IntVarP(&scheduledJobOptions.BackupPolicyID, "backup-policy-id", "p", 0, "backup policy id")

	cmdScheduledJob.MarkFlagRequired("backup-policy-id")
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

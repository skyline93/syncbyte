package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/cmd/syncbytectl/grpc"
	"github.com/spf13/cobra"
)

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "create backuppolicy,backupjob",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var cmdBackupPolicy = &cobra.Command{
	Use:   "backuppolicy",
	Short: "backuppolicy",
	Run: func(cmd *cobra.Command, args []string) {
		var resourceArgs interface{}

		if createBackupPolicyOptions.ResourceType == "nas" {
			resourceArgs = grpc.NasResourceArgs{Dir: createBackupPolicyOptions.BackupPath}
		}

		plID, err := grpc.CreateBackupPolicy(
			createBackupPolicyOptions.ResourceName,
			createBackupPolicyOptions.ResourceType,
			resourceArgs,
			createBackupPolicyOptions.Retention,
		)
		if err != nil {
			log.Errorf("create backup policy error, err: %v", err)
			os.Exit(1)
		}

		log.Infof("create backup policy successed, policy_id: %d", plID)
	},
}

var cmdBackupJob = &cobra.Command{
	Use:   "backupjob",
	Short: "backupjob",
	Run: func(cmd *cobra.Command, args []string) {
		jobID, err := grpc.StartBackup(startBackupOptions.ResourceID)
		if err != nil {
			log.Errorf("start backup job error, err: %v", err)
			os.Exit(1)
		}
		log.Infof("start backup job successed, job_id: %d", jobID)
	},
}

type CreateBackupPolicyOptions struct {
	ResourceName string
	ResourceType string
	BackupPath   string
	Retention    int
}

type StartBackupOptions struct {
	ResourceID uint
}

var createBackupPolicyOptions CreateBackupPolicyOptions
var startBackupOptions StartBackupOptions

func init() {
	cmdRoot.AddCommand(cmdCreate)
	cmdCreate.AddCommand(cmdBackupPolicy)
	cmdCreate.AddCommand(cmdBackupJob)

	bp := cmdBackupPolicy.Flags()
	bp.IntVarP(&createBackupPolicyOptions.Retention, "retention", "r", 7, "backup set save retention")
	bp.StringVarP(&createBackupPolicyOptions.ResourceType, "type", "t", "", "resource type, example: 'nas'")
	bp.StringVarP(&createBackupPolicyOptions.ResourceName, "name", "n", "", "resource name")
	bp.StringVarP(&createBackupPolicyOptions.BackupPath, "backup-path", "p", "", "backup path, only nas resource")

	cmdBackupPolicy.MarkFlagRequired("retention")
	cmdBackupPolicy.MarkFlagRequired("type")
	cmdBackupPolicy.MarkFlagRequired("name")

	bs := cmdBackupJob.Flags()
	bs.UintVarP(&startBackupOptions.ResourceID, "resource-id", "i", 0, "resource id")

	cmdBackupJob.MarkFlagRequired("resource-id")
}

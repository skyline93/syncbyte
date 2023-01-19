package main

import (
	"fmt"
	"os"

	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use:   "syncbyte-engine",
	Short: "syncbyte engine is a backup engine",
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

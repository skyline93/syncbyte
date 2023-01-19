package main

import (
	"fmt"
	"os"

	"github.com/skyline93/syncbyte-go/internal/agent"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use:   "syncbyte-agent",
	Short: "syncbyte agent is a backup agent",
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

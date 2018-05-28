package cmd

import (
	"github.com/spf13/cobra"
	"github.com/procroner/ssh-run/core/server"
	"github.com/procroner/ssh-run/core/job"
	"github.com/procroner/ssh-run/core/db"
)

func Run() {
	var cmdServer = &cobra.Command{
		Use:   "server [add|delete|edit|list]",
		Short: "Add, delete, edit or list server",
		Long:  "Command server is for server management, providing add, delete, edit or list commands",
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdServerList = &cobra.Command{
		Use:   "list",
		Short: "List servers",
		Long:  "List all servers",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			server.List()
		},
	}

	var cmdJob = &cobra.Command{
		Use:   "job [add|delete|edit|list]",
		Short: "Add, delete, edit or list job",
		Long:  "Command job is for job management, providing add, delete, edit or list commands",
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdJobList = &cobra.Command{
		Use:   "list",
		Short: "List jobs",
		Long:  "List all jobs",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			job.List()
		},
	}

	var cmdJobRun = &cobra.Command{
		Use:   "run",
		Short: "Run all jobs",
		Long:  "Run all jobs",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			job.RunAll()
		},
	}

	var cmdTable = &cobra.Command{
		Use:   "table",
		Short: "Table management",
		Long:  "Init tables",
		Args:  cobra.MinimumNArgs(1),
	}

	var tableName string
	var cmdTableInit = &cobra.Command{
		Use:   "init",
		Short: "Init tables",
		Long:  "Init tables including jobs, logs and servers",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			db.InitTables(tableName)
		},
	}
	cmdTableInit.Flags().StringVarP(&tableName, "table", "t", "", "tables names to init, multiple tables separated by comma")

	var rootCmd = &cobra.Command{Use: "ssh-run"}
	rootCmd.AddCommand(cmdServer, cmdJob, cmdTable)
	cmdServer.AddCommand(cmdServerList)
	cmdJob.AddCommand(cmdJobList, cmdJobRun)
	cmdTable.AddCommand(cmdTableInit)
	rootCmd.Execute()
}

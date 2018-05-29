package cmd

import (
	"github.com/spf13/cobra"
	"github.com/procroner/ssh-run/cmd/cmdServer"
	"github.com/procroner/ssh-run/cmd/cmdJob"
	"github.com/procroner/ssh-run/cmd/cmdTable"
)

func Run() {
	//server command
	var cServer = &cobra.Command{
		Use:   "server [add|delete|edit|list|ping]",
		Short: "Add, delete, edit, list or ping server",
		Long:  "Command server is for server management, providing add, delete, edit or list commands",
		Args:  cobra.MinimumNArgs(1),
	}

	var cServerList = &cobra.Command{
		Use:   "list",
		Short: "List servers",
		Long:  "List all servers",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cmdServer.List()
		},
	}

	var cServerPing = &cobra.Command{
		Use:   "ping",
		Short: "Ping servers",
		Long:  "Ping all servers",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cmdServer.Ping()
		},
	}
	cServer.AddCommand(cServerList, cServerPing)

	//job command
	var cJob = &cobra.Command{
		Use:   "job [add|delete|edit|list]",
		Short: "Add, delete, edit or list cmdJob",
		Long:  "Command job is for job management, providing add, delete, edit or list commands",
		Args:  cobra.MinimumNArgs(1),
	}

	var cJobList = &cobra.Command{
		Use:   "list",
		Short: "List jobs",
		Long:  "List all jobs",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cmdJob.List()
		},
	}

	var jobId int
	var cJobRun = &cobra.Command{
		Use:   "run",
		Short: "Run all jobs",
		Long:  "Run all jobs",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cmdJob.Run(jobId)
		},
	}
	cJobRun.Flags().IntVarP(&jobId, "job", "j", 0, "jobs that willing be run, separated by comma")
	cJob.AddCommand(cJobList, cJobRun)

	//table command
	var cTable = &cobra.Command{
		Use:   "table",
		Short: "Table management",
		Long:  "Init tables",
		Args:  cobra.MinimumNArgs(1),
	}

	var tableName string
	var cTableMigrate = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate tables",
		Long:  "Migrate tables including jobs, logs and servers",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cmdTable.MigrateTables(tableName)
		},
	}
	cTableMigrate.Flags().StringVarP(&tableName, "table", "t", "", "tables names to init, multiple tables separated by comma")

	cTable.AddCommand(cTableMigrate)

	var rootCmd = &cobra.Command{Use: "ssh-run"}
	rootCmd.AddCommand(cServer, cJob, cTable)
	rootCmd.Execute()
}

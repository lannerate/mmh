package cmd

import (
	"github.com/mritd/mmh/mmh"
	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:     "go SERVER_NAME",
	Aliases: []string{"mgo"},
	Short:   "Login single server",
	Long: `
Login single server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			mmh.SingleLogin(args[0])
		} else {
			_ = cmd.Help()
		}
	},
	PreRun:  mmh.UpdateContextTimestampTask,
	PostRun: mmh.UpdateContextTimestamp,
}

func init() {
	RootCmd.AddCommand(goCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sshm",
	Short: "Gerenciador de conexões SSH",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd, listCmd, connectCmd, tagCmd, groupCmd)
}
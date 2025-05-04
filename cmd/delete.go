// cmd/delete.go
package cmd

import (
	"fmt"
	"github.com/andreixhz/sshm/internal"
	"github.com/andreixhz/sshm/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [alias]",
	Short: "Remove uma conexão SSH pelo alias",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		alias := args[0]
		hosts, _ := internal.LoadHosts()
		newHosts := []models.Host{}
		found := false
		for _, h := range hosts {
			if h.Alias != alias {
				newHosts = append(newHosts, h)
			} else {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("alias não encontrado")
		}
		return internal.SaveHosts(newHosts)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

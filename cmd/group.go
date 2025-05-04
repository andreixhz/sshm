package cmd

import (
	"fmt"
	"github.com/andreixhz/sshm/internal"
	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Gerencia grupos",
}

var groupRenameCmd = &cobra.Command{
	Use:   "rename [alias] [novoGrupo]",
	Short: "Renomeia o grupo de uma conexão",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, _ := internal.LoadHosts()
		for i, h := range hosts {
			if h.Alias == args[0] {
				h.Group = args[1]
				hosts[i] = h
				return internal.SaveHosts(hosts)
			}
		}
		return fmt.Errorf("alias não encontrado")
	},
}

func init() {
	groupCmd.AddCommand(groupRenameCmd)
}
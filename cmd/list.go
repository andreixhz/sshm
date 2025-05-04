package cmd

import (
	"fmt"
	"github.com/andreixhz/sshm/internal"
	"github.com/spf13/cobra"
)

var tagFilter string
var groupFilter string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista conexões SSH",
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, _ := internal.LoadHosts()
		for _, h := range hosts {
			if (tagFilter == "" || contains(h.Tags, tagFilter)) && (groupFilter == "" || h.Group == groupFilter) {
				fmt.Printf("%s -> %s@%s:%d [tags: %v] [group: %s]\n", h.Alias, h.User, h.Host, h.Port, h.Tags, h.Group)
			}
		}
		return nil
	},
}

func init() {
	listCmd.Flags().StringVar(&tagFilter, "tag", "", "Filtrar por tag")
	listCmd.Flags().StringVar(&groupFilter, "group", "", "Filtrar por grupo")
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

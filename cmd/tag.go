package cmd

import (
	"fmt"
	"github.com/andreixhz/sshm/internal"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Gerencia tags",
}

var tagAddCmd = &cobra.Command{
	Use:   "add [alias] [tag]",
	Short: "Adiciona uma tag a uma conexão",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, _ := internal.LoadHosts()
		for i, h := range hosts {
			if h.Alias == args[0] {
				h.Tags = append(h.Tags, args[1])
				hosts[i] = h
				return internal.SaveHosts(hosts)
			}
		}
		return fmt.Errorf("alias não encontrado")
	},
}

var tagRemoveCmd = &cobra.Command{
	Use:   "remove [alias] [tag]",
	Short: "Remove uma tag de uma conexão",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, _ := internal.LoadHosts()
		for i, h := range hosts {
			if h.Alias == args[0] {
				newTags := []string{}
				for _, t := range h.Tags {
					if t != args[1] {
						newTags = append(newTags, t)
					}
				}
				h.Tags = newTags
				hosts[i] = h
				return internal.SaveHosts(hosts)
			}
		}
		return fmt.Errorf("alias não encontrado")
	},
}

func init() {
	tagCmd.AddCommand(tagAddCmd, tagRemoveCmd)
}
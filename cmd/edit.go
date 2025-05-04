package cmd

import (
	"fmt"
	"github.com/andreixhz/sshm/internal"
	"github.com/spf13/cobra"
	
)

var (
	edAlias  string
	edHost   string
	edUser   string
	edPort   int
	edTags   []string
	edGroup  string
	edPasswd string
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edita uma conexão existente",
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, _ := internal.LoadHosts()
		for i, h := range hosts {
			if h.Alias == edAlias {
				if edHost != "" {
					h.Host = edHost
				}
				if edUser != "" {
					h.User = edUser
				}
				if cmd.Flags().Changed("port") {
					h.Port = edPort
				}
				if len(edTags) > 0 {
					h.Tags = edTags
				}
				if edGroup != "" {
					h.Group = edGroup
				}
				if edPasswd != "" {
					h.Password = edPasswd
				}
				hosts[i] = h
				return internal.SaveHosts(hosts)
			}
		}
		return fmt.Errorf("alias não encontrado")
	},
}

func init() {
	editCmd.Flags().StringVar(&edAlias, "alias", "", "Alias da conexão a ser editada")
	editCmd.Flags().StringVar(&edHost, "host", "", "Novo host")
	editCmd.Flags().StringVar(&edUser, "user", "", "Novo usuário")
	editCmd.Flags().IntVar(&edPort, "port", 0, "Nova porta")
	editCmd.Flags().StringSliceVar(&edTags, "tags", []string{}, "Novas tags")
	editCmd.Flags().StringVar(&edGroup, "group", "", "Novo grupo")
	editCmd.Flags().StringVar(&edPasswd, "password", "", "Nova senha")
	editCmd.MarkFlagRequired("alias")
	rootCmd.AddCommand(editCmd)
}
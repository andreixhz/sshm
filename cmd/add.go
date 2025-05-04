package cmd

import (
	"github.com/andreixhz/sshm/internal"
	"github.com/andreixhz/sshm/models"
	"github.com/spf13/cobra"
)

var (
	alias string
	host string
	user string
	port int
	tags []string
	group string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova conexão SSH",
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, _ := internal.LoadHosts()
		hosts = append(hosts, models.Host{
			Alias: alias,
			Host:  host,
			User:  user,
			Port:  port,
			Tags:  tags,
			Group: group,
		})
		return internal.SaveHosts(hosts)
	},
}

func init() {
	addCmd.Flags().StringVar(&alias, "alias", "", "Apelido da conexão")
	addCmd.Flags().StringVar(&host, "host", "", "Endereço IP ou hostname")
	addCmd.Flags().StringVar(&user, "user", "", "Usuário SSH")
	addCmd.Flags().IntVar(&port, "port", 22, "Porta SSH")
	addCmd.Flags().StringSliceVar(&tags, "tags", []string{}, "Lista de tags")
	addCmd.Flags().StringVar(&group, "group", "", "Grupo")
	addCmd.MarkFlagRequired("alias")
	addCmd.MarkFlagRequired("host")
	addCmd.MarkFlagRequired("user")
}
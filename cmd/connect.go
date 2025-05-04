package cmd

import (
	"fmt"
	"os/exec"
	"github.com/andreixhz/sshm/internal"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [alias]",
	Short: "Conecta-se via SSH usando o alias",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		alias := args[0]
		hosts, _ := internal.LoadHosts()
		for _, h := range hosts {
			if h.Alias == alias {
				// Se não tiver senha, usa o método tradicional do comando ssh
				if h.Password == "" {
					sshCmd := exec.Command("ssh", fmt.Sprintf("%s@%s", h.User, h.Host), "-p", fmt.Sprintf("%d", h.Port))
					sshCmd.Stdin = cmd.InOrStdin()
					sshCmd.Stdout = cmd.OutOrStdout()
					sshCmd.Stderr = cmd.ErrOrStderr()
					return sshCmd.Run()
				} else {
					// Se tiver senha, usa o comando sshpass para fornecer a senha automaticamente
					sshpassCmd := exec.Command("sshpass", "-p", h.Password, "ssh", 
						fmt.Sprintf("%s@%s", h.User, h.Host), 
						"-p", fmt.Sprintf("%d", h.Port),
						"-o", "StrictHostKeyChecking=no")
					sshpassCmd.Stdin = cmd.InOrStdin()
					sshpassCmd.Stdout = cmd.OutOrStdout()
					sshpassCmd.Stderr = cmd.ErrOrStderr()
					
					err := sshpassCmd.Run()
					if err != nil {
						// Se falhar com sshpass, informa ao usuário que pode precisar instalar o sshpass
						fmt.Fprintf(cmd.ErrOrStderr(), "Erro ao conectar com senha: %v\n", err)
						fmt.Fprintf(cmd.ErrOrStderr(), "Certifique-se que o sshpass está instalado no seu sistema.\n")
						
						// Tenta o método tradicional como fallback
						sshCmd := exec.Command("ssh", fmt.Sprintf("%s@%s", h.User, h.Host), "-p", fmt.Sprintf("%d", h.Port))
						sshCmd.Stdin = cmd.InOrStdin()
						sshCmd.Stdout = cmd.OutOrStdout()
						sshCmd.Stderr = cmd.ErrOrStderr()
						return sshCmd.Run()
					}
					return nil
				}
			}
		}
		return fmt.Errorf("alias não encontrado: %s", alias)
	},
}

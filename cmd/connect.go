package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/andreixhz/sshm/internal"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
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
				// Verifica se está sendo executado em um terminal interativo
				isTerm := false
				fdIn := int(os.Stdin.Fd())
				if term.IsTerminal(fdIn) {
					isTerm = true
				}

				// Se não tiver senha ou não estiver em um terminal interativo,
				// usa o método tradicional do comando ssh externo
				if h.Password == "" || !isTerm {
					sshCmd := exec.Command("ssh", fmt.Sprintf("%s@%s", h.User, h.Host), "-p", fmt.Sprintf("%d", h.Port))
					sshCmd.Stdin = cmd.InOrStdin()
					sshCmd.Stdout = cmd.OutOrStdout()
					sshCmd.Stderr = cmd.ErrOrStderr()
					return sshCmd.Run()
				}

				// Configuração da conexão SSH usando a biblioteca golang.org/x/crypto/ssh
				config := &ssh.ClientConfig{
					User: h.User,
					Auth: []ssh.AuthMethod{
						ssh.Password(h.Password),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
					Timeout:         10 * time.Second,
				}

				// Conectar ao servidor SSH
				addr := fmt.Sprintf("%s:%d", h.Host, h.Port)
				client, err := ssh.Dial("tcp", addr, config)
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "Erro ao conectar via SSH: %v\n", err)
					fmt.Fprintf(cmd.ErrOrStderr(), "Tentando método alternativo...\n")

					// Fallback para o método tradicional
					sshCmd := exec.Command("ssh", fmt.Sprintf("%s@%s", h.User, h.Host), "-p", fmt.Sprintf("%d", h.Port))
					sshCmd.Stdin = cmd.InOrStdin()
					sshCmd.Stdout = cmd.OutOrStdout()
					sshCmd.Stderr = cmd.ErrOrStderr()
					return sshCmd.Run()
				}
				defer client.Close()

				// Criar uma sessão
				session, err := client.NewSession()
				if err != nil {
					return fmt.Errorf("falha ao criar sessão SSH: %v", err)
				}
				defer session.Close()

				// Configurar terminal
				modes := ssh.TerminalModes{
					ssh.ECHO:          1,
					ssh.TTY_OP_ISPEED: 14400,
					ssh.TTY_OP_OSPEED: 14400,
				}

				termWidth, termHeight, err := term.GetSize(fdIn)
				if err != nil {
					termWidth, termHeight = 80, 24 // Valores padrão
				}

				if err := session.RequestPty("xterm", termHeight, termWidth, modes); err != nil {
					return fmt.Errorf("falha ao requisitar pty: %v", err)
				}

				// Configurar I/O
				session.Stdout = cmd.OutOrStdout()
				session.Stderr = cmd.ErrOrStderr()
				session.Stdin = cmd.InOrStdin()

				// Inicia shell
				if err := session.Shell(); err != nil {
					return fmt.Errorf("falha ao iniciar shell: %v", err)
				}

				// Aguarda até que a sessão termine
				return session.Wait()
			}
		}
		return fmt.Errorf("alias não encontrado: %s", alias)
	},
}

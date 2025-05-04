package cmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "set-password",
	Short: "Define uma senha global para conexões SSH",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Digite a senha global: ")
		var passwd string
		fmt.Scanln(&passwd)
		if strings.TrimSpace(passwd) == "" {
			return fmt.Errorf("senha inválida")
		}
		return os.WriteFile(".global_pass", []byte(passwd), 0600)
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
}

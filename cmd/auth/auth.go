package auth

import (
	"github.com/spf13/cobra"
	"github.com/timwehrle/act/cmd/auth/login"
)

var AuthCmd = &cobra.Command{
	Use:   "auth <command>",
	Short: "Authenticate ACT with Asana",
}

func init() {
	AuthCmd.AddCommand(login.LoginCmd)
}

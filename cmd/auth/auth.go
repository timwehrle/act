package auth

import (
	"github.com/spf13/cobra"
	"github.com/timwehrle/act/cmd/auth/login"
	"github.com/timwehrle/act/cmd/auth/logout"
	"github.com/timwehrle/act/cmd/auth/status"
)

var AuthCmd = &cobra.Command{
	Use:   "auth <command>",
	Short: "Authenticate ACT with Asana",
}

func init() {
	AuthCmd.AddCommand(login.LoginCmd)
	AuthCmd.AddCommand(logout.LogoutCmd)
	AuthCmd.AddCommand(status.StatusCmd)
}

package cmd

import (
	"github.com/snapp-incubator/timeoff/internal/app/timeoff/cmd/server"
	"github.com/spf13/cobra"
)

// NewRootCommand creates a new timeoff root command.
func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use: "timeoff",
	}

	server.Register(root)

	return root
}

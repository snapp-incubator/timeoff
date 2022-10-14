package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zeinababbasi/timeoff/internal/app/timeoff/cmd/server"
)

// NewRootCommand creates a new timeoff root command.
func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use: "timeoff",
	}

	server.Register(root)

	return root
}

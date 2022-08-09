package cmd

import (
	"github.com/spf13/cobra"
)

var Verbose bool

var baseCmd = &cobra.Command{
	Use:   "secret-watcher",
	Short: "Kubernetes controller to watch secrets.",
}

// GetRootCommand returns the root cobra command to be executed
func GetRootCommand() *cobra.Command {
	baseCmd.AddCommand(getServerCommand())
	return baseCmd
}

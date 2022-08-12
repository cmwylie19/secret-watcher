package cmd

import (
	"sync"

	"github.com/cmwylie19/secret-watcher/internal/server"
	"github.com/spf13/cobra"
)

var (

	// WaitGroup is used to wait for the program to finish goroutines.
	wg sync.WaitGroup
	// cfgPath is the path to the EnvoyGateway configuration file.
	cfgPath   string
	tlsKey    string
	tlsCert   string
	port      string
	label     string
	frequency string
)

// getServerCommand returns the server cobra command to be executed.
func getServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"serve"},
		Short:   "Serve Media Controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := &server.Server{}
			s.Init()
			return s.Serve(tlsKey, tlsCert, port, label)
		},
	}

	cmd.PersistentFlags().StringVarP(&label, "label", "l", "", "Search for secrets containing label")

	cmd.PersistentFlags().StringVarP(&tlsKey, "key", "", "",
		"Server private key for TLS. If not provided, TLS will not be used.")

	cmd.PersistentFlags().StringVarP(&tlsCert, "cert", "", "",
		"Server certificate for TLS. If not provided, TLS will not be used.")

	cmd.PersistentFlags().StringVarP(&port, "port", "p", "8080",
		"Port to expose the application.")

	return cmd
}

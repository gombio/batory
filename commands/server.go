package commands

import (
	"github.com/gombio/batory/server"
	"github.com/spf13/cobra"
)

var (
	address string
	port    string
)
var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "run in server mode",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewHttpServer(address, port)
		s.Execute()
	},
}

func init() {
	ServerCommand.Flags().StringVar(&address, "address", "0.0.0.0", "Bind to address")
	ServerCommand.Flags().StringVar(&port, "port", "8081", "Bind to port")
}

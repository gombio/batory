package commands

import (
	"os"
	"github.com/gombio/batory/fixtures"
	"github.com/spf13/cobra"
)

var DevDataCommand = &cobra.Command{
	Use:   "loadData",
	Short: "load dev data",
	Run: func(cmd *cobra.Command, args []string) {
		fixtures.LoadFixtures()
			os.Exit(0)
	},
}

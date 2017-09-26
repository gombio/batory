package commands

import (
	"fmt"
	"github.com/gombio/batory/config"
	"github.com/spf13/cobra"
)

var (
	globalFlags struct {
		DryRunMode bool
		ConfigDir  string
	}
)

var RootCmd = &cobra.Command{
	Use:   "batory",
	Short: "Service centre managment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run batory help")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&globalFlags.DryRunMode, "dry-run", "", false, "Run in dry-run mode")
	RootCmd.PersistentFlags().StringVar(&globalFlags.ConfigDir, "config-dir", "", "Change the default configuration directory")
	RootCmd.PersistentFlags().StringVar(&config.AppConfig.Database.Provider, "data-provider", "rethinkDb", "Choice data storage")
	RootCmd.PersistentFlags().StringVar(&config.AppConfig.Database.Address, "db-address", "127.0.0.1", "db address")
	RootCmd.PersistentFlags().StringVar(&config.AppConfig.Database.DbName, "db-name", "batory", "db name")
}

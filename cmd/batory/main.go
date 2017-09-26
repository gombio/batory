package main

import (
	"fmt"
	cmd "github.com/gombio/batory/commands"
	"os"
)

func main() {
	cmd.RootCmd.AddCommand(cmd.ServerCommand)
	cmd.RootCmd.AddCommand(cmd.DevDataCommand)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

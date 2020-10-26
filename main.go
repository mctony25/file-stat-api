package main

import (
	"file-stat/command"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	rootCmd := &cobra.Command{}

	cliCmd := (command.CliCommand{}).GetCommand()
	apiCmd := (command.ApiCommand{}).GetCommand()

	rootCmd.AddCommand(cliCmd, apiCmd)
	err := rootCmd.Execute()
	if nil != err {
		code, _ := fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}
}
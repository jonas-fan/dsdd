package main

import (
	"fmt"
	"os"

	"github.com/jonas-fan/dsdd/internal/cmd"
	"github.com/spf13/cobra"
)

var version string

func main() {
	name := fmt.Sprintf("Deep Security Diagnostic Debugger %s", version)

	command := &cobra.Command{
		Use:   "dsdd",
		Short: name,
		Long:  name,
	}

	for _, each := range cmd.NewCommands() {
		command.AddCommand(each)
	}

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
}

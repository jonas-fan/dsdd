package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	command := &cobra.Command{
		Use: "dsdd",
		Short: "Deep Security Diagnostic Debugger",
		Long: "Deep Security Diagnostic Debugger",
	}

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
}

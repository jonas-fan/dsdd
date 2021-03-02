package cmd

import (
	"github.com/jonas-fan/dsdd/internal/cmd/event"
	"github.com/jonas-fan/dsdd/internal/cmd/process"
	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	var commands = []*cobra.Command{
		event.NewCommand(),
		process.NewCommand(),
	}

	return commands
}

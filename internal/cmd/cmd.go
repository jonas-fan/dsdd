package cmd

import (
	"github.com/jonas-fan/dsdd/internal/cmd/event"
	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	var commands = []*cobra.Command{
		event.NewCommand(),
	}

	return commands
}

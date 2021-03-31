package cmd

import (
	"github.com/jonas-fan/dsdd/internal/cmd/events"
	"github.com/jonas-fan/dsdd/internal/cmd/logs"
	"github.com/jonas-fan/dsdd/internal/cmd/ps"
	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	var commands = []*cobra.Command{
		events.NewCommand(),
		logs.NewCommand(),
		ps.NewCommand(),
	}

	return commands
}

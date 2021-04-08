package cmd

import (
	"github.com/jonas-fan/dsdd/internal/cmd/events"
	"github.com/jonas-fan/dsdd/internal/cmd/logs"
	"github.com/jonas-fan/dsdd/internal/cmd/ps"
	"github.com/jonas-fan/dsdd/internal/cmd/system"
	"github.com/spf13/cobra"
)

func NewCommands() []*cobra.Command {
	commands := []*cobra.Command{
		events.NewCommand(),
		logs.NewCommand(),
		ps.NewCommand(),
		system.NewCommand(),
	}

	return commands
}

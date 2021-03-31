package logs

import (
	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	diagnostic.ReadLog()
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "logs",
		Short: "Show logs",
		Run:   run,
	}

	return command
}

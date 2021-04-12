package agent

import (
	"fmt"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	agent, err := diagnostic.NewAgent("Agent")

	if err != nil {
		panic(err)
	}

	fmt.Println(agent)
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "agent",
		Short: "Display host-specific data",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(true)

	return command
}

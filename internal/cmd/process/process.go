package process

import (
	"fmt"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	data, err := diagnostic.ReadProcess()

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "process",
		Short: "Display running process",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(false)

	return command
}

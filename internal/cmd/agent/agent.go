package agent

import (
	"fmt"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/agent"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	dsa, err := agent.NewAgent("Agent")

	if err != nil {
		panic(err)
	}

	fmt.Println("dsa.os =", dsa.System.Type)
	fmt.Println("dsa.os.name =", dsa.System.Name)
	fmt.Println("dsa.os.release =", dsa.System.Release)
	fmt.Println("dsa.timestamp =", dsa.Timestamp)
	fmt.Println("dsa.version =", dsa.Version)
	fmt.Println("dsa.guid =", dsa.Guid.Agent)
	fmt.Println("dsa.guid.manager =", dsa.Guid.Manager)

	for i, module := range dsa.Module {
		prefix := fmt.Sprintf("dsa.module%d.%s.%s", i, module.Type, module.Name)

		fmt.Printf("%s = %s\n", prefix, module.Version)

		for i, dep := range module.Dependency {
			fmt.Printf("%s.dependency%d = %s\n", prefix, i, dep.Name)
		}
	}
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

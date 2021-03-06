package events

import (
	"strings"

	"github.com/spf13/cobra"
)

var category string

func run(cmd *cobra.Command, args []string) {
	switch strings.ToLower(category) {
	case "sys", "system":
		readSystemEvent()
	default:
		panic("Unknown category: " + category)
	}
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "event",
		Short: "Display various events",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(false)
	flags.StringVarP(&category, "category", "c", "system", "Event category")

	return command
}

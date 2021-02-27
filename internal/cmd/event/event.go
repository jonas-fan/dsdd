package event

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var category string

var timeLayouts = [...]string{
	"January 2, 2006 15:04:05",
	"January 2, 2006 15:04:05 PM",
}

func toTime(value string) time.Time {
	for _, layout := range timeLayouts {
		if out, err := time.Parse(layout, value); err == nil {
			return out
		}
	}

	return time.Time{}
}

func run(cmd *cobra.Command, args []string) {
	switch strings.ToLower(category) {
	case "sys", "system":
		readSystemEvent(filepath.Join("Manager", "hostevents.csv"))
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

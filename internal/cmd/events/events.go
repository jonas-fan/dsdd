package events

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event/antimalware"
	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event/system"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/spf13/cobra"
)

var category string
var oneline bool

func writeEvents(writer io.Writer, events []event.Event) {
	for i := len(events) - 1; i >= 0; i-- {
		fmt.Fprintln(writer, events[i].String())
	}
}

func writeOnelineEvents(writer io.Writer, header []string, events []event.Event) {
	formatter := fmtutil.NewFormatter()

	formatter.Write(header...)

	for i := len(events) - 1; i >= 0; i-- {
		tokens := events[i].Column()

		formatter.Write(tokens...)
	}

	fmt.Fprintln(writer, formatter.String())
}

func run(cmd *cobra.Command, args []string) {
	var header []string
	var events []event.Event
	var err error

	switch strings.ToLower(category) {
	case "sys", "system":
		header = system.Header
		events, err = system.ReadSystemEventFrom(filepath.Join("Manager", "hostevents.csv"))
	case "am", "antimalware":
		header = antimalware.Header
		events, err = antimalware.ReadAntiMalwareEventFrom(filepath.Join("Manager", "antimalwareevents.csv"))
	default:
		panic("unknown category: " + category)
	}

	if err != nil {
		panic(err)
	}

	if oneline {
		writeOnelineEvents(os.Stdout, header, events)
	} else {
		writeEvents(os.Stdout, events)
	}
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "events",
		Short: "Display various events",
		Run:   run,
	}

	flags := command.Flags()
	flags.SetInterspersed(false)
	flags.StringVarP(&category, "category", "c", "system", "Event category")
	flags.BoolVarP(&oneline, "oneline", "", false, "Show information on the same line")

	return command
}

package events

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event/antimalware"
	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event/system"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/spf13/cobra"
)

var kind string
var oneline bool

func newTableViewer(kind string, events []event.Event) (*event.TableViewer, error) {
	var layout event.TableLayout

	switch strings.ToLower(kind) {
	case "sys", "system":
		layout = system.NewTableLayout()
	case "am", "antimalware":
		layout = antimalware.NewTableLayout()
	default:
		return nil, errors.New("unknown type: " + kind)
	}

	return event.NewTableViewer(layout, events), nil
}

func newReader(kind string) (*event.Reader, error) {
	var filename string
	var parser event.Parser

	switch strings.ToLower(kind) {
	case "sys", "system":
		filename, parser = filepath.Join("Manager", "hostevents.csv"), system.Parse
	case "am", "antimalware":
		filename, parser = filepath.Join("Manager", "antimalwareevents.csv"), antimalware.Parse
	default:
		return nil, errors.New("unknown type: " + kind)
	}

	return event.Open(filename, parser)
}

func run(cmd *cobra.Command, args []string) {
	reader, err := newReader(kind)

	if err != nil {
		panic(err)
	}

	defer reader.Close()

	events := make([]event.Event, 0)

	for {
		event, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		events = append(events, event)
	}

	if oneline {
		viewer, err := newTableViewer(kind, events)

		if err != nil {
			panic(err)
		}

		columns := viewer.Header()
		formatter := fmtutil.NewFormatter(columns...)

		for viewer.HasNext() {
			columns := viewer.Next()

			formatter.Write(columns...)
		}

		fmt.Println(formatter.String())
	} else {
		for _, each := range events {
			fmt.Println(each.String())
		}
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
	flags.StringVarP(&kind, "kind", "k", "system", "Event type")
	flags.BoolVarP(&oneline, "oneline", "", false, "Show information on the same line")

	return command
}

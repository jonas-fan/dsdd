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

func newViewer(kind string, events []event.Event) (event.TableViewer, error) {
	switch strings.ToLower(kind) {
	case "sys", "system":
		return system.NewTableViewer(events), nil
	case "am", "antimalware":
		return antimalware.NewTableViewer(events), nil
	default:
		return nil, errors.New("unknown type: " + kind)
	}
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
		viewer, err := newViewer(kind, events)

		if err != nil {
			panic(err)
		}

		columns := viewer.Header()

		formatter := fmtutil.NewFormatter()
		formatter.Write(columns...)

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

package events

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/antimalware"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/applicationcontrol"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/integritymonitoring"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/system"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/spf13/cobra"
)

var kind string
var details bool

func contains(slice []string, value string) bool {
	for _, each := range slice {
		if each == value {
			return true
		}
	}

	return false
}

func newTableViewer(kind string, events []event.Event) (*event.TableViewer, error) {
	var layout event.TableLayout
	var name = strings.ToLower(kind)

	switch {
	case contains(system.Alias(), name):
		layout = system.NewTableLayout()
	case contains(antimalware.Alias(), name):
		layout = antimalware.NewTableLayout()
	case contains(applicationcontrol.Alias(), name):
		layout = applicationcontrol.NewTableLayout()
	case contains(integritymonitoring.Alias(), name):
		layout = integritymonitoring.NewTableLayout()
	default:
		return nil, errors.New("unknown type: " + kind)
	}

	return event.NewTableViewer(layout, events), nil
}

func newReader(kind string) (*event.Reader, error) {
	var filename string
	var builder event.EventBuilder
	var name = strings.ToLower(kind)

	switch {
	case contains(system.Alias(), name):
		filename, builder = filepath.Join("Manager", "hostevents.csv"), system.New
	case contains(antimalware.Alias(), name):
		filename, builder = filepath.Join("Manager", "antimalwareevents.csv"), antimalware.New
	case contains(applicationcontrol.Alias(), name):
		filename, builder = filepath.Join("Manager", "appcontrolevents.csv"), applicationcontrol.New
	case contains(integritymonitoring.Alias(), name):
		filename, builder = filepath.Join("Manager", "integrityevents.csv"), integritymonitoring.New
	default:
		return nil, errors.New("unknown type: " + kind)
	}

	return event.Open(filename, builder)
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

	sort.SliceStable(events, func(lhs int, rhs int) bool {
		return events[lhs].Datetime() > events[rhs].Datetime()
	})

	if details {
		for _, each := range events {
			fmt.Println(each.String())
		}
	} else {
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
	flags.BoolVarP(&details, "details", "d", false, "Show details")

	return command
}

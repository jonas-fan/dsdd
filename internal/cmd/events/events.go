package events

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/ac"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/am"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/fw"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/im"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/ips"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/li"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/system"
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event/wrs"
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

func checkKind(kind string) (string, event.EventBuilder, event.TableLayout, error) {
	name := strings.ToLower(kind)

	switch {
	case contains(ac.Alias(), name):
		return filepath.Join("Manager", "appcontrolevents.csv"), ac.New, ac.NewTableLayout(), nil
	case contains(am.Alias(), name):
		return filepath.Join("Manager", "antimalwareevents.csv"), am.New, am.NewTableLayout(), nil
	case contains(fw.Alias(), name):
		return filepath.Join("Manager", "firewallevents.csv"), fw.New, fw.NewTableLayout(), nil
	case contains(im.Alias(), name):
		return filepath.Join("Manager", "integrityevents.csv"), im.New, im.NewTableLayout(), nil
	case contains(ips.Alias(), name):
		return filepath.Join("Manager", "dpievents.csv"), ips.New, ips.NewTableLayout(), nil
	case contains(li.Alias(), name):
		return filepath.Join("Manager", "loginspectionevents.csv"), li.New, li.NewTableLayout(), nil
	case contains(system.Alias(), name):
		return filepath.Join("Manager", "hostevents.csv"), system.New, system.NewTableLayout(), nil
	case contains(wrs.Alias(), name):
		return filepath.Join("Manager", "webreputationevents.csv"), wrs.New, wrs.NewTableLayout(), nil
	}

	return "", nil, nil, errors.New("unknown type: " + kind)
}

func run(cmd *cobra.Command, args []string) {
	filename, builder, layout, err := checkKind(kind)

	if err != nil {
		panic(err)
	}

	reader, err := event.Open(filename, builder)

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
		viewer := event.NewTableViewer(layout, events)
		columns := viewer.Header()
		formatter := fmtutil.NewFormatter(columns...)

		for viewer.HasNext() {
			columns := viewer.Next()

			formatter.Write(columns...)
		}

		fmt.Println(formatter.String())
	}
}

func validEventType() []string {
	return []string{
		ac.Alias()[0],
		am.Alias()[0],
		fw.Alias()[0],
		im.Alias()[0],
		ips.Alias()[0],
		li.Alias()[0],
		system.Alias()[0],
		wrs.Alias()[0],
	}
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "events",
		Short: "Display various events",
		Run:   run,
	}

	eventTypeDescription := fmt.Sprintf("Event type (e.g., \"%s\")", strings.Join(validEventType(), "\", \""))

	flags := command.Flags()
	flags.SetInterspersed(true)
	flags.StringVarP(&kind, "kind", "k", "sys", eventTypeDescription)
	flags.BoolVarP(&details, "details", "d", false, "Show details")

	return command
}

package events

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/jonas-fan/dsdd/pkg/fmtutil"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

const eventFormat = `Origin: %v <%v@%v>
Target: %v
Time:   %v
Level:  %v
Event:  %v | %v

%v
`

func serializeSystemEvent(event *diagnostic.SystemEvent) string {
	return fmt.Sprintf(eventFormat,
		event.EventOrigin,
		event.ActionBy,
		event.Manager,
		event.Target,
		event.Time,
		event.Level,
		event.EventId,
		event.Event,
		pretty.Indent(event.Description))
}

func writeSystemEvents(writer io.Writer, events []diagnostic.SystemEvent) {
	for i := len(events) - 1; i >= 0; i-- {
		fmt.Fprintln(writer, serializeSystemEvent(&events[i]))
	}
}

func writeSummarizedSystemEvents(writer io.Writer, events []diagnostic.SystemEvent) {
	formatter := fmtutil.NewFormatter()

	formatter.Write("TIME", "ORIGIN", "LEVEL", "EID", "EVENT")

	for i := len(events) - 1; i >= 0; i-- {
		event := &events[i]

		formatter.Write(event.Time, event.EventOrigin, event.Level, event.EventId, event.Event)
	}

	fmt.Fprintln(writer, formatter.String())
}

func showSystemEvent() {
	file, err := os.Open(filepath.Join("Manager", "hostevents.csv"))

	if err != nil {
		panic(err)
	}

	defer file.Close()

	events, err := diagnostic.ReadSystemEvent(file)

	if err != nil {
		panic(err)
	}

	if summarized {
		writeSummarizedSystemEvents(os.Stdout, events)
	} else {
		writeSystemEvents(os.Stdout, events)
	}
}

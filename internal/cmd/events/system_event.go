package events

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

const eventFormat = `Origin: %v <%v@%v>
Target: %v
Time:   %v
Level:  %v
Event:  %v | %v

%v
`

func serializeSystemEvent(event diagnostic.SystemEvent) string {
	var builder strings.Builder

	fmt.Fprintf(
		&builder,
		eventFormat,
		event.EventOrigin,
		event.ActionBy,
		event.Manager,
		event.Target,
		event.Time,
		event.Level,
		event.EventId,
		event.Event,
		pretty.Indent(event.Description))

	return builder.String()
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

	for i := len(events) - 1; i >= 0; i-- {
		fmt.Println(serializeSystemEvent(events[i]))
	}
}

package events

import (
	"fmt"
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

func serializeSystemEvent(e diagnostic.SystemEvent) string {
	var builder strings.Builder

	fmt.Fprintf(
		&builder,
		eventFormat,
		e.EventOrigin,
		e.ActionBy,
		e.Manager,
		e.Target,
		e.Time,
		e.Level,
		e.EventId,
		e.Event,
		pretty.Indent(e.Description))

	return builder.String()
}

func readSystemEvent() {
	events, err := diagnostic.ReadSystemEvent()

	if err != nil {
		panic(err)
	}

	for _, event := range events {
		fmt.Println(serializeSystemEvent(event))
	}
}

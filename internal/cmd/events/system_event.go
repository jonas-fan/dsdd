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
	events, err := diagnostic.ReadSystemEvent()

	if err != nil {
		panic(err)
	}

	for _, each := range events {
		fmt.Println(serializeSystemEvent(each))
	}
}

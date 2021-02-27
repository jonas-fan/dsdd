package event

import (
	"fmt"
	"os"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/encoding/csv"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

const eventFormat = `Origin: %v <%v@%v>
Target: %v
Time:   %v
Level:  %v
Event:  %v | %v

%v
`

type SystemEvent struct {
	ActionBy    string
	Description string
	Event       string
	EventId     string
	EventOrigin string
	Level       string
	Manager     string
	Target      string
	Time        string
}

func (e *SystemEvent) assign(key string, value string) {
	switch strings.ToLower(key) {
	case "action by":
		e.ActionBy = value
	case "description":
		e.Description = value
	case "event":
		e.Event = value
	case "event id":
		e.EventId = value
	case "event origin":
		e.EventOrigin = value
	case "level":
		e.Level = value
	case "manager":
		e.Manager = value
	case "target":
		e.Target = value
	case "time":
		e.Time = fmt.Sprint(toTime(value))
	}
}

func (e *SystemEvent) Assign(keys []string, values []string) {
	for index, key := range keys {
		e.assign(key, values[index])
	}
}

func (e *SystemEvent) String() string {
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

func readSystemEvent(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	events := []SystemEvent{}
	err = csv.ReadAll(file, &events)

	if err != nil {
		panic(err)
	}

	for _, event := range events {
		fmt.Println(event.String())
	}
}

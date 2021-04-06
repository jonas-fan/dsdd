package system

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

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

const eventFormat = `Origin: %v <%v@%v>
Target: %v
Time:   %v
Level:  %v
Event:  %v | %v

%v
`

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
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	default:
		// don't bother about these
	}
}

// Time implements the `event.Event` interface.
func (e *SystemEvent) Datetime() string {
	return e.Time
}

// String implements the `event.Event` interface.
func (e *SystemEvent) String() string {
	return fmt.Sprintf(eventFormat,
		e.EventOrigin,
		e.ActionBy,
		e.Manager,
		e.Target,
		e.Time,
		e.Level,
		e.EventId,
		e.Event,
		pretty.Indent(e.Description))
}

func Parse(header []string, fields []string) event.Event {
	e := &SystemEvent{}

	for i := range header {
		e.assign(header[i], fields[i])
	}

	return e
}

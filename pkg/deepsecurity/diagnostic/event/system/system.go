package system

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type SystemEvent struct {
	Time        string
	EventOrigin string
	ActionBy    string
	Manager     string
	Target      string
	Level       string
	EventId     string
	Event       string
	Description string
}

const template = `Origin: %v <%v@%v>
Target: %v
Time:   %v
Level:  %v
Event:  %v | %v

%v
`

// Assign implements the `event.Event` interface.
func (e *SystemEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "action by":
		e.ActionBy = value
	case "manager":
		e.Manager = value
	case "target":
		e.Target = value
	case "level":
		e.Level = value
	case "event id":
		e.EventId = value
	case "event":
		e.Event = value
	case "description":
		e.Description = value
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *SystemEvent) String() string {
	return fmt.Sprintf(template,
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

// Datetime implements the `event.Event` interface.
func (e *SystemEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &SystemEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"system", "sys", "host"}
}

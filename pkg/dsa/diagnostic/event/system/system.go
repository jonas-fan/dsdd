package system

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/encoding/csv"
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

var Header = []string{"TIME", "ORIGIN", "LEVEL", "EID", "EVENT"}

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

// Assign implements `csv.Assigner` and helps with SystemEvent initialization.
func (e *SystemEvent) Assign(keys []string, values []string) {
	for index, key := range keys {
		e.assign(key, values[index])
	}
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

// Column implements the `event.Event` interface.
func (e *SystemEvent) Column() []string {
	return []string{e.Time, e.EventOrigin, e.Level, e.EventId, e.Event}
}

// ReadSystemEvent returns the system events from a reader.
func ReadSystemEvent(reader io.Reader) ([]event.Event, error) {
	events := []SystemEvent{}

	if err := csv.ReadAll(reader, &events); err != nil {
		return nil, err
	}

	out := make([]event.Event, len(events))

	for i := range events {
		out[i] = &events[i]
	}

	return out, nil
}

// ReadSystemEventFrom returns the system events from a file.
func ReadSystemEventFrom(filename string) ([]event.Event, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ReadSystemEvent(file)
}

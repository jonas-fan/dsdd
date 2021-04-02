package diagnostic

import (
	"fmt"
	"io"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/encoding/csv"
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
		e.Time = fmt.Sprint(toTime(value).Format("2006-01-02 15:04:05"))
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

// ReadSystemEvent returns the system events from a reader.
func ReadSystemEvent(reader io.Reader) ([]SystemEvent, error) {
	events := []SystemEvent{}

	if err := csv.ReadAll(reader, &events); err != nil {
		return nil, err
	}

	return events, nil
}

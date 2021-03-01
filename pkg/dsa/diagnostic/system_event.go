package diagnostic

import (
	"fmt"
	"os"
	"path/filepath"
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
	field := strings.ToLower(key)

	switch field {
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

// ReadSystemEvent reads the specific system events file then returns a slice.
func ReadSystemEventFrom(filename string) ([]SystemEvent, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	events := []SystemEvent{}

	if err = csv.ReadAll(file, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// ReadSystemEvent reads the system events file
// (e.g., `Manager/hostevents.csv`) then returns a slice.
func ReadSystemEvent() ([]SystemEvent, error) {
	return ReadSystemEventFrom(filepath.Join("Manager", "hostevents.csv"))
}

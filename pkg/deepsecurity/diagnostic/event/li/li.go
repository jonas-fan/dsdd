package li

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type LogInspectionEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	Reason      string
	Severity    string
	Groups      string
	ProgramName string
	Location    string
	Description string
	Event       string
}

const template = `Origin: %v <%v>
Time:    %v
Reason:  %v | %v
Gropus:  %v
Program: %v

%v

%v

%v
`

// Assign implements the `event.Event` interface.
func (e *LogInspectionEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "computer":
		e.Computer = value
	case "reason":
		e.Reason = value
	case "severity":
		e.Severity = value
	case "groups":
		e.Groups = value
	case "program name":
		e.ProgramName = value
	case "location":
		e.Location = value
	case "description":
		e.Description = value
	case "event":
		e.Event = value
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *LogInspectionEvent) String() string {
	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		e.Reason,
		e.Severity,
		e.Groups,
		e.ProgramName,
		pretty.Indent(e.Description),
		pretty.Indent(e.Location),
		pretty.Indent(e.Event),
	)
}

// Datetime implements the `event.Event` interface.
func (e *LogInspectionEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &LogInspectionEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"li", "loginspection"}
}

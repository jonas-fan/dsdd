package integritymonitoring

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type IntegrityMonitoringEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	Reason      string
	Change      string
	Rank        string
	Severity    string
	Type        string
	Key         string
	Description string
	User        string
	Process     string
}

const template = `Origin:   %v <%v>
Time:     %v
Reason:   %v | %v
Change:   %v | %v | %v
By:       %v | %v

%v
`

// Assign implements the `event.Event` interface.
func (e *IntegrityMonitoringEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "computer":
		e.Computer = value
	case "reason":
		e.Reason = value
	case "change":
		e.Change = value
	case "rank":
		e.Rank = value
	case "severity":
		e.Severity = value
	case "type":
		e.Type = value
	case "key":
		e.Key = value
	case "description":
		e.Description = value
	case "user":
		e.User = value
	case "process":
		e.Process = value
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *IntegrityMonitoringEvent) String() string {
	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		e.Reason,
		e.Severity,
		e.Change,
		e.Type,
		e.Key,
		e.User,
		e.Process,
		pretty.Indent(e.Description))
}

// Datetime implements the `event.Event` interface.
func (e *IntegrityMonitoringEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &IntegrityMonitoringEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"integritymonitoring", "integrity", "im"}
}

package wrs

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type WebReputationEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	Risk        string
	Rank        string
	Url         string
}

const template = `Origin: %v <%v>
Time:   %v
Risk:   %v
Rank:   %v

%v
`

// Assign implements the `event.Event` interface.
func (e *WebReputationEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "computer":
		e.Computer = value
	case "risk":
		e.Risk = value
	case "rank":
		e.Rank = value
	case "url":
		e.Url = value
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *WebReputationEvent) String() string {
	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		e.Risk,
		e.Rank,
		pretty.Indent(e.Url),
	)
}

// Datetime implements the `event.Event` interface.
func (e *WebReputationEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &WebReputationEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"wrs", "webreputation"}
}

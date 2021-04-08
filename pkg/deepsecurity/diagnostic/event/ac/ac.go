package ac

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type ApplicationControlEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	Ruleset     string
	Reason      string
	Event       string
	RepeatCount string
	Action      string
	Path        string
	File        string
	Md5         string
	Sha1        string
	Sha256      string
	UserName    string
	ProcessName string
}

const template = `Origin: %v <%v>
Time:   %v
Reason: %v | %v
Event:  %v
By:     %v | %v
Action: %v

%v

%v
%v
%v
`

// Assign implements the `event.Event` interface.
func (e *ApplicationControlEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "computer":
		e.Computer = value
	case "ruleset":
		e.Ruleset = value
	case "reason":
		e.Reason = value
	case "event":
		e.Event = value
	case "repeat count":
		e.RepeatCount = value
	case "action":
		e.Action = value
	case "path":
		e.Path = value
	case "file":
		e.File = value
	case "md5":
		e.Md5 = event.ToLowerOrNA(value)
	case "sha1":
		e.Sha1 = event.ToLowerOrNA(value)
	case "sha256":
		e.Sha256 = event.ToLowerOrNA(value)
	case "user name":
		e.UserName = value
	case "process name":
		e.ProcessName = value
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *ApplicationControlEvent) String() string {
	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		e.Reason,
		e.Ruleset,
		e.Event,
		e.UserName,
		e.ProcessName,
		e.Action,
		pretty.Indent(e.Path+e.File),
		pretty.Indent("md5:"+e.Md5),
		pretty.Indent("sha1:"+e.Sha1),
		pretty.Indent("sha256:"+e.Sha256),
	)
}

// Datetime implements the `event.Event` interface.
func (e *ApplicationControlEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &ApplicationControlEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"ac", "appcontrol", "applicationcontrol"}
}

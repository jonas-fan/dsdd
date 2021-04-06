package applicationcontrol

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event"
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

const eventFormat = `Origin: %v <%v>
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

func (e *ApplicationControlEvent) assign(key string, value string) {
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
		if value == "" {
			value = "n/a"
		}
		e.Md5 = strings.ToLower(value)
	case "sha1":
		if value == "" {
			value = "n/a"
		}
		e.Sha1 = strings.ToLower(value)
	case "sha256":
		if value == "" {
			value = "n/a"
		}
		e.Sha256 = strings.ToLower(value)
	case "user name":
		e.UserName = value
	case "process name":
		e.ProcessName = value
	default:
		// don't bother about these
	}
}

// Time implements the `event.Event` interface.
func (e *ApplicationControlEvent) Datetime() string {
	return e.Time
}

// String implements the `event.Event` interface.
func (e *ApplicationControlEvent) String() string {
	return fmt.Sprintf(eventFormat,
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
		pretty.Indent("sha256:"+e.Sha256))
}

func Parse(header []string, fields []string) event.Event {
	e := &ApplicationControlEvent{}

	for i := range header {
		e.assign(header[i], fields[i])
	}

	return e
}

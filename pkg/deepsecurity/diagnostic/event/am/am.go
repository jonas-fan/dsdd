package am

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type AntiMalwareEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	ScanType    string
	Reason      string
	VirusType   string
	Malware     string
	Action      string
	Infection   string
	Md5         string
	Sha1        string
	Sha256      string
}

const template = `Origin:  %v <%v>
Time:    %v
Reason:  %v | %v
Malware: %v | %v
Action:  %v

%v

%v
%v
%v
`

// Assign implements the `event.Event` interface.
func (e *AntiMalwareEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "computer":
		e.Computer = value
	case "scan type":
		e.ScanType = value
	case "reason":
		e.Reason = value
	case "major virus type":
		e.VirusType = value
	case "malware":
		e.Malware = value
	case "action taken":
		e.Action = value
	case "infected file(s)":
		e.Infection = value
	case "file md5":
		e.Md5 = event.ToLowerOrNA(value)
	case "file sha-1":
		e.Sha1 = event.ToLowerOrNA(value)
	case "file sha-256":
		e.Sha256 = event.ToLowerOrNA(value)
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *AntiMalwareEvent) String() string {
	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		e.Reason,
		e.ScanType,
		e.VirusType,
		e.Malware,
		e.Action,
		pretty.Indent(e.Infection),
		pretty.Indent("md5:"+e.Md5),
		pretty.Indent("sha1:"+e.Sha1),
		pretty.Indent("sha256:"+e.Sha256),
	)
}

// Datetime implements the `event.Event` interface.
func (e *AntiMalwareEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &AntiMalwareEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"am", "anti-malware", "antimalware"}
}

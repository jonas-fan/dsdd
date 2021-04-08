package ips

import (
	"fmt"
	"strings"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
	"github.com/jonas-fan/dsdd/pkg/pretty"
)

type NetInterface struct {
	IP   string
	Port string
	MAC  string
}

type Packet struct {
	Protocol string
	Flags    string
	Data     PackatData
}

type PackatData struct {
	Size string
}

type IntrusionPreventionEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	Reason      string
	Action      string
	Direction   string
	Flow        string
	Interface   string
	Source      NetInterface
	Destination NetInterface
	Packet      Packet
}

const template = `Origin:    %v <%v>
Time:      %v
Reason:    %v
Action:    %v
Direction: %v
Flow:      %v

%v
%v
%v
%v

%v
%v

%v
%v

%v
%v
`

// Assign implements the `event.Event` interface.
func (e *IntrusionPreventionEvent) Assign(key string, value string) {
	switch strings.ToLower(key) {
	case "time":
		e.Time = fmt.Sprint(event.ToTime(value).Format("2006-01-02 15:04:05"))
	case "event origin":
		e.EventOrigin = value
	case "computer":
		e.Computer = value
	case "reason":
		e.Reason = value
	case "action":
		e.Action = value
	case "direction":
		e.Direction = value
	case "flow":
		e.Flow = value
	case "source ip":
		e.Source.IP = value
	case "source port":
		e.Source.Port = value
	case "source mac":
		e.Source.MAC = value
	case "destination ip":
		e.Destination.IP = value
	case "destination port":
		e.Destination.Port = value
	case "destination mac":
		e.Destination.MAC = value
	case "protocol":
		e.Packet.Protocol = value
	case "flags":
		e.Packet.Flags = value
	case "packet size":
		e.Packet.Data.Size = value
	default:
		// don't bother about these
	}
}

// String implements the `event.Event` interface.
func (e *IntrusionPreventionEvent) String() string {
	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		pretty.Indent(""),
	)
}

// Datetime implements the `event.Event` interface.
func (e *IntrusionPreventionEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &IntrusionPreventionEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"ips", "intrusion", "intrusionprevention"}
}

package fw

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

type Frame struct {
	Type string
}

type Packet struct {
	Protocol string
	Flags    string
	Data     PackatData
}

type PackatData struct {
	Size string
}

type FirewallEvent struct {
	Time        string
	EventOrigin string
	Computer    string
	Reason      string
	Action      string
	Direction   string
	Source      NetInterface
	Destination NetInterface
	Frame       Frame
	Packet      Packet
}

const template = `Origin:    %v <%v>
Time:      %v
Reason:    %v
Action:    %v
Direction: %v

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
func (e *FirewallEvent) Assign(key string, value string) {
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
	case "frame type":
		e.Frame.Type = value
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
func (e *FirewallEvent) String() string {
	source := fmt.Sprintf("%s:%s (%s)", e.Source.IP, e.Source.Port, e.Source.MAC)
	destination := fmt.Sprintf("%s:%s (%s)", e.Destination.IP, e.Destination.Port, e.Destination.MAC)

	return fmt.Sprintf(template,
		e.EventOrigin,
		e.Computer,
		e.Time,
		e.Reason,
		e.Action,
		e.Direction,
		pretty.Indent("Packet:"),
		pretty.Indent("    Frame Type: "+e.Frame.Type),
		pretty.Indent("    Protocol:   "+e.Packet.Protocol),
		pretty.Indent("    Flags:      "+e.Packet.Flags),
		pretty.Indent("Packet Data:"),
		pretty.Indent("    Size:       "+e.Packet.Data.Size),
		pretty.Indent("Source:"),
		pretty.Indent("    "+source),
		pretty.Indent("Destination:"),
		pretty.Indent("    "+destination),
	)
}

// Datetime implements the `event.Event` interface.
func (e *FirewallEvent) Datetime() string {
	return e.Time
}

// New returns a new `event.Event`.
func New() event.Event {
	return &FirewallEvent{}
}

// Alias returns alias of this pacakge.
func Alias() []string {
	return []string{"fw", "firewall"}
}

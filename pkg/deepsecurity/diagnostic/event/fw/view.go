package fw

import (
	"fmt"

	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
)

type TableLayout struct {
	header []string
}

func (v *TableLayout) Header() []string {
	return v.header
}

func (v *TableLayout) Columns(event event.Event) []string {
	e := event.(*FirewallEvent)

	return []string{
		e.Time,
		e.Reason,
		e.Action,
		e.Direction,
		fmt.Sprintf("%s:%s", e.Source.IP, e.Source.Port),
		fmt.Sprintf("%s:%s", e.Destination.IP, e.Destination.Port),
		e.Frame.Type,
		e.Packet.Protocol,
		e.Packet.Flags,
		e.Packet.Data.Size,
	}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{
			"TIME",
			"REASON",
			"ACTION",
			"DIRECTION",
			"SOURCE",
			"DESTINATION",
			"TYPE",
			"PROTO",
			"FLAGS",
			"DATA",
		},
	}
}

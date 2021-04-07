package integritymonitoring

import (
	"github.com/jonas-fan/dsdd/pkg/deepsecurity/diagnostic/event"
)

type TableLayout struct {
	header []string
}

func (v *TableLayout) Header() []string {
	return v.header
}

func (v *TableLayout) Columns(event event.Event) []string {
	e := event.(*IntegrityMonitoringEvent)

	return []string{e.Time, e.Reason, e.Change, e.Process, e.Type, e.Key}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{"TIME", "REASON", "CHANGE", "BY", "TYPE", "KEY"},
	}
}

package system

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
	e := event.(*SystemEvent)

	return []string{e.Time, e.EventOrigin, e.Level, e.EventId, e.Event}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{"TIME", "ORIGIN", "LEVEL", "EVENT ID", "EVENT"},
	}
}

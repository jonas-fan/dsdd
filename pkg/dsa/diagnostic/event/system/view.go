package system

import (
	"github.com/jonas-fan/dsdd/pkg/dsa/diagnostic/event"
)

type TableViewer struct {
	header  []string
	columns []event.Event
	index   int
}

func (v *TableViewer) Header() []string {
	return v.header
}

func (v *TableViewer) HasNext() bool {
	return v.index < len(v.columns)
}

func (v *TableViewer) Next() []string {
	e := v.columns[v.index].(*SystemEvent)

	v.index++

	return []string{e.Time, e.EventOrigin, e.Level, e.EventId, e.Event}
}

func NewTableViewer(events []event.Event) event.TableViewer {
	return &TableViewer{
		header:  []string{"TIME", "ORIGIN", "LEVEL", "EID", "EVENT"},
		columns: events,
	}
}

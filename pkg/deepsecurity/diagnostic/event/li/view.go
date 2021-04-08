package li

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
	e := event.(*LogInspectionEvent)

	return []string{e.Time, e.Reason, e.Severity, e.Description}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{"TIME", "REASON", "SEVERITY", "DESCRIPTION"},
	}
}

package wrs

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
	e := event.(*WebReputationEvent)

	return []string{e.Time, e.Risk, e.Rank, e.Url}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{"TIME", "RISK", "RANK", "URL"},
	}
}

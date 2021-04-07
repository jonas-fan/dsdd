package applicationcontrol

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
	e := event.(*ApplicationControlEvent)
	digest := fmt.Sprintf("%.12s", e.Sha256)
	file := e.Path + e.File

	return []string{e.Time, e.Reason, e.Event, e.Action, digest, file}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{"TIME", "REASON", "EVENT", "ACTION", "DIGEST", "PATH"},
	}
}

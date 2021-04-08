package am

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
	e := event.(*AntiMalwareEvent)
	malware := fmt.Sprintf("%s:%s", e.VirusType, e.Malware)

	return []string{e.Time, e.Reason, e.ScanType, malware, e.Action, e.Infection}
}

func NewTableLayout() event.TableLayout {
	return &TableLayout{
		header: []string{"TIME", "REASON", "BY", "MALWARE", "ACTION", "PATH"},
	}
}

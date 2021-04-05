package event

type TableLayout interface {
	Header() []string
	Columns(event Event) []string
}

type TableViewer struct {
	layout  TableLayout
	columns []Event
	index   int
}

func (v *TableViewer) Header() []string {
	return v.layout.Header()
}

func (v *TableViewer) HasNext() bool {
	return v.index < len(v.columns)
}

func (v *TableViewer) Next() []string {
	e := v.columns[v.index]

	v.index++

	return v.layout.Columns(e)
}

func NewTableViewer(layout TableLayout, events []Event) *TableViewer {
	return &TableViewer{
		layout:  layout,
		columns: events,
	}
}

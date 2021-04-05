package event

type TableViewer interface {
	Header() []string
	HasNext() bool
	Next() []string
}

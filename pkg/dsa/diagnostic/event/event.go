package event

import (
	"encoding/csv"
	"os"
)

type Event interface {
	Assign(key string, value string)
	String() string
	Datetime() string
}

type EventBuilder func() Event

type Reader struct {
	file    *os.File
	reader  *csv.Reader
	builder EventBuilder
	header  []string
}

// Open returns a new event reader.
func Open(filename string, builder EventBuilder) (*Reader, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	reader := &Reader{
		file:    file,
		reader:  csv.NewReader(file),
		builder: builder,
	}

	return reader, nil
}

// Close closes the file descriptor.
func (r *Reader) Close() {
	r.file.Close()
}

func (r *Reader) build(header []string, fields []string) Event {
	event := r.builder()

	for i := range header {
		event.Assign(header[i], fields[i])
	}

	return event
}

// Read returns the next record.
func (r *Reader) Read() (Event, error) {
	fields, err := r.reader.Read()

	if err != nil {
		return nil, err
	}

	if r.header == nil {
		r.header = fields

		return r.Read()
	}

	return r.build(r.header, fields), nil
}

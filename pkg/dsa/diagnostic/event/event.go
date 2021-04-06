package event

import (
	"encoding/csv"
	"os"
)

type Event interface {
	Datetime() string
	String() string
}

type Parser func(header []string, fields []string) Event

type Reader struct {
	file   *os.File
	reader *csv.Reader
	parser Parser
	header []string
}

// Open returns a new event reader.
func Open(filename string, parser Parser) (*Reader, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	reader := &Reader{
		file:   file,
		reader: csv.NewReader(file),
		parser: parser,
	}

	return reader, nil
}

// Close closes the file descriptor.
func (r *Reader) Close() {
	r.file.Close()
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

	return r.parser(r.header, fields), nil
}

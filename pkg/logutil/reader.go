package logutil

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type SplitFunc func(string) bool

type Reader struct {
	scanner *bufio.Scanner
	split   SplitFunc
	caches  []string
	next    string
}

func makeScanFunc(pattern string) SplitFunc {
	var express = regexp.MustCompile(pattern)

	return func(token string) bool {
		return express.MatchString(token)
	}
}

// HasNext checks whether the next token is ready to read.
func (r *Reader) HasNext() bool {
	var line string

	for r.scanner.Scan() {
		line = r.scanner.Text()

		if len(r.caches) > 0 && r.split(line) {
			break
		}

		r.caches = append(r.caches, line)
	}

	if len(r.caches) == 0 {
		return false
	}

	r.next = strings.Join(r.caches, "\n")
	r.caches = r.caches[:0]

	if len(line) > 0 {
		r.caches = append(r.caches, line)
	}

	return true
}

// Next returns the next item.
func (r *Reader) Next() string {
	return r.next
}

// NewReader returns a new Reader.
func NewReader(reader io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(reader),
		split:   makeScanFunc(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`),
		caches:  make([]string, 0),
	}
}

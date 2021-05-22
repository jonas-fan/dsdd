package fmtutil

import (
	"fmt"
	"strings"
)

type Align int

const (
	LeftAlign Align = iota
	RightAlign
)

type Format struct {
	width int
	align Align
}

type Formatter struct {
	formats []Format
	lines   [][]string
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func align2rune(align Align) rune {
	switch align {
	case LeftAlign:
		return '-'
	case RightAlign:
		return '+'
	default:
		panic("invalid argument")
	}
}

func (f *Formatter) makeFormats() []string {
	out := make([]string, len(f.formats))
	last := len(out) - 1

	for index := range out {
		align := f.formats[index].align
		width := f.formats[index].width

		if (index != last) || (align == RightAlign) {
			out[index] = fmt.Sprintf("%%%c%ds", align2rune(align), width)
		} else {
			out[index] = "%s"
		}
	}

	return out
}

// Align sets the alignment on the index-th element.
func (f *Formatter) Align(index int, align Align) {
	if index >= len(f.formats) {
		panic("index out of range")
	}

	f.formats[index].align = align
}

// Write writes data to formatter.
func (f *Formatter) Write(args ...string) {
	if f.formats == nil {
		f.formats = make([]Format, len(args))
		f.lines = make([][]string, 0)
	} else if len(f.formats) != len(args) {
		panic("unexpected number of arguments")
	}

	for index, arg := range args {
		format := &f.formats[index]

		format.width = max(format.width, len(arg))
	}

	f.lines = append(f.lines, args)
}

// String represents the formatted string.
func (f *Formatter) String() string {
	var builder strings.Builder
	formats := f.makeFormats()

	for index, line := range f.lines {
		if index > 0 {
			fmt.Fprintln(&builder)
		}

		for index, token := range line {
			if index > 0 {
				fmt.Fprintf(&builder, "  ")
			}

			fmt.Fprintf(&builder, formats[index], token)
		}
	}

	return builder.String()
}

// NewFormatter returns a new formatter.
func NewFormatter(args ...string) *Formatter {
	formatter := &Formatter{}

	if len(args) > 0 {
		formatter.Write(args...)
	}

	return formatter
}

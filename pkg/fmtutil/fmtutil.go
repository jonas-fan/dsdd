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

type Formatter struct {
	widths []int
	aligns []Align
	lines  [][]string
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func format(align Align) rune {
	switch align {
	case LeftAlign:
		return '-'
	case RightAlign:
		return '+'
	default:
		panic("invalid argument")
	}
}

func (f *Formatter) formats() []string {
	out := make([]string, len(f.widths))
	last := len(f.widths) - 1

	for index, width := range f.widths {
		align := f.aligns[index]

		if (index != last) || (align == RightAlign) {
			out[index] = fmt.Sprintf("%%%c%ds", format(align), width)
		} else {
			out[index] = "%s"
		}
	}

	return out
}

// Align sets the alignment on the index-th element.
func (f *Formatter) Align(index int, align Align) {
	if index >= len(f.aligns) {
		panic("index out of range")
	}

	f.aligns[index] = align
}

// Write writes data to formatter.
func (f *Formatter) Write(args ...string) {
	if f.widths == nil {
		f.widths = make([]int, len(args))
		f.aligns = make([]Align, len(args))
		f.lines = make([][]string, 0)
	} else if len(f.widths) != len(args) {
		panic("unexpected number of arguments")
	}

	for index, arg := range args {
		f.widths[index] = max(f.widths[index], len(arg))
	}

	f.lines = append(f.lines, args)
}

// String represents the formatted string.
func (f *Formatter) String() string {
	var builder strings.Builder
	formats := f.formats()

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

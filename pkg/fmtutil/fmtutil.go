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
	lines [][]string
	width []int
	align []Align
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func (f *Formatter) formats() []string {
	out := make([]string, len(f.width))
	align := []rune{
		LeftAlign:  '-',
		RightAlign: '+',
	}

	for index, width := range f.width {
		out[index] = fmt.Sprintf("%%%c%ds", align[f.align[index]], width)
	}

	out[len(out)-1] = "%s"

	return out
}

func (f *Formatter) Align(index int, align Align) {
	if index >= len(f.align) {
		panic("index out of range")
	}

	f.align[index] = align
}

func (f *Formatter) Write(args ...string) {
	if f.width == nil {
		f.lines = make([][]string, 0)
		f.width = make([]int, len(args))
		f.align = make([]Align, len(args))
	} else if len(f.width) != len(args) {
		panic("Unexpected number of arguments")
	}

	for index, arg := range args {
		f.width[index] = max(f.width[index], len(arg))
	}

	f.lines = append(f.lines, args)
}

func (f *Formatter) String() string {
	lines := make([]string, 0, len(f.lines))
	tokens := make([]string, len(f.width))
	formats := f.formats()

	for _, line := range f.lines {
		for index, token := range line {
			tokens[index] = fmt.Sprintf(formats[index], token)
		}

		lines = append(lines, strings.Join(tokens, "  "))
	}

	return strings.Join(lines, "\n")
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

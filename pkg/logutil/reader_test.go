package logutil

import (
	"strings"
	"testing"
)

func TestReaderSingleLineLog(t *testing.T) {
	lines := []string{
		"2020-12-19 15:30:01.123456 [+0800] | foo",
		"2020-12-19 15:30:02.123457 [+0800] | bar",
		"2020-12-19 15:30:03.123458 [+0800] | baz",
	}

	blob := strings.Join(lines, "\n")
	reader := NewReader(strings.NewReader(blob))

	for count := 0; reader.HasNext(); count++ {
		if reader.Next() != lines[count] {
			t.Error("failure")
		}
	}

	t.Log("success")
}

func TestReaderMultipleLinesLog(t *testing.T) {
	lines := []string{
		"2020-12-19 15:30:01.123456 [+0800] | foo",
		"2020-12-19 15:30:02.123457 [+0800] | bar\nbaz\nqux\n\n\n",
		"2020-12-19 15:30:03.123458 [+0800] | foobar",
	}

	blob := strings.Join(lines, "\n")
	reader := NewReader(strings.NewReader(blob))

	for count := 0; reader.HasNext(); count++ {
		if reader.Next() != lines[count] {
			t.Error("failure")
		}
	}

	t.Log("success")
}

package streamio

import (
	"bytes"
	"context"
	"testing"
)

const NEWLINE byte = '\n'

func pull(from *Stream) []byte {
	out := make([]byte, 0)

	for each := range from.Chan() {
		if len(out) > 0 {
			out = append(out, NEWLINE)
		}

		out = append(out, each...)
	}

	return out
}

func testStream(blob []byte) bool {
	stream := NewStream(context.Background(), bytes.NewReader(blob))

	return bytes.Equal(blob, pull(stream))
}

func TestStreamSingleLine(t *testing.T) {
	lines := [][]byte{
		[]byte("2020-12-19 15:30:01.123456 [+0800] | foo"),
	}

	blob := bytes.Join(lines, []byte{NEWLINE})

	if testStream(blob) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

func TestStreamMultipleLines(t *testing.T) {
	lines := [][]byte{
		[]byte("2020-12-19 15:30:01.123456 [+0800] | foo"),
		[]byte("2020-12-19 15:30:02.123457 [+0800] | bar"),
		[]byte("2020-12-19 15:30:03.123458 [+0800] | baz"),
	}

	blob := bytes.Join(lines, []byte{NEWLINE})

	if testStream(blob) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

func TestStreamMultipleLinesWithDelimiter(t *testing.T) {
	lines := [][]byte{
		[]byte("2020-12-19 15:30:01.123456 [+0800] | foo"),
		[]byte("2020-12-19 15:30:02.123457 [+0800] | bar\nbaz\nqux\n\n\n"),
		[]byte("2020-12-19 15:30:03.123458 [+0800] | baz"),
	}

	blob := bytes.Join(lines, []byte{NEWLINE})

	if testStream(blob) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

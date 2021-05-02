package fmtutil

import (
	"testing"
)

func TestWriteNothing(t *testing.T) {
	f := NewFormatter()

	expected := ""

	if f.String() != expected {
		t.Error("unexpected result")
	}
}

func TestWithoutWrite(t *testing.T) {
	f := NewFormatter("1", "2", "3")

	expected := "1  2  3"

	if f.String() != expected {
		t.Error("unexpected result")
	}
}

func TestWriteSingleLine(t *testing.T) {
	f := NewFormatter()
	f.Write("1", "2", "3")

	expected := "1  2  3"

	if f.String() != expected {
		t.Error("unexpected result")
	}
}

func TestWriteMultipleLines(t *testing.T) {
	f := NewFormatter()
	f.Write("1", "2", "3")
	f.Write("111", "222", "333")
	f.Write("11", "22", "33")

	expected := `1    2    3
111  222  333
11   22   33`

	if f.String() != expected {
		t.Error("unexpected result")
	}
}

func TestWritePanicMismatchedNumOfArgs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not panic")
		}
	}()

	f := NewFormatter()
	f.Write("1", "2", "3")
	f.Write("1", "2", "3", "4")
}

func TestAlignSingleLine(t *testing.T) {
	f := NewFormatter()
	f.Write("1", "2", "3")
	f.Align(0, RightAlign)
	f.Align(1, RightAlign)
	f.Align(2, LeftAlign)

	expected := "1  2  3"

	if f.String() != expected {
		t.Error("unexpected result")
	}
}

func TestAlignMultipleLines(t *testing.T) {
	f := NewFormatter()
	f.Write("111", "222", "333")
	f.Write("11", "22", "33")
	f.Write("1", "2", "3")
	f.Align(0, RightAlign)
	f.Align(1, LeftAlign)
	f.Align(2, RightAlign)

	expected := `111  222  333
 11  22    33
  1  2      3`

	if f.String() != expected {
		t.Error("unexpected result")
	}
}

func TestAlignPanicOutOfRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not panic")
		}
	}()

	f := NewFormatter("1", "2", "3")
	f.Align(3, RightAlign)
}

func TestAlignPanicInvalidArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not panic")
		}
	}()

	f := NewFormatter("1", "2", "3")
	f.Align(0, 100)
	f.String()
}

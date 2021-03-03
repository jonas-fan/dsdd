package fmtutil

import (
	"testing"
)

func TestWriteSingleLine(t *testing.T) {
	f := NewFormatter()
	f.Write("1", "2", "3")

	expected := `1  2  3`

	if f.String() != expected {
		t.Error("Unexpected result")
	}
}

func TestWriteMultipleLines01(t *testing.T) {
	f := NewFormatter()
	f.Write("1", "2", "3")
	f.Write("11", "22", "33")
	f.Write("111", "222", "333")

	expected := `1    2    3
11   22   33
111  222  333`

	if f.String() != expected {
		t.Error("Unexpected result")
	}
}

func TestWriteMultipleLines02(t *testing.T) {
	f := NewFormatter()
	f.Write("111", "222", "333")
	f.Write("11", "22", "33")
	f.Write("1", "2", "3")

	expected := `111  222  333
11   22   33
1    2    3`

	if f.String() != expected {
		t.Error("Unexpected result")
	}
}

func TestWritePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic")
		}
	}()

	f := NewFormatter()
	f.Write("1", "2", "3")
	f.Write("1", "2", "3", "4")
}

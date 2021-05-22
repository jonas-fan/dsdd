package pretty

import (
	"testing"
)

func check(t *testing.T, out string, expected string) {
	if out != expected {
		t.Errorf("expected: %s, got: %s", expected, out)
	}
}

func TestIndentEmptyLine(t *testing.T) {
	check(t, Indent(""), "    ")
}

func TestIndentSingleLine(t *testing.T) {
	check(t, Indent("hello"), "    hello")
}

func TestIndentMultipleLines(t *testing.T) {
	check(t, Indent("hello\nworld"), "    hello\n    world")
}

func TestIndentWithEmptyLine(t *testing.T) {
	check(t, IndentWith("", "\t"), "\t")
}

func TestIndentWithSingleLine(t *testing.T) {
	check(t, IndentWith("hello", "\t"), "\thello")
}

func TestIndentWithMultipleLines(t *testing.T) {
	check(t, IndentWith("hello\nworld", "\t"), "\thello\n\tworld")
}

package pretty

import (
	"strings"
)

// IndentWith adds specified spaces in front of each line of message.
func IndentWith(message string, spaces string) string {
	return spaces + strings.ReplaceAll(message, "\n", "\n"+spaces)
}

// Indent adds 4 spaces in front of each line of message.
func Indent(message string) string {
	return IndentWith(message, "    ")
}

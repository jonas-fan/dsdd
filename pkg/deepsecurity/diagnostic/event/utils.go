package event

import (
	"strings"
	"time"
)

var timeLayouts = [...]string{
	"January 2, 2006 15:04:05",
	"January 2, 2006 15:04:05 PM",
}

// ToTime converts a string to time.
func ToTime(value string) time.Time {
	for _, layout := range timeLayouts {
		if out, err := time.Parse(layout, value); err == nil {
			return out
		}
	}

	return time.Time{}
}

// ToLowerOrNA converts a string to lowercase or n/a.
func ToLowerOrNA(value string) string {
	if value == "" {
		return "n/a"
	}

	return strings.ToLower(value)
}

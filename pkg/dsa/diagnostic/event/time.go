package event

import (
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

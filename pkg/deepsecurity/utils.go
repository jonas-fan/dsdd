package deepsecurity

import (
	"strings"
)

// ToOS represents an operating system string based on a given platform.
func ToOS(platform string) string {
	platform = strings.ToLower(platform)

	switch {
	case strings.HasPrefix(platform, "aix"):
		return "AIX"
	case strings.HasPrefix(platform, "solaris"):
		return "Solaris"
	case strings.HasPrefix(platform, "windows"):
		return "Windows"
	default:
		return "Linux"
	}
}

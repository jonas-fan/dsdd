package deepsecurity

import (
	"strings"
)

func ToOS(platform string) string {
	platform = strings.ToLower(platform)

	switch {
	case strings.HasPrefix(platform, "aix"):
		return "aix"
	case strings.HasPrefix(platform, "solaris"):
		return "solaris"
	case strings.HasPrefix(platform, "windows"):
		return "windows"
	default:
		return "linux"
	}
}

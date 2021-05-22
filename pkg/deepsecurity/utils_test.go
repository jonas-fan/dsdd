package deepsecurity

import (
	"testing"
)

func check(t *testing.T, out string, expected string) {
	if out != expected {
		t.Errorf("expected: %s, got: %s", expected, out)
	}
}

func TestToOSLinux(t *testing.T) {
	check(t, ToOS("Amazon"), "Linux")
	check(t, ToOS("Debian"), "Linux")
	check(t, ToOS("RedHat"), "Linux")
	check(t, ToOS("Ubuntu"), "Linux")
}

func TestToOSWindows(t *testing.T) {
	check(t, ToOS("windows"), "Windows")
	check(t, ToOS("WINDOWS"), "Windows")
}

func TestToOSAIX(t *testing.T) {
	check(t, ToOS("aix"), "AIX")
	check(t, ToOS("AIX"), "AIX")
}

func TestToOSSolaris(t *testing.T) {
	check(t, ToOS("solaris"), "Solaris")
	check(t, ToOS("SOLARIS"), "Solaris")
}

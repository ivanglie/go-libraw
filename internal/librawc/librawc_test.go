//go:build cgo

package librawc

import (
	"strings"
	"testing"
)

func TestLinkedVersion(t *testing.T) {
	version := Version()
	if strings.TrimSpace(version) == "" {
		t.Fatal("Version() is empty")
	}

	number := VersionNumber()
	if number <= 0 {
		t.Fatalf("VersionNumber() = %d, want positive", number)
	}

	t.Logf("linked LibRaw version: %s (%d)", version, number)
}

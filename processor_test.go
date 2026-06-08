package libraw

import (
	"strings"
	"testing"
)

func TestNewProcessorCloseIsIdempotent(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if p.Closed() {
		t.Fatal("new processor reports closed")
	}

	if err := p.Close(); err != nil {
		t.Fatalf("first Close() error = %v", err)
	}
	if !p.Closed() {
		t.Fatal("processor reports open after Close")
	}
	if err := p.Close(); err != nil {
		t.Fatalf("second Close() error = %v", err)
	}
}

func TestNilProcessorCloseIsSafe(t *testing.T) {
	var p *Processor
	if err := p.Close(); err != nil {
		t.Fatalf("nil Close() error = %v", err)
	}
	if !p.Closed() {
		t.Fatal("nil Processor should report closed")
	}
}

func TestVersion(t *testing.T) {
	version := Version()
	if strings.TrimSpace(version) == "" {
		t.Fatal("Version() is empty")
	}
	if VersionNumber() <= 0 {
		t.Fatalf("VersionNumber() = %d, want positive", VersionNumber())
	}
}

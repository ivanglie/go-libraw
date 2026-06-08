//go:build !cgo

// Package librawc contains the internal cgo bridge to LibRaw.
package librawc

import "errors"

// ErrCGODisabled reports that cgo is required for LibRaw bindings.
var ErrCGODisabled = errors.New("libraw: cgo is disabled; enable cgo and install LibRaw development headers")

// Handle is an unavailable LibRaw handle when cgo is disabled.
type Handle struct{}

// New reports that cgo is required for LibRaw.
func New(uint) (*Handle, error) {
	return nil, ErrCGODisabled
}

// Close releases the unavailable handle.
func (h *Handle) Close() {}

// Version returns an empty version when cgo is disabled.
func Version() string {
	return ""
}

// VersionNumber returns zero when cgo is disabled.
func VersionNumber() int {
	return 0
}

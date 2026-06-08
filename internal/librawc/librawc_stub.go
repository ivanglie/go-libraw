//go:build !cgo

// Package librawc contains the internal cgo bridge to LibRaw.
package librawc

import "errors"

// ErrCGODisabled reports that cgo is required for LibRaw bindings.
var ErrCGODisabled = errors.New("libraw: cgo is disabled; enable cgo and install LibRaw development headers")

// errorCodeCGODisabled is a non-zero LibRaw status returned by stub open paths so
// callers surface a typed error. These paths are unreachable in practice because
// New already fails when cgo is disabled.
const errorCodeCGODisabled = -1

// Handle is an unavailable LibRaw handle when cgo is disabled.
type Handle struct{}

// New reports that cgo is required for LibRaw.
func New(uint) (*Handle, error) {
	return nil, ErrCGODisabled
}

// Close releases the unavailable handle.
func (h *Handle) Close() {}

// OpenFile reports that cgo is required for LibRaw.
func (h *Handle) OpenFile(string) int {
	return int(errorCodeCGODisabled)
}

// OpenBuffer reports that cgo is required for LibRaw.
func (h *Handle) OpenBuffer([]byte) int {
	return int(errorCodeCGODisabled)
}

// BayerParams holds the geometry and decoding flags for OpenBayer.
type BayerParams struct {
	RawWidth     uint16
	RawHeight    uint16
	LeftMargin   uint16
	TopMargin    uint16
	RightMargin  uint16
	BottomMargin uint16
	ProcFlags    uint8
	BayerPattern uint8
	UnusedBits   uint
	OtherFlags   uint
	BlackLevel   uint
}

// OpenBayer reports that cgo is required for LibRaw.
func (h *Handle) OpenBayer([]byte, BayerParams) int {
	return int(errorCodeCGODisabled)
}

// Recycle is a no-op when cgo is disabled.
func (h *Handle) Recycle() {}

// RecycleDatastream is a no-op when cgo is disabled.
func (h *Handle) RecycleDatastream() {}

// Version returns an empty version when cgo is disabled.
func Version() string {
	return ""
}

// VersionNumber returns zero when cgo is disabled.
func VersionNumber() int {
	return 0
}

// StrError returns an empty string when cgo is disabled.
func StrError(int) string {
	return ""
}

// StrProgress returns an empty string when cgo is disabled.
func StrProgress(int) string {
	return ""
}

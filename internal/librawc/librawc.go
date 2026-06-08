//go:build cgo

// Package librawc contains the internal cgo bridge to LibRaw.
package librawc

/*
#cgo darwin CFLAGS: -I/opt/homebrew/opt/libraw/include
#cgo darwin LDFLAGS: -L/opt/homebrew/opt/libraw/lib -lraw
#cgo linux LDFLAGS: -lraw
#include <libraw/libraw.h>
*/
import "C"

import "errors"

// ErrInitFailed reports that LibRaw returned a nil handle during initialization.
var ErrInitFailed = errors.New("libraw: libraw_init returned nil")

// Handle wraps a libraw_data_t pointer.
type Handle struct {
	ptr *C.libraw_data_t
}

// New initializes a LibRaw handle.
func New(flags uint) (*Handle, error) {
	ptr := C.libraw_init(C.uint(flags))
	if ptr == nil {
		return nil, ErrInitFailed
	}
	return &Handle{ptr: ptr}, nil
}

// Close releases the LibRaw handle.
func (h *Handle) Close() {
	if h == nil || h.ptr == nil {
		return
	}
	C.libraw_close(h.ptr)
	h.ptr = nil
}

// Version returns the linked LibRaw runtime version string.
func Version() string {
	return C.GoString(C.libraw_version())
}

// VersionNumber returns the linked LibRaw runtime version number.
func VersionNumber() int {
	return int(C.libraw_versionNumber())
}

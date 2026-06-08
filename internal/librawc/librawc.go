//go:build cgo

// Package librawc contains the internal cgo bridge to LibRaw.
package librawc

/*
#cgo linux pkg-config: libraw
#cgo darwin,arm64 CFLAGS: -I/opt/homebrew/opt/libraw/include
#cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/libraw/lib -lraw
#cgo darwin,amd64 CFLAGS: -I/usr/local/opt/libraw/include
#cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/libraw/lib -lraw
#include <stdlib.h>
#include <libraw/libraw.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

// ErrInitFailed reports that LibRaw returned a nil handle during initialization.
var ErrInitFailed = errors.New("libraw: libraw_init returned nil")

// Handle wraps a libraw_data_t pointer.
//
// LibRaw's buffer and Bayer open paths keep a pointer to the input bytes rather
// than copying them, so the bytes must outlive processing. To keep ownership on
// the C side and satisfy cgo pointer rules, OpenBuffer and OpenBayer copy the
// input into C memory retained here. The copy is freed when the handle is
// recycled, closed, or reused for another buffer.
type Handle struct {
	ptr  *C.libraw_data_t
	cbuf unsafe.Pointer
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
	h.freeBuffer()
}

// freeBuffer releases the retained C copy of an input buffer, if any.
func (h *Handle) freeBuffer() {
	if h.cbuf != nil {
		C.free(h.cbuf)
		h.cbuf = nil
	}
}

// OpenFile opens a RAW file by path and returns the LibRaw status code.
func (h *Handle) OpenFile(path string) int {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return int(C.libraw_open_file(h.ptr, cpath))
}

// OpenBuffer opens RAW bytes held in memory and returns the LibRaw status code.
//
// The bytes are copied into C memory retained by the handle, so the caller's
// slice does not need to outlive the call.
func (h *Handle) OpenBuffer(data []byte) int {
	h.freeBuffer()
	if len(data) == 0 {
		return int(C.libraw_open_buffer(h.ptr, nil, 0))
	}
	h.cbuf = C.CBytes(data)
	return int(C.libraw_open_buffer(h.ptr, h.cbuf, C.size_t(len(data))))
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

// OpenBayer opens raw Bayer samples and returns the LibRaw status code.
//
// The bytes are copied into C memory retained by the handle because LibRaw keeps
// a pointer to them; the copy is freed on recycle, close, or reuse.
func (h *Handle) OpenBayer(data []byte, p BayerParams) int {
	h.freeBuffer()
	var ptr *C.uchar
	if len(data) > 0 {
		h.cbuf = C.CBytes(data)
		ptr = (*C.uchar)(h.cbuf)
	}
	return int(C.libraw_open_bayer(
		h.ptr,
		ptr,
		C.uint(len(data)),
		C.ushort(p.RawWidth),
		C.ushort(p.RawHeight),
		C.ushort(p.LeftMargin),
		C.ushort(p.TopMargin),
		C.ushort(p.RightMargin),
		C.ushort(p.BottomMargin),
		C.uchar(p.ProcFlags),
		C.uchar(p.BayerPattern),
		C.uint(p.UnusedBits),
		C.uint(p.OtherFlags),
		C.uint(p.BlackLevel),
	))
}

// Recycle resets LibRaw state so the handle can open another input, releasing
// any retained input buffer.
func (h *Handle) Recycle() {
	if h == nil || h.ptr == nil {
		return
	}
	C.libraw_recycle(h.ptr)
	h.freeBuffer()
}

// RecycleDatastream releases the open input datastream while keeping decoded
// state, matching libraw_recycle_datastream.
func (h *Handle) RecycleDatastream() {
	if h == nil || h.ptr == nil {
		return
	}
	C.libraw_recycle_datastream(h.ptr)
}

// Version returns the linked LibRaw runtime version string.
func Version() string {
	return C.GoString(C.libraw_version())
}

// VersionNumber returns the linked LibRaw runtime version number.
func VersionNumber() int {
	return int(C.libraw_versionNumber())
}

// StrError returns the LibRaw message for an error code.
func StrError(code int) string {
	return C.GoString(C.libraw_strerror(C.int(code)))
}

// StrProgress returns the LibRaw message for a progress stage.
func StrProgress(progress int) string {
	return C.GoString(C.libraw_strprogress(C.enum_LibRaw_progress(progress)))
}

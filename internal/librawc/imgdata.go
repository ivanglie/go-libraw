//go:build cgo

package librawc

/*
#include <stdlib.h>
#include <libraw/libraw.h>
*/
import "C"

import "unsafe"

// Color returns LibRaw's color index for the sensor pixel at (row, col),
// mapping the camera's color filter array. It wraps libraw_COLOR.
func (h *Handle) Color(row, col int) int {
	return int(C.libraw_COLOR(h.ptr, C.int(row), C.int(col)))
}

// RawWidth returns the raw image width in pixels (libraw_get_raw_width).
func (h *Handle) RawWidth() int {
	return int(C.libraw_get_raw_width(h.ptr))
}

// RawHeight returns the raw image height in pixels (libraw_get_raw_height).
func (h *Handle) RawHeight() int {
	return int(C.libraw_get_raw_height(h.ptr))
}

// RawImage returns a Go copy of the single-channel raw Bayer buffer
// (imgdata.rawdata.raw_image), or nil if no such buffer is available.
//
// The buffer is row-padded: its length is (raw_pitch/2)*raw_height samples.
func (h *Handle) RawImage() []uint16 {
	img := h.ptr.rawdata.raw_image
	if img == nil {
		return nil
	}
	height := int(h.ptr.sizes.raw_height)
	pitch := int(h.ptr.sizes.raw_pitch)
	if pitch == 0 {
		pitch = int(h.ptr.sizes.raw_width) * 2
	}
	n := (pitch / 2) * height
	if n <= 0 {
		return nil
	}
	src := unsafe.Slice((*uint16)(unsafe.Pointer(img)), n)
	out := make([]uint16, n)
	copy(out, src)
	return out
}

// FourChannels returns a Go copy of the 4-channel postprocessing image buffer
// (imgdata.image) as a flat slice of [4]uint16 pixels, or nil when the buffer
// is not available (Raw2Image or DcrawProcess must have run first).
//
// The length is iheight*iwidth. Channel assignment follows the CFA pattern
// (typically RGBG for Bayer sensors).
func (h *Handle) FourChannels() [][4]uint16 {
	if h.ptr.image == nil {
		return nil
	}
	n := int(h.ptr.sizes.iheight) * int(h.ptr.sizes.iwidth)
	if n <= 0 {
		return nil
	}
	src := unsafe.Slice(h.ptr.image, n)
	out := make([][4]uint16, n)
	for i, px := range src {
		out[i] = [4]uint16{uint16(px[0]), uint16(px[1]), uint16(px[2]), uint16(px[3])}
	}
	return out
}

// ThumbnailData returns a Go copy of the unpacked thumbnail bytes
// (imgdata.thumbnail.thumb), or nil if no thumbnail data is present.
func (h *Handle) ThumbnailData() []byte {
	thumb := h.ptr.thumbnail.thumb
	n := int(h.ptr.thumbnail.tlength)
	if thumb == nil || n <= 0 {
		return nil
	}
	src := unsafe.Slice((*byte)(unsafe.Pointer(thumb)), n)
	out := make([]byte, n)
	copy(out, src)
	return out
}

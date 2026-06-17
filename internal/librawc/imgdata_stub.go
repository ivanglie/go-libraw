//go:build !cgo

package librawc

// Color returns zero when cgo is disabled.
func (h *Handle) Color(int, int) int { return 0 }

// RawWidth returns zero when cgo is disabled.
func (h *Handle) RawWidth() int { return 0 }

// RawHeight returns zero when cgo is disabled.
func (h *Handle) RawHeight() int { return 0 }

// RawImage returns nil when cgo is disabled.
func (h *Handle) RawImage() []uint16 { return nil }

// FourChannels returns nil when cgo is disabled.
func (h *Handle) FourChannels() [][4]uint16 { return nil }

// ThumbnailData returns nil when cgo is disabled.
func (h *Handle) ThumbnailData() []byte { return nil }

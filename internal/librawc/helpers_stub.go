//go:build !cgo

package librawc

// CameraList returns nil when cgo is disabled.
func CameraList() []string { return nil }

// CameraCount returns zero when cgo is disabled.
func CameraCount() int { return 0 }

// Capabilities returns zero when cgo is disabled.
func Capabilities() uint { return 0 }

// UnpackFunctionName returns an empty string when cgo is disabled.
func (h *Handle) UnpackFunctionName() string { return "" }

// DecoderInfo returns zero values when cgo is disabled.
func (h *Handle) DecoderInfo() (string, uint32, int) { return "", 0, 0 }

// IWidth returns zero when cgo is disabled.
func (h *Handle) IWidth() int { return 0 }

// IHeight returns zero when cgo is disabled.
func (h *Handle) IHeight() int { return 0 }

// CamMul returns zero when cgo is disabled.
func (h *Handle) CamMul(int) float32 { return 0 }

// PreMul returns zero when cgo is disabled.
func (h *Handle) PreMul(int) float32 { return 0 }

// RGBCam returns zero when cgo is disabled.
func (h *Handle) RGBCam(int, int) float32 { return 0 }

// ColorMaximum returns zero when cgo is disabled.
func (h *Handle) ColorMaximum() int { return 0 }

// AdjustToRawInsetCrop reports unsupported when cgo is disabled.
func (h *Handle) AdjustToRawInsetCrop(uint, float32) (int, bool) { return 0, false }

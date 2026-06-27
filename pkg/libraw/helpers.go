package libraw

import (
	"errors"

	"github.com/strawberryworks/go-libraw/internal/librawc"
)

const (
	camMulLen  = 4 // camera/pre multipliers
	rgbCamRows = 3 // rgb_cam matrix rows
	rgbCamCols = 4 // rgb_cam matrix columns
)

// ErrUnsupported reports that the operation requires a newer LibRaw than the one
// the binary is linked against.
var ErrUnsupported = errors.New("libraw: operation not supported by the linked LibRaw")

// DecoderInfo describes the decoder LibRaw selected for the open image.
type DecoderInfo struct {
	Name  string
	Flags uint32
}

// CameraList returns the camera models LibRaw supports. The slice length equals
// CameraCount. The list is static; it is rebuilt on each call rather than cached.
func CameraList() []string { return librawc.CameraList() }

// CameraCount returns the number of camera models LibRaw supports.
func CameraCount() int { return librawc.CameraCount() }

// Capabilities returns LibRaw's runtime capability flags.
func Capabilities() uint { return librawc.Capabilities() }

// UnpackFunctionName returns the name of the unpack function for the open image.
func (p *Processor) UnpackFunctionName() (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return "", ErrClosed
	}
	return p.handle.UnpackFunctionName(), nil
}

// DecoderInfo returns the decoder name and flags for the open image.
func (p *Processor) DecoderInfo() (DecoderInfo, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return DecoderInfo{}, ErrClosed
	}
	name, flags, code := p.handle.DecoderInfo()
	if err := ToError(ErrorCode(code)); err != nil {
		return DecoderInfo{}, err
	}
	return DecoderInfo{Name: name, Flags: flags}, nil
}

// IWidth returns the output image width in pixels.
func (p *Processor) IWidth() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.IWidth(), nil
}

// IHeight returns the output image height in pixels.
func (p *Processor) IHeight() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.IHeight(), nil
}

// CamMul returns the camera white-balance multiplier at index (0..3).
func (p *Processor) CamMul(index int) (float32, error) {
	if err := validateIndex("camera multiplier", index, camMulLen); err != nil {
		return 0, err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.CamMul(index), nil
}

// PreMul returns the pre-multiplier at index (0..3).
func (p *Processor) PreMul(index int) (float32, error) {
	if err := validateIndex("pre multiplier", index, camMulLen); err != nil {
		return 0, err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.PreMul(index), nil
}

// RGBCam returns the rgb_cam color matrix element at (row, col), where row is in
// 0..2 and col in 0..3.
func (p *Processor) RGBCam(row, col int) (float32, error) {
	if err := validateIndex("rgb_cam row", row, rgbCamRows); err != nil {
		return 0, err
	}
	if err := validateIndex("rgb_cam column", col, rgbCamCols); err != nil {
		return 0, err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.RGBCam(row, col), nil
}

// ColorMaximum returns the maximum color value for the open image.
func (p *Processor) ColorMaximum() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.ColorMaximum(), nil
}

// AdjustToRawInsetCrop adjusts the output crop to the raw inset using the given
// mask and maximum crop fraction. It requires LibRaw 0.22+; against an older
// LibRaw it returns ErrUnsupported.
func (p *Processor) AdjustToRawInsetCrop(mask uint, maxCrop float32) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	code, supported := p.handle.AdjustToRawInsetCrop(mask, maxCrop)
	if !supported {
		return ErrUnsupported
	}
	return ToError(ErrorCode(code))
}

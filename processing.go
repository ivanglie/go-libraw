package libraw

import "github.com/ivanglie/go-libraw/internal/librawc"

// ProcessedImage is an in-memory image or thumbnail produced by MemImage or
// MemThumb. Data is owned by Go; the underlying C allocation is freed before the
// value is returned, so there is nothing for the caller to release.
type ProcessedImage struct {
	// Type is a LIBRAW_IMAGE_* format constant (bitmap, JPEG, ...).
	Type int
	// Height and Width are the image dimensions in pixels.
	Height uint16
	Width  uint16
	// Colors is the number of color components per pixel.
	Colors uint16
	// Bits is the bit depth per component.
	Bits uint16
	// Data holds the raw image or compressed thumbnail bytes.
	Data []byte
}

// Unpack decodes the RAW pixel data of the opened input.
func (p *Processor) Unpack() error {
	return p.withHandle(func(h *librawc.Handle) int { return h.Unpack() })
}

// UnpackThumb decodes the embedded thumbnail of the opened input.
func (p *Processor) UnpackThumb() error {
	return p.withHandle(func(h *librawc.Handle) int { return h.UnpackThumb() })
}

// UnpackThumbEx decodes the thumbnail at the given index.
func (p *Processor) UnpackThumbEx(index int) error {
	return p.withHandle(func(h *librawc.Handle) int { return h.UnpackThumbEx(index) })
}

// SubtractBlack applies LibRaw's black-level subtraction pass.
//
// It operates on the postprocessing image buffer, so Raw2Image (or DcrawProcess)
// must run first. Calling it on an unbuilt buffer is undefined in LibRaw and can
// hang or crash.
func (p *Processor) SubtractBlack() error {
	return p.withHandleVoid(func(h *librawc.Handle) { h.SubtractBlack() })
}

// Raw2Image copies unpacked RAW data into the postprocessing image buffer.
//
// Unpack must run first.
func (p *Processor) Raw2Image() error {
	return p.withHandle(func(h *librawc.Handle) int { return h.Raw2Image() })
}

// FreeImage releases the postprocessing image buffer.
func (p *Processor) FreeImage() error {
	return p.withHandleVoid(func(h *librawc.Handle) { h.FreeImage() })
}

// AdjustSizesInfoOnly recomputes output sizes without producing an image.
func (p *Processor) AdjustSizesInfoOnly() error {
	return p.withHandle(func(h *librawc.Handle) int { return h.AdjustSizesInfoOnly() })
}

// DcrawProcess runs LibRaw's dcraw-equivalent postprocessing.
func (p *Processor) DcrawProcess() error {
	return p.withHandle(func(h *librawc.Handle) int { return h.DcrawProcess() })
}

// WritePPMTiff writes the processed image to path. The format is PPM or TIFF
// depending on the output parameters (PPM by default).
func (p *Processor) WritePPMTiff(path string) error {
	return p.withHandle(func(h *librawc.Handle) int { return h.DcrawPPMTiffWriter(path) })
}

// WriteThumb writes the unpacked thumbnail to path in its native format.
func (p *Processor) WriteThumb(path string) error {
	return p.withHandle(func(h *librawc.Handle) int { return h.DcrawThumbWriter(path) })
}

// MemImage renders the processed image into memory.
//
// DcrawProcess must run first. The returned image owns its bytes.
func (p *Processor) MemImage() (*ProcessedImage, error) {
	return p.makeMem((*librawc.Handle).MakeMemImage)
}

// MemThumb renders the unpacked thumbnail into memory.
//
// UnpackThumb must run first. The returned image owns its bytes.
func (p *Processor) MemThumb() (*ProcessedImage, error) {
	return p.makeMem((*librawc.Handle).MakeMemThumb)
}

// makeMem runs a mem-image producer under the lock and converts the result.
func (p *Processor) makeMem(fn func(*librawc.Handle) (librawc.ProcessedImage, int)) (*ProcessedImage, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return nil, ErrClosed
	}
	img, code := fn(p.handle)
	if err := ToError(ErrorCode(code)); err != nil {
		return nil, err
	}
	return &ProcessedImage{
		Type:   img.Type,
		Height: img.Height,
		Width:  img.Width,
		Colors: img.Colors,
		Bits:   img.Bits,
		Data:   img.Data,
	}, nil
}

// withHandleVoid runs a void handle operation under the lock, returning ErrClosed
// when the Processor is closed.
func (p *Processor) withHandleVoid(fn func(*librawc.Handle)) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	fn(p.handle)
	return nil
}

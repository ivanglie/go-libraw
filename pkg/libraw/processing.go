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
//
// An input must be opened first; otherwise Unpack returns ErrInvalidState.
func (p *Processor) Unpack() error {
	return p.staged(
		func() error { return p.requireState("Unpack", stateOpened) },
		func(h *librawc.Handle) int { return h.Unpack() },
		func() { p.advance(stateUnpacked) },
	)
}

// UnpackThumb decodes the embedded thumbnail of the opened input.
//
// An input must be opened first; otherwise UnpackThumb returns ErrInvalidState.
func (p *Processor) UnpackThumb() error {
	return p.staged(
		func() error { return p.requireState("UnpackThumb", stateOpened) },
		func(h *librawc.Handle) int { return h.UnpackThumb() },
		func() { p.thumbReady = true },
	)
}

// UnpackThumbEx decodes the thumbnail at the given index.
//
// An input must be opened first; otherwise UnpackThumbEx returns ErrInvalidState.
func (p *Processor) UnpackThumbEx(index int) error {
	return p.staged(
		func() error { return p.requireState("UnpackThumbEx", stateOpened) },
		func(h *librawc.Handle) int { return h.UnpackThumbEx(index) },
		func() { p.thumbReady = true },
	)
}

// SubtractBlack applies LibRaw's black-level subtraction pass.
//
// It operates on the postprocessing image buffer, so Raw2Image (or DcrawProcess)
// must run first. Calling it on an unbuilt buffer is undefined in LibRaw and can
// hang or crash, so SubtractBlack returns ErrInvalidState in that case rather
// than forwarding the call.
func (p *Processor) SubtractBlack() error {
	return p.stagedVoid(
		func() error { return p.requireState("SubtractBlack", stateImageBuilt) },
		func(h *librawc.Handle) { h.SubtractBlack() },
		nil,
	)
}

// Raw2Image copies unpacked RAW data into the postprocessing image buffer.
//
// Unpack must run first; otherwise Raw2Image returns ErrInvalidState.
func (p *Processor) Raw2Image() error {
	return p.staged(
		func() error { return p.requireState("Raw2Image", stateUnpacked) },
		func(h *librawc.Handle) int { return h.Raw2Image() },
		func() { p.advance(stateImageBuilt) },
	)
}

// FreeImage releases the postprocessing image buffer, returning the pipeline to
// its unpacked state so a fresh image must be built before reuse.
//
// Raw2Image (or DcrawProcess) must run first; otherwise FreeImage returns
// ErrInvalidState.
func (p *Processor) FreeImage() error {
	return p.stagedVoid(
		func() error { return p.requireState("FreeImage", stateImageBuilt) },
		func(h *librawc.Handle) { h.FreeImage() },
		func() { p.state = stateUnpacked },
	)
}

// AdjustSizesInfoOnly recomputes output sizes without producing an image.
//
// An input must be opened first; otherwise it returns ErrInvalidState.
func (p *Processor) AdjustSizesInfoOnly() error {
	return p.staged(
		func() error { return p.requireState("AdjustSizesInfoOnly", stateOpened) },
		func(h *librawc.Handle) int { return h.AdjustSizesInfoOnly() },
		nil,
	)
}

// DcrawProcess runs LibRaw's dcraw-equivalent postprocessing.
//
// Unpack must run first; otherwise DcrawProcess returns ErrInvalidState.
func (p *Processor) DcrawProcess() error {
	return p.staged(
		func() error { return p.requireState("DcrawProcess", stateUnpacked) },
		func(h *librawc.Handle) int { return h.DcrawProcess() },
		func() { p.advance(stateProcessed) },
	)
}

// WritePPMTiff writes the processed image to path. The format is PPM or TIFF
// depending on the output parameters (PPM by default).
//
// DcrawProcess must run first; otherwise WritePPMTiff returns ErrInvalidState.
func (p *Processor) WritePPMTiff(path string) error {
	return p.staged(
		func() error { return p.requireState("WritePPMTiff", stateProcessed) },
		func(h *librawc.Handle) int { return h.DcrawPPMTiffWriter(path) },
		nil,
	)
}

// WriteThumb writes the unpacked thumbnail to path in its native format.
//
// UnpackThumb must run first; otherwise WriteThumb returns ErrInvalidState.
func (p *Processor) WriteThumb(path string) error {
	return p.staged(
		func() error { return p.requireThumb("WriteThumb") },
		func(h *librawc.Handle) int { return h.DcrawThumbWriter(path) },
		nil,
	)
}

// MemImage renders the processed image into memory.
//
// DcrawProcess must run first; otherwise MemImage returns ErrInvalidState. The
// returned image owns its bytes.
func (p *Processor) MemImage() (*ProcessedImage, error) {
	return p.makeMem(
		func() error { return p.requireState("MemImage", stateProcessed) },
		(*librawc.Handle).MakeMemImage,
	)
}

// MemThumb renders the unpacked thumbnail into memory.
//
// UnpackThumb must run first; otherwise MemThumb returns ErrInvalidState. The
// returned image owns its bytes.
func (p *Processor) MemThumb() (*ProcessedImage, error) {
	return p.makeMem(
		func() error { return p.requireThumb("MemThumb") },
		(*librawc.Handle).MakeMemThumb,
	)
}

// makeMem runs a mem-image producer under the lock and converts the result. It
// rejects the call with precond before forwarding to LibRaw.
func (p *Processor) makeMem(precond func() error, fn func(*librawc.Handle) (librawc.ProcessedImage, int)) (*ProcessedImage, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return nil, ErrClosed
	}
	if precond != nil {
		if err := precond(); err != nil {
			return nil, err
		}
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

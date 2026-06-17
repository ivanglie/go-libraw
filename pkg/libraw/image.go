package libraw

import "errors"

// ErrNoImageData reports that the requested image, raw, or thumbnail buffer is
// not available. Typically the relevant decode step has not run yet, or the
// format does not provide that buffer.
var ErrNoImageData = errors.New("libraw: no image data available")

// Color returns LibRaw's color index for the sensor pixel at (row, col),
// mapping the camera's color filter array (libraw_COLOR). It is valid after the
// input is opened.
func (p *Processor) Color(row, col int) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.Color(row, col), nil
}

// RawWidth returns the raw image width in pixels.
func (p *Processor) RawWidth() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.RawWidth(), nil
}

// RawHeight returns the raw image height in pixels.
func (p *Processor) RawHeight() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return 0, ErrClosed
	}
	return p.handle.RawHeight(), nil
}

// RawImage returns a copy of the single-channel raw Bayer buffer.
//
// Unpack must run first. The returned slice is owned by the caller and remains
// valid after the Processor is closed. The data is row-padded; its length is
// (raw_pitch/2)*raw_height samples. ErrNoImageData is returned when the opened
// image has no single-channel raw buffer (for example, non-Bayer formats).
func (p *Processor) RawImage() ([]uint16, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return nil, ErrClosed
	}
	data := p.handle.RawImage()
	if data == nil {
		return nil, ErrNoImageData
	}
	return data, nil
}

// FourChannels returns a copy of the 4-channel postprocessing image buffer
// built by Raw2Image or DcrawProcess.
//
// Length is IHeight*IWidth (see Metadata().Sizes). Channel assignment depends
// on the CFA pattern; for Bayer sensors it is typically RGBG. ErrNoImageData
// is returned when the buffer is empty (Raw2Image was not run or is not
// applicable to this image type).
func (p *Processor) FourChannels() ([][4]uint16, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return nil, ErrClosed
	}
	if err := p.requireState("FourChannels", stateImageBuilt); err != nil {
		return nil, err
	}
	data := p.handle.FourChannels()
	if data == nil {
		return nil, ErrNoImageData
	}
	return data, nil
}

// ThumbnailData returns a copy of the unpacked thumbnail bytes.
//
// UnpackThumb must run first; otherwise ThumbnailData returns ErrInvalidState.
// The returned slice is owned by the caller and remains valid after the
// Processor is closed. ErrNoImageData is returned when no thumbnail data is
// present despite the unpack step completing.
func (p *Processor) ThumbnailData() ([]byte, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return nil, ErrClosed
	}
	if err := p.requireThumb("ThumbnailData"); err != nil {
		return nil, err
	}
	data := p.handle.ThumbnailData()
	if data == nil {
		return nil, ErrNoImageData
	}
	return data, nil
}

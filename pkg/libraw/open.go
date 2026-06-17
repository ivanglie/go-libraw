package libraw

import "github.com/ivanglie/go-libraw/internal/librawc"

// BayerParams describes the geometry and decoding flags for OpenBayer.
//
// The field set mirrors libraw_open_bayer. BayerPattern accepts the
// LIBRAW_OPENBAYER_* constants.
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

// OpenFile opens a RAW file by path.
//
// On success LibRaw metadata becomes available for later helpers. Opening a new
// input replaces any input previously opened on this Processor.
func (p *Processor) OpenFile(path string) error {
	return p.open(func(h *librawc.Handle) int {
		return h.OpenFile(path)
	})
}

// OpenBuffer opens RAW bytes held in memory.
//
// The bytes are copied into memory owned by LibRaw, so the caller may reuse or
// release data once OpenBuffer returns.
func (p *Processor) OpenBuffer(data []byte) error {
	return p.open(func(h *librawc.Handle) int {
		return h.OpenBuffer(data)
	})
}

// OpenBayer opens raw Bayer samples described by params.
//
// As with OpenBuffer, the bytes are copied into LibRaw-owned memory.
func (p *Processor) OpenBayer(data []byte, params BayerParams) error {
	return p.open(func(h *librawc.Handle) int {
		return h.OpenBayer(data, librawc.BayerParams(params))
	})
}

// open runs a LibRaw open operation under the lock. LibRaw recycles any prior
// input before parsing, so the pipeline is reset before the call regardless of
// outcome: a failed open leaves the Processor with no input (stateInit) rather
// than stale readiness, and a successful open advances to stateOpened.
func (p *Processor) open(fn func(*librawc.Handle) int) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	p.state = stateInit
	p.thumbReady = false
	if err := ToError(ErrorCode(fn(p.handle))); err != nil {
		return err
	}
	p.state = stateOpened
	return nil
}

// Recycle resets the Processor so it can open another input without allocating a
// new handle. Metadata and decoded buffers from the previous input are cleared.
func (p *Processor) Recycle() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	p.handle.Recycle()
	p.state = stateInit
	p.thumbReady = false
	return nil
}

// RecycleDatastream releases the open input datastream while keeping decoded
// state, matching libraw_recycle_datastream.
func (p *Processor) RecycleDatastream() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	p.handle.RecycleDatastream()
	return nil
}

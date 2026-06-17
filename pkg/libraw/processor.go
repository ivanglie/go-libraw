package libraw

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ivanglie/go-libraw/internal/librawc"
)

// ErrClosed reports that an operation requires an open Processor.
var ErrClosed = errors.New("libraw: processor is closed")

// ErrInvalidState reports that a pipeline operation was called before its
// prerequisites were met (for example SubtractBlack before Raw2Image). LibRaw
// is undefined on an unbuilt buffer and can hang or crash, so the binding
// rejects out-of-order calls instead of forwarding them.
var ErrInvalidState = errors.New("libraw: operation called out of order")

// pipelineState tracks how far the LibRaw processing pipeline has advanced for
// the current input. States are ordered: each stage requires a minimum state
// and advances it. Thumbnail decoding is tracked separately (see thumbReady).
type pipelineState int

const (
	stateInit       pipelineState = iota // handle created, no input opened
	stateOpened                          // input opened, metadata available
	stateUnpacked                        // RAW pixel data decoded (Unpack)
	stateImageBuilt                      // postprocessing buffer built (Raw2Image)
	stateProcessed                       // dcraw postprocessing done (DcrawProcess)
)

// String returns a human-readable name used in ErrInvalidState messages.
func (s pipelineState) String() string {
	switch s {
	case stateInit:
		return "no input opened"
	case stateOpened:
		return "input opened"
	case stateUnpacked:
		return "RAW unpacked"
	case stateImageBuilt:
		return "image built"
	case stateProcessed:
		return "image processed"
	default:
		return "unknown"
	}
}

// Option configures a Processor at construction time.
type Option func(*options)

type options struct {
	flags           uint
	outputParams    *OutputParams
	rawUnpackParams *RawUnpackParams
}

// WithFlags sets the flags passed to LibRaw during handle initialization.
func WithFlags(flags uint) Option {
	return func(o *options) {
		o.flags = flags
	}
}

// Processor owns a LibRaw processing handle.
//
// Processor methods are safe to call concurrently for lifecycle operations.
// Future image-processing methods may have narrower concurrency guarantees when
// they map to mutable LibRaw operations.
type Processor struct {
	mu         sync.Mutex
	handle     *librawc.Handle
	closed     bool
	state      pipelineState
	thumbReady bool
}

// NewProcessor creates a LibRaw processor handle.
func NewProcessor(opts ...Option) (*Processor, error) {
	var cfg options
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}

	handle, err := librawc.New(cfg.flags)
	if err != nil {
		return nil, fmt.Errorf("libraw: create processor: %w", err)
	}
	if cfg.outputParams != nil {
		handle.SetOutputParams(librawc.OutputParams(*cfg.outputParams))
	}
	if cfg.rawUnpackParams != nil {
		if len([]byte(cfg.rawUnpackParams.P4ShotOrder)) > p4ShotOrderLen {
			handle.Close()
			return nil, fmt.Errorf("libraw: P4ShotOrder length %d exceeds %d bytes", len([]byte(cfg.rawUnpackParams.P4ShotOrder)), p4ShotOrderLen)
		}
		handle.SetRawUnpackParams(librawc.RawUnpackParams(*cfg.rawUnpackParams))
	}

	return &Processor{handle: handle}, nil
}

// Close releases the underlying LibRaw handle.
//
// Close is idempotent. Calling Close on an already closed Processor returns nil.
func (p *Processor) Close() error {
	if p == nil {
		return nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil
	}

	if p.handle != nil {
		p.handle.Close()
		p.handle = nil
	}
	p.closed = true
	return nil
}

// Closed reports whether the Processor has been closed.
func (p *Processor) Closed() bool {
	if p == nil {
		return true
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	return p.closed
}

// requireState verifies the pipeline has reached min. The caller must hold p.mu.
func (p *Processor) requireState(op string, min pipelineState) error {
	if p.state < min {
		return fmt.Errorf("%w: %s requires %s, current state is %q", ErrInvalidState, op, min, p.state)
	}
	return nil
}

// requireThumb verifies a thumbnail has been unpacked. The caller must hold p.mu.
func (p *Processor) requireThumb(op string) error {
	if !p.thumbReady {
		return fmt.Errorf("%w: %s requires UnpackThumb first", ErrInvalidState, op)
	}
	return nil
}

// advance raises the pipeline state to s, never lowering it. The caller must
// hold p.mu.
func (p *Processor) advance(s pipelineState) {
	if s > p.state {
		p.state = s
	}
}

// staged runs an int-returning pipeline stage under the lock. It rejects the
// call with precond before forwarding to LibRaw, then runs post on success.
// precond and post may be nil.
func (p *Processor) staged(precond func() error, fn func(*librawc.Handle) int, post func()) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	if precond != nil {
		if err := precond(); err != nil {
			return err
		}
	}
	if err := ToError(ErrorCode(fn(p.handle))); err != nil {
		return err
	}
	if post != nil {
		post()
	}
	return nil
}

// withHandleVoid runs a void handle operation under the lock, returning
// ErrClosed when the Processor is closed. It carries no pipeline-state guard
// and is used by parameter setters and callback registration, which are valid
// at any point in the lifecycle.
func (p *Processor) withHandleVoid(fn func(*librawc.Handle)) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	fn(p.handle)
	return nil
}

// stagedVoid is staged for LibRaw operations that return no status code.
func (p *Processor) stagedVoid(precond func() error, fn func(*librawc.Handle), post func()) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return ErrClosed
	}
	if precond != nil {
		if err := precond(); err != nil {
			return err
		}
	}
	fn(p.handle)
	if post != nil {
		post()
	}
	return nil
}

// Version returns the linked LibRaw runtime version string.
func Version() string {
	return librawc.Version()
}

// VersionNumber returns the linked LibRaw runtime version number.
func VersionNumber() int {
	return librawc.VersionNumber()
}

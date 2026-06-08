package libraw

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ivanglie/go-libraw/internal/librawc"
)

// ErrClosed reports that an operation requires an open Processor.
var ErrClosed = errors.New("libraw: processor is closed")

// Option configures a Processor at construction time.
type Option func(*options)

type options struct {
	flags uint
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
	mu     sync.Mutex
	handle *librawc.Handle
	closed bool
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

// Version returns the linked LibRaw runtime version string.
func Version() string {
	return librawc.Version()
}

// VersionNumber returns the linked LibRaw runtime version number.
func VersionNumber() int {
	return librawc.VersionNumber()
}

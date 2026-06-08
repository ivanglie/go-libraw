// Package libraw provides Go bindings for LibRaw.
//
// This module is intentionally split into a small public API and an internal
// cgo bridge. Public types own Go-level lifecycle and error semantics. The
// internal bridge owns direct calls into LibRaw and should not be imported by
// consumers.
//
// LibRaw itself must be available to the Go toolchain when cgo is enabled. On
// macOS with Homebrew, the initial bridge looks under /opt/homebrew/opt/libraw.
// More complete build discovery is tracked separately in the workflowr queue.
//
// Processor values own a LibRaw handle. Call Close when finished. Close is
// idempotent, so deferred cleanup is safe even if earlier cleanup already ran.
package libraw

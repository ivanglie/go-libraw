package libraw

import "github.com/ivanglie/go-libraw/internal/librawc"

// Callback threading, panics, and lifetime:
//
//   - LibRaw invokes these callbacks synchronously, on the goroutine that calls
//     the triggering method (OpenFile, Unpack, DcrawProcess, ...). They are not
//     run concurrently by this package, but LibRaw makes no promise about
//     ordering or frequency beyond its own behavior.
//   - A panic inside a callback is recovered at the C boundary so it never
//     crashes the process. A panicking ProgressFunc is treated as cancellation
//     (the processing call returns LIBRAW_CANCELLED_BY_CALLBACK); a panic in any
//     other callback is recovered and ignored.
//   - Handlers are released when the Processor is closed. LibRaw's recycle also
//     clears its C-side handlers, so re-register after Recycle if needed.

// ProgressFunc receives LibRaw progress events. Returning a non-zero value
// cancels the in-flight processing call, which then returns a typed error.
type ProgressFunc func(stage Progress, iteration, expected int) int

// DataErrorFunc receives LibRaw I/O error notifications with the file name and
// byte offset where the error occurred.
type DataErrorFunc func(file string, offset int64)

// TagParserFunc receives EXIF or maker-note tag events during identification.
// The underlying LibRaw stream is not exposed; only tag metadata is provided.
type TagParserFunc func(tag, typ, length int, order uint32, base int64)

// SetProgressHandler registers fn to receive progress events. Pass nil to
// disable the handler.
func (p *Processor) SetProgressHandler(fn ProgressFunc) error {
	return p.withHandleVoid(func(h *librawc.Handle) {
		if fn == nil {
			h.SetProgressCallback(nil)
			return
		}
		h.SetProgressCallback(func(stage, iteration, expected int) int {
			return fn(Progress(stage), iteration, expected)
		})
	})
}

// SetDataErrorHandler registers fn to receive I/O error notifications. Pass nil
// to disable the handler.
func (p *Processor) SetDataErrorHandler(fn DataErrorFunc) error {
	return p.withHandleVoid(func(h *librawc.Handle) {
		if fn == nil {
			h.SetDataErrorCallback(nil)
			return
		}
		h.SetDataErrorCallback(librawc.DataErrorFunc(fn))
	})
}

// SetExifParserHandler registers fn to receive EXIF tag events. Pass nil to
// disable the handler.
func (p *Processor) SetExifParserHandler(fn TagParserFunc) error {
	return p.withHandleVoid(func(h *librawc.Handle) {
		if fn == nil {
			h.SetExifCallback(nil)
			return
		}
		h.SetExifCallback(librawc.ExifParserFunc(fn))
	})
}

// SetMakerNotesHandler registers fn to receive maker-note tag events. Pass nil
// to disable the handler.
//
// LibRaw exposed libraw_set_makernotes_handler in 0.22; when linked against an
// older LibRaw this registers fn but the callback never fires.
func (p *Processor) SetMakerNotesHandler(fn TagParserFunc) error {
	return p.withHandleVoid(func(h *librawc.Handle) {
		if fn == nil {
			h.SetMakernotesCallback(nil)
			return
		}
		h.SetMakernotesCallback(librawc.ExifParserFunc(fn))
	})
}

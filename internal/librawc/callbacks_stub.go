//go:build !cgo

package librawc

// ProgressFunc receives LibRaw progress events. A non-zero return cancels the
// current processing call.
type ProgressFunc func(stage, iteration, expected int) int

// DataErrorFunc receives LibRaw I/O error notifications.
type DataErrorFunc func(file string, offset int64)

// ExifParserFunc receives EXIF or maker-note tag events.
type ExifParserFunc func(tag, typ, length int, order uint32, base int64)

// SetProgressCallback is a no-op when cgo is disabled.
func (h *Handle) SetProgressCallback(ProgressFunc) {}

// SetDataErrorCallback is a no-op when cgo is disabled.
func (h *Handle) SetDataErrorCallback(DataErrorFunc) {}

// SetExifCallback is a no-op when cgo is disabled.
func (h *Handle) SetExifCallback(ExifParserFunc) {}

// SetMakernotesCallback is a no-op when cgo is disabled.
func (h *Handle) SetMakernotesCallback(ExifParserFunc) {}

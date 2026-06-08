//go:build cgo

package librawc

/*
#include <stdint.h>
#include <libraw/libraw.h>

extern int goLibrawProgress(void *data, int stage, int iter, int expected);
extern void goLibrawDataError(void *data, const char *file, long long offset);
extern void goLibrawExif(void *data, int tag, int type, int len,
                         unsigned int ord, void *ifp, long long base);
extern void goLibrawMakernotes(void *data, int tag, int type, int len,
                               unsigned int ord, void *ifp, long long base);

static void go_set_progress(libraw_data_t *lr, uintptr_t d) {
	libraw_set_progress_handler(lr, (progress_callback)goLibrawProgress, (void *)d);
}
static void go_set_dataerror(libraw_data_t *lr, uintptr_t d) {
	libraw_set_dataerror_handler(lr, (data_callback)goLibrawDataError, (void *)d);
}
static void go_set_exif(libraw_data_t *lr, uintptr_t d) {
	libraw_set_exifparser_handler(lr, (exif_parser_callback)goLibrawExif, (void *)d);
}
static void go_set_makernotes(libraw_data_t *lr, uintptr_t d) {
	libraw_set_makernotes_handler(lr, (exif_parser_callback)goLibrawMakernotes, (void *)d);
}
*/
import "C"

import "runtime/cgo"

// ProgressFunc receives LibRaw progress events. A non-zero return cancels the
// current processing call.
type ProgressFunc func(stage, iteration, expected int) int

// DataErrorFunc receives LibRaw I/O error notifications.
type DataErrorFunc func(file string, offset int64)

// ExifParserFunc receives EXIF or maker-note tag events. The underlying LibRaw
// stream pointer is not exposed.
type ExifParserFunc func(tag, typ, length int, order uint32, base int64)

// callbacks holds the Go functions registered for one handle. It is referenced
// by a cgo.Handle whose value is passed to LibRaw as the callback data pointer.
type callbacks struct {
	progress   ProgressFunc
	dataError  DataErrorFunc
	exif       ExifParserFunc
	makernotes ExifParserFunc
}

// ensureCallbacks lazily creates the callback registry and its cgo.Handle.
func (h *Handle) ensureCallbacks() *callbacks {
	if h.cb == nil {
		h.cb = &callbacks{}
		h.cbHandle = cgo.NewHandle(h.cb)
	}
	return h.cb
}

// releaseCallbacks deletes the cgo.Handle so the registry no longer retains the
// Go callbacks. It is called on Close.
func (h *Handle) releaseCallbacks() {
	if h.cbHandle != 0 {
		h.cbHandle.Delete()
		h.cbHandle = 0
	}
	h.cb = nil
}

// SetProgressCallback registers fn as the progress handler.
func (h *Handle) SetProgressCallback(fn ProgressFunc) {
	h.ensureCallbacks().progress = fn
	C.go_set_progress(h.ptr, C.uintptr_t(h.cbHandle))
}

// SetDataErrorCallback registers fn as the data-error handler.
func (h *Handle) SetDataErrorCallback(fn DataErrorFunc) {
	h.ensureCallbacks().dataError = fn
	C.go_set_dataerror(h.ptr, C.uintptr_t(h.cbHandle))
}

// SetExifCallback registers fn as the EXIF parser handler.
func (h *Handle) SetExifCallback(fn ExifParserFunc) {
	h.ensureCallbacks().exif = fn
	C.go_set_exif(h.ptr, C.uintptr_t(h.cbHandle))
}

// SetMakernotesCallback registers fn as the maker-notes handler.
func (h *Handle) SetMakernotesCallback(fn ExifParserFunc) {
	h.ensureCallbacks().makernotes = fn
	C.go_set_makernotes(h.ptr, C.uintptr_t(h.cbHandle))
}

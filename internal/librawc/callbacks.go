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
	// libraw_set_makernotes_handler was added in LibRaw 0.22; on older
	// libraries it is a no-op so the callback simply never fires.
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0, 22, 0)
	libraw_set_makernotes_handler(lr, (exif_parser_callback)goLibrawMakernotes, (void *)d);
#else
	(void)lr;
	(void)d;
	(void)goLibrawMakernotes;
#endif
}

// The clear helpers unregister a handler by installing a NULL callback so LibRaw
// stops invoking the Go trampoline entirely, rather than leaving it registered.
static void go_clear_progress(libraw_data_t *lr) {
	libraw_set_progress_handler(lr, NULL, NULL);
}
static void go_clear_dataerror(libraw_data_t *lr) {
	libraw_set_dataerror_handler(lr, NULL, NULL);
}
static void go_clear_exif(libraw_data_t *lr) {
	libraw_set_exifparser_handler(lr, NULL, NULL);
}
static void go_clear_makernotes(libraw_data_t *lr) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0, 22, 0)
	libraw_set_makernotes_handler(lr, NULL, NULL);
#else
	(void)lr;
#endif
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

// reregisterCallbacks reinstalls all previously registered callbacks on the C
// side. It must be called after libraw_recycle, which zeroes all C-side state.
func (h *Handle) reregisterCallbacks() {
	if h.cb == nil || h.cbHandle == 0 {
		return
	}
	if h.cb.progress != nil {
		C.go_set_progress(h.ptr, C.uintptr_t(h.cbHandle))
	}
	if h.cb.dataError != nil {
		C.go_set_dataerror(h.ptr, C.uintptr_t(h.cbHandle))
	}
	if h.cb.exif != nil {
		C.go_set_exif(h.ptr, C.uintptr_t(h.cbHandle))
	}
	if h.cb.makernotes != nil {
		C.go_set_makernotes(h.ptr, C.uintptr_t(h.cbHandle))
	}
}

// SetProgressCallback registers fn as the progress handler. A nil fn unregisters
// the handler on the LibRaw side so the trampoline is no longer invoked.
func (h *Handle) SetProgressCallback(fn ProgressFunc) {
	if fn == nil {
		if h.cb != nil {
			h.cb.progress = nil
		}
		C.go_clear_progress(h.ptr)
		return
	}
	h.ensureCallbacks().progress = fn
	C.go_set_progress(h.ptr, C.uintptr_t(h.cbHandle))
}

// SetDataErrorCallback registers fn as the data-error handler. A nil fn
// unregisters the handler on the LibRaw side.
func (h *Handle) SetDataErrorCallback(fn DataErrorFunc) {
	if fn == nil {
		if h.cb != nil {
			h.cb.dataError = nil
		}
		C.go_clear_dataerror(h.ptr)
		return
	}
	h.ensureCallbacks().dataError = fn
	C.go_set_dataerror(h.ptr, C.uintptr_t(h.cbHandle))
}

// SetExifCallback registers fn as the EXIF parser handler. A nil fn unregisters
// the handler on the LibRaw side.
func (h *Handle) SetExifCallback(fn ExifParserFunc) {
	if fn == nil {
		if h.cb != nil {
			h.cb.exif = nil
		}
		C.go_clear_exif(h.ptr)
		return
	}
	h.ensureCallbacks().exif = fn
	C.go_set_exif(h.ptr, C.uintptr_t(h.cbHandle))
}

// SetMakernotesCallback registers fn as the maker-notes handler. A nil fn
// unregisters the handler on the LibRaw side.
func (h *Handle) SetMakernotesCallback(fn ExifParserFunc) {
	if fn == nil {
		if h.cb != nil {
			h.cb.makernotes = nil
		}
		C.go_clear_makernotes(h.ptr)
		return
	}
	h.ensureCallbacks().makernotes = fn
	C.go_set_makernotes(h.ptr, C.uintptr_t(h.cbHandle))
}

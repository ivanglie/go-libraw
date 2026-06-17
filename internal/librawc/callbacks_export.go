//go:build cgo

package librawc

import "C"

import (
	"runtime/cgo"
	"unsafe"
)

// lookupCallbacks recovers the *callbacks registry from a LibRaw data pointer.
func lookupCallbacks(data unsafe.Pointer) *callbacks {
	if data == nil {
		return nil
	}
	cb, _ := cgo.Handle(uintptr(data)).Value().(*callbacks)
	return cb
}

// recoverCallback swallows a panic escaping a Go callback so it never crosses
// the C boundary.
func recoverCallback() { _ = recover() }

//export goLibrawProgress
func goLibrawProgress(data unsafe.Pointer, stage, iteration, expected C.int) (ret C.int) {
	// A panic in a progress callback is treated as cancellation so the failure
	// surfaces as a LibRaw error rather than crossing into C.
	defer func() {
		if r := recover(); r != nil {
			ret = 1
		}
	}()
	cb := lookupCallbacks(data)
	if cb == nil || cb.progress == nil {
		return 0
	}
	return C.int(cb.progress(int(stage), int(iteration), int(expected)))
}

//export goLibrawDataError
func goLibrawDataError(data unsafe.Pointer, file *C.char, offset C.longlong) {
	defer recoverCallback()
	cb := lookupCallbacks(data)
	if cb == nil || cb.dataError == nil {
		return
	}
	cb.dataError(C.GoString(file), int64(offset))
}

//export goLibrawExif
func goLibrawExif(data unsafe.Pointer, tag, typ, length C.int, ord C.uint, ifp unsafe.Pointer, base C.longlong) {
	defer recoverCallback()
	cb := lookupCallbacks(data)
	if cb == nil || cb.exif == nil {
		return
	}
	cb.exif(int(tag), int(typ), int(length), uint32(ord), int64(base))
}

//export goLibrawMakernotes
func goLibrawMakernotes(data unsafe.Pointer, tag, typ, length C.int, ord C.uint, ifp unsafe.Pointer, base C.longlong) {
	defer recoverCallback()
	cb := lookupCallbacks(data)
	if cb == nil || cb.makernotes == nil {
		return
	}
	cb.makernotes(int(tag), int(typ), int(length), uint32(ord), int64(base))
}

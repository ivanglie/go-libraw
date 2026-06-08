//go:build cgo

package librawc

/*
#include <libraw/libraw.h>

// libraw_adjust_to_raw_inset_crop was added in LibRaw 0.22; on older libraries
// report it as unsupported via *supported.
static int go_adjust_inset_crop(libraw_data_t *lr, unsigned mask, float maxcrop, int *supported) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0, 22, 0)
	*supported = 1;
	return libraw_adjust_to_raw_inset_crop(lr, mask, maxcrop);
#else
	(void)lr;
	(void)mask;
	(void)maxcrop;
	*supported = 0;
	return 0;
#endif
}
*/
import "C"

import "unsafe"

// CameraList returns LibRaw's supported camera names as a Go slice of copies.
func CameraList() []string {
	count := int(C.libraw_cameraCount())
	if count <= 0 {
		return nil
	}
	arr := C.libraw_cameraList()
	if arr == nil {
		return nil
	}
	names := unsafe.Slice(arr, count)
	out := make([]string, 0, count)
	for _, name := range names {
		if name == nil {
			break
		}
		out = append(out, C.GoString(name))
	}
	return out
}

// CameraCount returns the number of cameras LibRaw supports.
func CameraCount() int {
	return int(C.libraw_cameraCount())
}

// Capabilities returns LibRaw's runtime capability flags.
func Capabilities() uint {
	return uint(C.libraw_capabilities())
}

// UnpackFunctionName returns the name of the unpack function for the open image.
func (h *Handle) UnpackFunctionName() string {
	return C.GoString(C.libraw_unpack_function_name(h.ptr))
}

// DecoderInfo returns the decoder name, decoder flags, and LibRaw status code.
func (h *Handle) DecoderInfo() (string, uint32, int) {
	var info C.libraw_decoder_info_t
	code := int(C.libraw_get_decoder_info(h.ptr, &info))
	return C.GoString(info.decoder_name), uint32(info.decoder_flags), code
}

// IWidth returns the output image width in pixels.
func (h *Handle) IWidth() int { return int(C.libraw_get_iwidth(h.ptr)) }

// IHeight returns the output image height in pixels.
func (h *Handle) IHeight() int { return int(C.libraw_get_iheight(h.ptr)) }

// CamMul returns the camera white-balance multiplier at index (0..3).
func (h *Handle) CamMul(index int) float32 {
	return float32(C.libraw_get_cam_mul(h.ptr, C.int(index)))
}

// PreMul returns the pre-multiplier at index (0..3).
func (h *Handle) PreMul(index int) float32 {
	return float32(C.libraw_get_pre_mul(h.ptr, C.int(index)))
}

// RGBCam returns the rgb_cam color matrix element at (row, col).
func (h *Handle) RGBCam(row, col int) float32 {
	return float32(C.libraw_get_rgb_cam(h.ptr, C.int(row), C.int(col)))
}

// ColorMaximum returns the maximum color value for the open image.
func (h *Handle) ColorMaximum() int {
	return int(C.libraw_get_color_maximum(h.ptr))
}

// AdjustToRawInsetCrop adjusts the output crop to the raw inset. The second
// return reports whether the linked LibRaw supports the call (0.22+).
func (h *Handle) AdjustToRawInsetCrop(mask uint, maxCrop float32) (int, bool) {
	var supported C.int
	code := int(C.go_adjust_inset_crop(h.ptr, C.uint(mask), C.float(maxCrop), &supported))
	return code, supported != 0
}

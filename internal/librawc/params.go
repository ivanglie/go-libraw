//go:build cgo

package librawc

/*
#include <stdlib.h>
#include <libraw/libraw.h>

static int go_libraw_get_use_p1_correction(libraw_output_params_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->use_p1_correction;
#else
	(void)p;
	return 0;
#endif
}

static void go_libraw_set_use_p1_correction(libraw_output_params_t *p, int v) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	p->use_p1_correction = v;
#else
	(void)p;
	(void)v;
#endif
}
*/
import "C"

import "unsafe"

// OutputParams mirrors libraw_output_params_t. The four profile/path fields are
// exposed as Go strings; when set, copies are retained on the handle and freed
// on Close or on the next SetOutputParams call.
type OutputParams struct {
	Greybox          [4]uint32
	Cropbox          [4]uint32
	Aber             [4]float64
	Gamm             [6]float64
	UserMul          [4]float32
	Bright           float32
	Threshold        float32
	HalfSize         int
	FourColorRGB     int
	Highlight        int
	UseAutoWB        int
	UseCameraWB      int
	UseCameraMatrix  int
	OutputColor      int
	OutputProfile    string
	CameraProfile    string
	BadPixels        string
	DarkFrame        string
	OutputBPS        int
	OutputTIFF       int
	OutputFlags      int
	UserFlip         int
	UserQual         int
	UserBlack        int
	UserCblack       [4]int
	UserSat          int
	MedPasses        int
	AutoBrightThr    float32
	AdjustMaximumThr float32
	NoAutoBright     int
	UseFujiRotate    int
	UseP1Correction  int
	GreenMatching    int
	DCBIterations    int
	DCBEnhanceFL     int
	FBDDNoiseRD      int
	ExpCorrec        int
	ExpShift         float32
	ExpPreser        float32
	NoAutoScale      int
	NoInterpolation  int
}

// RawUnpackParams mirrors libraw_raw_unpack_params_t.
//
// The upstream custom_camera_strings (char**) field is intentionally not
// mirrored; it is documented as unsupported.
type RawUnpackParams struct {
	UseRawSpeed             int
	UseDNGSDK               int
	Options                 uint32
	ShotSelect              uint32
	Specials                uint32
	MaxRawMemoryMB          uint32
	SonyARW2PosterizationTh int
	CoolscanNEFGamma        float32
	P4ShotOrder             string
}

// GetOutputParams reads the handle's libraw_output_params_t into Go.
func (h *Handle) GetOutputParams() OutputParams {
	p := &h.ptr.params
	out := OutputParams{
		Bright:           float32(p.bright),
		Threshold:        float32(p.threshold),
		HalfSize:         int(p.half_size),
		FourColorRGB:     int(p.four_color_rgb),
		Highlight:        int(p.highlight),
		UseAutoWB:        int(p.use_auto_wb),
		UseCameraWB:      int(p.use_camera_wb),
		UseCameraMatrix:  int(p.use_camera_matrix),
		OutputColor:      int(p.output_color),
		OutputProfile:    C.GoString(p.output_profile),
		CameraProfile:    C.GoString(p.camera_profile),
		BadPixels:        C.GoString(p.bad_pixels),
		DarkFrame:        C.GoString(p.dark_frame),
		OutputBPS:        int(p.output_bps),
		OutputTIFF:       int(p.output_tiff),
		OutputFlags:      int(p.output_flags),
		UserFlip:         int(p.user_flip),
		UserQual:         int(p.user_qual),
		UserBlack:        int(p.user_black),
		UserSat:          int(p.user_sat),
		MedPasses:        int(p.med_passes),
		AutoBrightThr:    float32(p.auto_bright_thr),
		AdjustMaximumThr: float32(p.adjust_maximum_thr),
		NoAutoBright:     int(p.no_auto_bright),
		UseFujiRotate:    int(p.use_fuji_rotate),
		UseP1Correction:  int(C.go_libraw_get_use_p1_correction(p)),
		GreenMatching:    int(p.green_matching),
		DCBIterations:    int(p.dcb_iterations),
		DCBEnhanceFL:     int(p.dcb_enhance_fl),
		FBDDNoiseRD:      int(p.fbdd_noiserd),
		ExpCorrec:        int(p.exp_correc),
		ExpShift:         float32(p.exp_shift),
		ExpPreser:        float32(p.exp_preser),
		NoAutoScale:      int(p.no_auto_scale),
		NoInterpolation:  int(p.no_interpolation),
	}
	for i := 0; i < 4; i++ {
		out.Greybox[i] = uint32(p.greybox[i])
		out.Cropbox[i] = uint32(p.cropbox[i])
		out.Aber[i] = float64(p.aber[i])
		out.UserMul[i] = float32(p.user_mul[i])
		out.UserCblack[i] = int(p.user_cblack[i])
	}
	for i := 0; i < 6; i++ {
		out.Gamm[i] = float64(p.gamm[i])
	}
	return out
}

// SetOutputParams writes a full libraw_output_params_t from Go, retaining C
// copies of the string fields on the handle.
func (h *Handle) SetOutputParams(in OutputParams) {
	p := &h.ptr.params
	p.bright = C.float(in.Bright)
	p.threshold = C.float(in.Threshold)
	p.half_size = C.int(in.HalfSize)
	p.four_color_rgb = C.int(in.FourColorRGB)
	p.highlight = C.int(in.Highlight)
	p.use_auto_wb = C.int(in.UseAutoWB)
	p.use_camera_wb = C.int(in.UseCameraWB)
	p.use_camera_matrix = C.int(in.UseCameraMatrix)
	p.output_color = C.int(in.OutputColor)
	p.output_bps = C.int(in.OutputBPS)
	p.output_tiff = C.int(in.OutputTIFF)
	p.output_flags = C.int(in.OutputFlags)
	p.user_flip = C.int(in.UserFlip)
	p.user_qual = C.int(in.UserQual)
	p.user_black = C.int(in.UserBlack)
	p.user_sat = C.int(in.UserSat)
	p.med_passes = C.int(in.MedPasses)
	p.auto_bright_thr = C.float(in.AutoBrightThr)
	p.adjust_maximum_thr = C.float(in.AdjustMaximumThr)
	p.no_auto_bright = C.int(in.NoAutoBright)
	p.use_fuji_rotate = C.int(in.UseFujiRotate)
	C.go_libraw_set_use_p1_correction(p, C.int(in.UseP1Correction))
	p.green_matching = C.int(in.GreenMatching)
	p.dcb_iterations = C.int(in.DCBIterations)
	p.dcb_enhance_fl = C.int(in.DCBEnhanceFL)
	p.fbdd_noiserd = C.int(in.FBDDNoiseRD)
	p.exp_correc = C.int(in.ExpCorrec)
	p.exp_shift = C.float(in.ExpShift)
	p.exp_preser = C.float(in.ExpPreser)
	p.no_auto_scale = C.int(in.NoAutoScale)
	p.no_interpolation = C.int(in.NoInterpolation)
	for i := 0; i < 4; i++ {
		p.greybox[i] = C.uint(in.Greybox[i])
		p.cropbox[i] = C.uint(in.Cropbox[i])
		p.aber[i] = C.double(in.Aber[i])
		p.user_mul[i] = C.float(in.UserMul[i])
		p.user_cblack[i] = C.int(in.UserCblack[i])
	}
	for i := 0; i < 6; i++ {
		p.gamm[i] = C.double(in.Gamm[i])
	}

	h.freeParamStrings()
	p.output_profile = h.retainParamString(0, in.OutputProfile)
	p.camera_profile = h.retainParamString(1, in.CameraProfile)
	p.bad_pixels = h.retainParamString(2, in.BadPixels)
	p.dark_frame = h.retainParamString(3, in.DarkFrame)
}

// retainParamString stores a C copy of s at slot i and returns its pointer, or
// nil for an empty string.
func (h *Handle) retainParamString(i int, s string) *C.char {
	if s == "" {
		return nil
	}
	c := C.CString(s)
	h.paramStrings[i] = unsafe.Pointer(c)
	return c
}

// freeParamStrings releases all retained output-params string copies.
func (h *Handle) freeParamStrings() {
	for i := range h.paramStrings {
		if h.paramStrings[i] != nil {
			C.free(h.paramStrings[i])
			h.paramStrings[i] = nil
		}
	}
}

// GetRawUnpackParams reads the handle's libraw_raw_unpack_params_t into Go.
func (h *Handle) GetRawUnpackParams() RawUnpackParams {
	r := &h.ptr.rawparams
	return RawUnpackParams{
		UseRawSpeed:             int(r.use_rawspeed),
		UseDNGSDK:               int(r.use_dngsdk),
		Options:                 uint32(r.options),
		ShotSelect:              uint32(r.shot_select),
		Specials:                uint32(r.specials),
		MaxRawMemoryMB:          uint32(r.max_raw_memory_mb),
		SonyARW2PosterizationTh: int(r.sony_arw2_posterization_thr),
		CoolscanNEFGamma:        float32(r.coolscan_nef_gamma),
		P4ShotOrder:             C.GoString(&r.p4shot_order[0]),
	}
}

// SetRawUnpackParams writes a libraw_raw_unpack_params_t from Go. The
// custom_camera_strings field is left untouched (unsupported).
func (h *Handle) SetRawUnpackParams(in RawUnpackParams) {
	r := &h.ptr.rawparams
	r.use_rawspeed = C.int(in.UseRawSpeed)
	r.use_dngsdk = C.int(in.UseDNGSDK)
	r.options = C.uint(in.Options)
	r.shot_select = C.uint(in.ShotSelect)
	r.specials = C.uint(in.Specials)
	r.max_raw_memory_mb = C.uint(in.MaxRawMemoryMB)
	r.sony_arw2_posterization_thr = C.int(in.SonyARW2PosterizationTh)
	r.coolscan_nef_gamma = C.float(in.CoolscanNEFGamma)

	// p4shot_order is a fixed char[5]; copy up to 4 bytes plus a NUL.
	for i := range r.p4shot_order {
		r.p4shot_order[i] = 0
	}
	b := []byte(in.P4ShotOrder)
	if len(b) > len(r.p4shot_order)-1 {
		b = b[:len(r.p4shot_order)-1]
	}
	for i, c := range b {
		r.p4shot_order[i] = C.char(c)
	}
}

// Ergonomic setters mapping to the libraw_set_* helpers.

// SetDemosaic selects the demosaic algorithm (output_params.user_qual).
func (h *Handle) SetDemosaic(v int) { C.libraw_set_demosaic(h.ptr, C.int(v)) }

// SetOutputColor selects the output color space.
func (h *Handle) SetOutputColor(v int) { C.libraw_set_output_color(h.ptr, C.int(v)) }

// SetAdjustMaximumThr sets the maximum-value adjustment threshold.
func (h *Handle) SetAdjustMaximumThr(v float32) {
	C.libraw_set_adjust_maximum_thr(h.ptr, C.float(v))
}

// SetUserMul sets one user white-balance multiplier.
func (h *Handle) SetUserMul(index int, v float32) {
	C.libraw_set_user_mul(h.ptr, C.int(index), C.float(v))
}

// SetOutputBPS sets the output bit depth (8 or 16).
func (h *Handle) SetOutputBPS(v int) { C.libraw_set_output_bps(h.ptr, C.int(v)) }

// SetGamma sets one gamma curve coefficient.
func (h *Handle) SetGamma(index int, v float32) {
	C.libraw_set_gamma(h.ptr, C.int(index), C.float(v))
}

// SetNoAutoBright toggles auto-brightening.
func (h *Handle) SetNoAutoBright(v int) { C.libraw_set_no_auto_bright(h.ptr, C.int(v)) }

// SetBright sets the brightness multiplier.
func (h *Handle) SetBright(v float32) { C.libraw_set_bright(h.ptr, C.float(v)) }

// SetHighlight sets the highlight-recovery mode.
func (h *Handle) SetHighlight(v int) { C.libraw_set_highlight(h.ptr, C.int(v)) }

// SetFBDDNoiseRD sets the FBDD noise-reduction mode.
func (h *Handle) SetFBDDNoiseRD(v int) { C.libraw_set_fbdd_noiserd(h.ptr, C.int(v)) }

// SetOutputTIFF toggles TIFF output (vs PPM).
func (h *Handle) SetOutputTIFF(v int) { C.libraw_set_output_tif(h.ptr, C.int(v)) }

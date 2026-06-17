//go:build !cgo

// Package librawc contains the internal cgo bridge to LibRaw.
package librawc

import "errors"

// ErrCGODisabled reports that cgo is required for LibRaw bindings.
var ErrCGODisabled = errors.New("libraw: cgo is disabled; enable cgo and install LibRaw development headers")

// errorCodeCGODisabled is a non-zero LibRaw status returned by stub open paths so
// callers surface a typed error. These paths are unreachable in practice because
// New already fails when cgo is disabled.
const errorCodeCGODisabled = -1

// Handle is an unavailable LibRaw handle when cgo is disabled.
type Handle struct{}

// New reports that cgo is required for LibRaw.
func New(uint) (*Handle, error) {
	return nil, ErrCGODisabled
}

// Close releases the unavailable handle.
func (h *Handle) Close() {}

// OpenFile reports that cgo is required for LibRaw.
func (h *Handle) OpenFile(string) int {
	return int(errorCodeCGODisabled)
}

// OpenBuffer reports that cgo is required for LibRaw.
func (h *Handle) OpenBuffer([]byte) int {
	return int(errorCodeCGODisabled)
}

// BayerParams holds the geometry and decoding flags for OpenBayer.
type BayerParams struct {
	RawWidth     uint16
	RawHeight    uint16
	LeftMargin   uint16
	TopMargin    uint16
	RightMargin  uint16
	BottomMargin uint16
	ProcFlags    uint8
	BayerPattern uint8
	UnusedBits   uint
	OtherFlags   uint
	BlackLevel   uint
}

// OpenBayer reports that cgo is required for LibRaw.
func (h *Handle) OpenBayer([]byte, BayerParams) int {
	return int(errorCodeCGODisabled)
}

// Recycle is a no-op when cgo is disabled.
func (h *Handle) Recycle() {}

// RecycleDatastream is a no-op when cgo is disabled.
func (h *Handle) RecycleDatastream() {}

// Unpack reports that cgo is required for LibRaw.
func (h *Handle) Unpack() int { return int(errorCodeCGODisabled) }

// UnpackThumb reports that cgo is required for LibRaw.
func (h *Handle) UnpackThumb() int { return int(errorCodeCGODisabled) }

// UnpackThumbEx reports that cgo is required for LibRaw.
func (h *Handle) UnpackThumbEx(int) int { return int(errorCodeCGODisabled) }

// SubtractBlack is a no-op when cgo is disabled.
func (h *Handle) SubtractBlack() {}

// Raw2Image reports that cgo is required for LibRaw.
func (h *Handle) Raw2Image() int { return int(errorCodeCGODisabled) }

// FreeImage is a no-op when cgo is disabled.
func (h *Handle) FreeImage() {}

// AdjustSizesInfoOnly reports that cgo is required for LibRaw.
func (h *Handle) AdjustSizesInfoOnly() int { return int(errorCodeCGODisabled) }

// DcrawProcess reports that cgo is required for LibRaw.
func (h *Handle) DcrawProcess() int { return int(errorCodeCGODisabled) }

// DcrawPPMTiffWriter reports that cgo is required for LibRaw.
func (h *Handle) DcrawPPMTiffWriter(string) int { return int(errorCodeCGODisabled) }

// DcrawThumbWriter reports that cgo is required for LibRaw.
func (h *Handle) DcrawThumbWriter(string) int { return int(errorCodeCGODisabled) }

// ProcessedImage is a Go copy of a libraw_processed_image_t.
type ProcessedImage struct {
	Type   int
	Height uint16
	Width  uint16
	Colors uint16
	Bits   uint16
	Data   []byte
}

// OutputParams mirrors libraw_output_params_t when cgo is available.
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

// RawUnpackParams mirrors libraw_raw_unpack_params_t when cgo is available.
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

// GetOutputParams reports zero values when cgo is disabled.
func (h *Handle) GetOutputParams() OutputParams { return OutputParams{} }

// SetOutputParams is a no-op when cgo is disabled.
func (h *Handle) SetOutputParams(OutputParams) {}

// GetRawUnpackParams reports zero values when cgo is disabled.
func (h *Handle) GetRawUnpackParams() RawUnpackParams { return RawUnpackParams{} }

// SetRawUnpackParams is a no-op when cgo is disabled.
func (h *Handle) SetRawUnpackParams(RawUnpackParams) {}

// SetDemosaic is a no-op when cgo is disabled.
func (h *Handle) SetDemosaic(int) {}

// SetOutputColor is a no-op when cgo is disabled.
func (h *Handle) SetOutputColor(int) {}

// SetAdjustMaximumThr is a no-op when cgo is disabled.
func (h *Handle) SetAdjustMaximumThr(float32) {}

// SetUserMul is a no-op when cgo is disabled.
func (h *Handle) SetUserMul(int, float32) {}

// SetOutputBPS is a no-op when cgo is disabled.
func (h *Handle) SetOutputBPS(int) {}

// SetGamma is a no-op when cgo is disabled.
func (h *Handle) SetGamma(int, float32) {}

// SetNoAutoBright is a no-op when cgo is disabled.
func (h *Handle) SetNoAutoBright(int) {}

// SetBright is a no-op when cgo is disabled.
func (h *Handle) SetBright(float32) {}

// SetHighlight is a no-op when cgo is disabled.
func (h *Handle) SetHighlight(int) {}

// SetFBDDNoiseRD is a no-op when cgo is disabled.
func (h *Handle) SetFBDDNoiseRD(int) {}

// SetOutputTIFF is a no-op when cgo is disabled.
func (h *Handle) SetOutputTIFF(int) {}

// MakeMemImage reports that cgo is required for LibRaw.
func (h *Handle) MakeMemImage() (ProcessedImage, int) {
	return ProcessedImage{}, int(errorCodeCGODisabled)
}

// MakeMemThumb reports that cgo is required for LibRaw.
func (h *Handle) MakeMemThumb() (ProcessedImage, int) {
	return ProcessedImage{}, int(errorCodeCGODisabled)
}

// Version returns an empty version when cgo is disabled.
func Version() string {
	return ""
}

// VersionNumber returns zero when cgo is disabled.
func VersionNumber() int {
	return 0
}

// StrError returns an empty string when cgo is disabled.
func StrError(int) string {
	return ""
}

// StrProgress returns an empty string when cgo is disabled.
func StrProgress(int) string {
	return ""
}

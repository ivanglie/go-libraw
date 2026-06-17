package libraw

import (
	"fmt"

	"github.com/ivanglie/go-libraw/internal/librawc"
)

const (
	userMulLen     = 4
	gammaLen       = 6
	p4ShotOrderLen = 4
)

// OutputParams mirrors LibRaw's libraw_output_params_t.
//
// The string fields name files or profiles consumed by LibRaw. SetOutputParams
// copies those strings into memory retained by the Processor.
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

// RawUnpackParams mirrors LibRaw's libraw_raw_unpack_params_t.
//
// LibRaw's custom_camera_strings field is not exposed because it is a
// null-terminated char** list with ownership semantics that need a dedicated API.
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

// WithOutputParams applies output parameters when the Processor is created.
func WithOutputParams(params OutputParams) Option {
	return func(o *options) {
		o.outputParams = &params
	}
}

// WithRawUnpackParams applies raw unpack parameters when the Processor is created.
func WithRawUnpackParams(params RawUnpackParams) Option {
	return func(o *options) {
		o.rawUnpackParams = &params
	}
}

// OutputParams returns a copy of the current LibRaw output parameters.
func (p *Processor) OutputParams() (OutputParams, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return OutputParams{}, ErrClosed
	}
	return OutputParams(p.handle.GetOutputParams()), nil
}

// SetOutputParams replaces the current LibRaw output parameters.
func (p *Processor) SetOutputParams(params OutputParams) error {
	return p.withHandleVoid(func(h *librawc.Handle) {
		h.SetOutputParams(librawc.OutputParams(params))
	})
}

// RawUnpackParams returns a copy of the current LibRaw raw unpack parameters.
func (p *Processor) RawUnpackParams() (RawUnpackParams, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return RawUnpackParams{}, ErrClosed
	}
	return RawUnpackParams(p.handle.GetRawUnpackParams()), nil
}

// SetRawUnpackParams replaces the current LibRaw raw unpack parameters.
func (p *Processor) SetRawUnpackParams(params RawUnpackParams) error {
	if len([]byte(params.P4ShotOrder)) > p4ShotOrderLen {
		return fmt.Errorf("libraw: P4ShotOrder length %d exceeds %d bytes", len([]byte(params.P4ShotOrder)), p4ShotOrderLen)
	}
	return p.withHandleVoid(func(h *librawc.Handle) {
		h.SetRawUnpackParams(librawc.RawUnpackParams(params))
	})
}

// SetDemosaic selects the demosaic algorithm (output_params.user_qual).
func (p *Processor) SetDemosaic(value int) error {
	if value < 0 {
		return fmt.Errorf("libraw: demosaic value %d is negative", value)
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetDemosaic(value) })
}

// SetOutputColor selects the output color space.
func (p *Processor) SetOutputColor(value int) error {
	if value < 0 {
		return fmt.Errorf("libraw: output color value %d is negative", value)
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetOutputColor(value) })
}

// SetAdjustMaximumThreshold sets the maximum-value adjustment threshold.
func (p *Processor) SetAdjustMaximumThreshold(value float32) error {
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetAdjustMaximumThr(value) })
}

// SetUserMul sets one user white-balance multiplier.
func (p *Processor) SetUserMul(index int, value float32) error {
	if err := validateIndex("user multiplier", index, userMulLen); err != nil {
		return err
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetUserMul(index, value) })
}

// SetOutputBPS sets the output bit depth. LibRaw accepts 8 or 16 bits.
func (p *Processor) SetOutputBPS(bits int) error {
	if bits != 8 && bits != 16 {
		return fmt.Errorf("libraw: output bits per sample %d is invalid; want 8 or 16", bits)
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetOutputBPS(bits) })
}

// SetGamma sets one gamma curve coefficient.
func (p *Processor) SetGamma(index int, value float32) error {
	if err := validateIndex("gamma", index, gammaLen); err != nil {
		return err
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetGamma(index, value) })
}

// SetNoAutoBright toggles auto-brightening off.
func (p *Processor) SetNoAutoBright(disable bool) error {
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetNoAutoBright(boolInt(disable)) })
}

// SetBright sets the brightness multiplier.
func (p *Processor) SetBright(value float32) error {
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetBright(value) })
}

// SetHighlight sets the highlight-recovery mode.
func (p *Processor) SetHighlight(value int) error {
	if value < 0 {
		return fmt.Errorf("libraw: highlight value %d is negative", value)
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetHighlight(value) })
}

// SetFBDDNoiseReduction sets the FBDD noise-reduction mode.
func (p *Processor) SetFBDDNoiseReduction(value int) error {
	if value < 0 {
		return fmt.Errorf("libraw: FBDD noise reduction value %d is negative", value)
	}
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetFBDDNoiseRD(value) })
}

// SetOutputTIFF toggles TIFF output instead of PPM output.
func (p *Processor) SetOutputTIFF(enabled bool) error {
	return p.withHandleVoid(func(h *librawc.Handle) { h.SetOutputTIFF(boolInt(enabled)) })
}

func validateIndex(name string, index, length int) error {
	if index < 0 || index >= length {
		return fmt.Errorf("libraw: %s index %d out of range [0,%d)", name, index, length)
	}
	return nil
}

func boolInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

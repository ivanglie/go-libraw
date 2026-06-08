package libraw

import (
	"errors"
	"testing"
)

func TestOutputParamsRoundTrip(t *testing.T) {
	p := openProcessor(t)

	want := OutputParams{
		Greybox:          [4]uint32{1, 2, 3, 4},
		Cropbox:          [4]uint32{5, 6, 7, 8},
		Aber:             [4]float64{1.1, 1.2, 1.3, 1.4},
		Gamm:             [6]float64{2.2, 4.5, 0.1, 0.2, 0.3, 0.4},
		UserMul:          [4]float32{1.0, 1.1, 1.2, 1.3},
		Bright:           1.25,
		Threshold:        0.05,
		HalfSize:         1,
		FourColorRGB:     1,
		Highlight:        2,
		UseAutoWB:        1,
		UseCameraWB:      1,
		UseCameraMatrix:  1,
		OutputColor:      LIBRAW_COLORSPACE_sRGB,
		OutputProfile:    "out.icc",
		CameraProfile:    "camera.icc",
		BadPixels:        "badpixels.txt",
		DarkFrame:        "darkframe.pgm",
		OutputBPS:        16,
		OutputTIFF:       1,
		OutputFlags:      LIBRAW_OUTPUT_FLAGS_PPMMETA,
		UserFlip:         3,
		UserQual:         3,
		UserBlack:        128,
		UserCblack:       [4]int{1, 2, 3, 4},
		UserSat:          4095,
		MedPasses:        2,
		AutoBrightThr:    0.01,
		AdjustMaximumThr: 0.7,
		NoAutoBright:     1,
		UseFujiRotate:    1,
		GreenMatching:    1,
		DCBIterations:    2,
		DCBEnhanceFL:     1,
		FBDDNoiseRD:      1,
		ExpCorrec:        1,
		ExpShift:         0.5,
		ExpPreser:        0.75,
		NoAutoScale:      1,
		NoInterpolation:  1,
	}
	if VersionNumber() >= 0x1600 {
		want.UseP1Correction = 1
	}

	if err := p.SetOutputParams(want); err != nil {
		t.Fatalf("SetOutputParams() error = %v", err)
	}
	got, err := p.OutputParams()
	if err != nil {
		t.Fatalf("OutputParams() error = %v", err)
	}
	if got != want {
		t.Fatalf("OutputParams() = %+v, want %+v", got, want)
	}
}

func TestRawUnpackParamsRoundTrip(t *testing.T) {
	p := openProcessor(t)

	want := RawUnpackParams{
		UseRawSpeed:             1,
		UseDNGSDK:               1,
		Options:                 7,
		ShotSelect:              2,
		Specials:                3,
		MaxRawMemoryMB:          256,
		SonyARW2PosterizationTh: 100,
		CoolscanNEFGamma:        2.4,
		P4ShotOrder:             "0123",
	}
	if err := p.SetRawUnpackParams(want); err != nil {
		t.Fatalf("SetRawUnpackParams() error = %v", err)
	}
	got, err := p.RawUnpackParams()
	if err != nil {
		t.Fatalf("RawUnpackParams() error = %v", err)
	}
	if got != want {
		t.Fatalf("RawUnpackParams() = %+v, want %+v", got, want)
	}
}

func TestParameterSetters(t *testing.T) {
	p := openProcessor(t)

	ops := []struct {
		name string
		fn   func() error
	}{
		{"SetDemosaic", func() error { return p.SetDemosaic(3) }},
		{"SetOutputColor", func() error { return p.SetOutputColor(LIBRAW_COLORSPACE_sRGB) }},
		{"SetAdjustMaximumThreshold", func() error { return p.SetAdjustMaximumThreshold(0.6) }},
		{"SetUserMul", func() error { return p.SetUserMul(2, 1.5) }},
		{"SetOutputBPS", func() error { return p.SetOutputBPS(16) }},
		{"SetGamma", func() error { return p.SetGamma(1, 4.5) }},
		{"SetNoAutoBright", func() error { return p.SetNoAutoBright(true) }},
		{"SetBright", func() error { return p.SetBright(1.3) }},
		{"SetHighlight", func() error { return p.SetHighlight(1) }},
		{"SetFBDDNoiseReduction", func() error { return p.SetFBDDNoiseReduction(1) }},
		{"SetOutputTIFF", func() error { return p.SetOutputTIFF(true) }},
	}
	for _, op := range ops {
		t.Run(op.name, func(t *testing.T) {
			if err := op.fn(); err != nil {
				t.Fatalf("%s() error = %v", op.name, err)
			}
		})
	}

	params, err := p.OutputParams()
	if err != nil {
		t.Fatalf("OutputParams() error = %v", err)
	}
	if params.UserQual != 3 {
		t.Fatalf("UserQual = %d, want 3", params.UserQual)
	}
	if params.OutputColor != LIBRAW_COLORSPACE_sRGB {
		t.Fatalf("OutputColor = %d, want %d", params.OutputColor, LIBRAW_COLORSPACE_sRGB)
	}
	if params.OutputBPS != 16 {
		t.Fatalf("OutputBPS = %d, want 16", params.OutputBPS)
	}
	if params.Gamm[1] != 4.5 {
		t.Fatalf("Gamm[1] = %v, want 4.5", params.Gamm[1])
	}
	if params.UserMul[2] != 1.5 {
		t.Fatalf("UserMul[2] = %v, want 1.5", params.UserMul[2])
	}
	if params.NoAutoBright != 1 || params.OutputTIFF != 1 {
		t.Fatalf("NoAutoBright/OutputTIFF = %d/%d, want 1/1", params.NoAutoBright, params.OutputTIFF)
	}
}

func TestParameterSetterValidation(t *testing.T) {
	p := openProcessor(t)

	tests := map[string]func() error{
		"user mul negative":     func() error { return p.SetUserMul(-1, 1) },
		"user mul high":         func() error { return p.SetUserMul(4, 1) },
		"gamma negative":        func() error { return p.SetGamma(-1, 1) },
		"gamma high":            func() error { return p.SetGamma(6, 1) },
		"invalid bps":           func() error { return p.SetOutputBPS(12) },
		"negative demosaic":     func() error { return p.SetDemosaic(-1) },
		"negative output color": func() error { return p.SetOutputColor(-1) },
		"negative highlight":    func() error { return p.SetHighlight(-1) },
		"negative fbdd noiserd": func() error { return p.SetFBDDNoiseReduction(-1) },
		"long p4 shot order":    func() error { return p.SetRawUnpackParams(RawUnpackParams{P4ShotOrder: "01234"}) },
	}
	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			if err := fn(); err == nil {
				t.Fatal("setter returned nil error")
			}
		})
	}
}

func TestNewProcessorWithParameterOptions(t *testing.T) {
	p, err := NewProcessor(
		WithOutputParams(OutputParams{OutputBPS: 16, Bright: 1.1}),
		WithRawUnpackParams(RawUnpackParams{P4ShotOrder: "3210"}),
	)
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	t.Cleanup(func() { _ = p.Close() })

	out, err := p.OutputParams()
	if err != nil {
		t.Fatalf("OutputParams() error = %v", err)
	}
	if out.OutputBPS != 16 || out.Bright != 1.1 {
		t.Fatalf("output params = %+v, want OutputBPS=16 Bright=1.1", out)
	}
	raw, err := p.RawUnpackParams()
	if err != nil {
		t.Fatalf("RawUnpackParams() error = %v", err)
	}
	if raw.P4ShotOrder != "3210" {
		t.Fatalf("P4ShotOrder = %q, want %q", raw.P4ShotOrder, "3210")
	}
}

func TestNewProcessorRejectsInvalidRawUnpackOption(t *testing.T) {
	_, err := NewProcessor(WithRawUnpackParams(RawUnpackParams{P4ShotOrder: "01234"}))
	if err == nil {
		t.Fatal("NewProcessor() returned nil error")
	}
}

func TestParameterApplicationDuringProcessing(t *testing.T) {
	p := openProcessor(t)

	if err := p.SetOutputColor(LIBRAW_COLORSPACE_sRGB); err != nil {
		t.Fatalf("SetOutputColor() error = %v", err)
	}
	if err := p.SetOutputBPS(8); err != nil {
		t.Fatalf("SetOutputBPS() error = %v", err)
	}
	if err := p.SetGamma(0, 2.2); err != nil {
		t.Fatalf("SetGamma(0) error = %v", err)
	}
	if err := p.SetGamma(1, 4.5); err != nil {
		t.Fatalf("SetGamma(1) error = %v", err)
	}
	if err := p.SetBright(1.2); err != nil {
		t.Fatalf("SetBright() error = %v", err)
	}
	if err := p.SetDemosaic(3); err != nil {
		t.Fatalf("SetDemosaic() error = %v", err)
	}
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
	if err := p.DcrawProcess(); err != nil {
		t.Fatalf("DcrawProcess() error = %v", err)
	}
	img, err := p.MemImage()
	if err != nil {
		t.Fatalf("MemImage() error = %v", err)
	}
	if img.Bits != 8 {
		t.Fatalf("MemImage().Bits = %d, want 8", img.Bits)
	}
}

func TestParamsAfterCloseReturnErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	ops := map[string]func() error{
		"OutputParams": func() error {
			_, err := p.OutputParams()
			return err
		},
		"SetOutputParams":           func() error { return p.SetOutputParams(OutputParams{}) },
		"RawUnpackParams":           func() error { _, err := p.RawUnpackParams(); return err },
		"SetRawUnpackParams":        func() error { return p.SetRawUnpackParams(RawUnpackParams{}) },
		"SetDemosaic":               func() error { return p.SetDemosaic(0) },
		"SetOutputColor":            func() error { return p.SetOutputColor(0) },
		"SetAdjustMaximumThreshold": func() error { return p.SetAdjustMaximumThreshold(0) },
		"SetUserMul":                func() error { return p.SetUserMul(0, 1) },
		"SetOutputBPS":              func() error { return p.SetOutputBPS(8) },
		"SetGamma":                  func() error { return p.SetGamma(0, 1) },
		"SetNoAutoBright":           func() error { return p.SetNoAutoBright(true) },
		"SetBright":                 func() error { return p.SetBright(1) },
		"SetHighlight":              func() error { return p.SetHighlight(0) },
		"SetFBDDNoiseReduction":     func() error { return p.SetFBDDNoiseReduction(0) },
		"SetOutputTIFF":             func() error { return p.SetOutputTIFF(true) },
	}
	for name, op := range ops {
		t.Run(name, func(t *testing.T) {
			if err := op(); !errors.Is(err, ErrClosed) {
				t.Fatalf("%s after Close() error = %v, want ErrClosed", name, err)
			}
		})
	}
}

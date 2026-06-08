package libraw

import (
	"errors"
	"testing"
)

func TestCameraList(t *testing.T) {
	count := CameraCount()
	if count <= 0 {
		t.Fatalf("CameraCount() = %d, want positive", count)
	}
	list := CameraList()
	if len(list) != count {
		t.Fatalf("len(CameraList()) = %d, want %d", len(list), count)
	}
	if list[0] == "" {
		t.Fatal("CameraList()[0] is empty")
	}
}

func TestCapabilities(t *testing.T) {
	// Capabilities reflects build-time flags; just confirm the call is wired.
	_ = Capabilities()
}

func TestUnpackFunctionNameAndDecoderInfo(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	name, err := p.UnpackFunctionName()
	if err != nil {
		t.Fatalf("UnpackFunctionName() error = %v", err)
	}
	if name == "" {
		t.Fatal("UnpackFunctionName() is empty")
	}

	dec, err := p.DecoderInfo()
	if err != nil {
		t.Fatalf("DecoderInfo() error = %v", err)
	}
	if dec.Name == "" {
		t.Fatal("DecoderInfo().Name is empty")
	}
}

func TestColorAndDimensionHelpers(t *testing.T) {
	p := developed(t) // open + unpack + dcraw process

	w, err := p.IWidth()
	if err != nil {
		t.Fatalf("IWidth() error = %v", err)
	}
	h, err := p.IHeight()
	if err != nil {
		t.Fatalf("IHeight() error = %v", err)
	}
	if w <= 0 || h <= 0 {
		t.Fatalf("output dimensions = %dx%d, want positive", w, h)
	}

	max, err := p.ColorMaximum()
	if err != nil {
		t.Fatalf("ColorMaximum() error = %v", err)
	}
	if max <= 0 {
		t.Fatalf("ColorMaximum() = %d, want positive", max)
	}

	for i := 0; i < camMulLen; i++ {
		if _, err := p.CamMul(i); err != nil {
			t.Fatalf("CamMul(%d) error = %v", i, err)
		}
		if _, err := p.PreMul(i); err != nil {
			t.Fatalf("PreMul(%d) error = %v", i, err)
		}
	}
	if _, err := p.RGBCam(0, 0); err != nil {
		t.Fatalf("RGBCam(0,0) error = %v", err)
	}
}

func TestColorHelperValidation(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	tests := []struct {
		name string
		fn   func() error
	}{
		{"CamMul(-1)", func() error { _, e := p.CamMul(-1); return e }},
		{"CamMul(4)", func() error { _, e := p.CamMul(4); return e }},
		{"PreMul(4)", func() error { _, e := p.PreMul(4); return e }},
		{"RGBCam(3,0)", func() error { _, e := p.RGBCam(3, 0); return e }},
		{"RGBCam(0,4)", func() error { _, e := p.RGBCam(0, 4); return e }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err == nil {
				t.Fatalf("%s returned nil error for invalid index", tt.name)
			}
		})
	}
}

func TestAdjustToRawInsetCrop(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	err := p.AdjustToRawInsetCrop(1, 0.5)
	// libraw_adjust_to_raw_inset_crop exists from LibRaw 0.22 (version number
	// 0x001600). Older libraries report ErrUnsupported.
	if VersionNumber() >= 0x001600 {
		if errors.Is(err, ErrUnsupported) {
			t.Fatalf("AdjustToRawInsetCrop() = ErrUnsupported on LibRaw %s", Version())
		}
	} else if !errors.Is(err, ErrUnsupported) {
		t.Fatalf("AdjustToRawInsetCrop() error = %v, want ErrUnsupported on LibRaw %s", err, Version())
	}
}

func TestHelpersAfterCloseReturnErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	ops := map[string]func() error{
		"UnpackFunctionName":   func() error { _, e := p.UnpackFunctionName(); return e },
		"DecoderInfo":          func() error { _, e := p.DecoderInfo(); return e },
		"IWidth":               func() error { _, e := p.IWidth(); return e },
		"IHeight":              func() error { _, e := p.IHeight(); return e },
		"CamMul":               func() error { _, e := p.CamMul(0); return e },
		"PreMul":               func() error { _, e := p.PreMul(0); return e },
		"RGBCam":               func() error { _, e := p.RGBCam(0, 0); return e },
		"ColorMaximum":         func() error { _, e := p.ColorMaximum(); return e },
		"AdjustToRawInsetCrop": func() error { return p.AdjustToRawInsetCrop(0, 0) },
	}
	for name, op := range ops {
		t.Run(name, func(t *testing.T) {
			if err := op(); !errors.Is(err, ErrClosed) {
				t.Fatalf("%s after Close() error = %v, want ErrClosed", name, err)
			}
		})
	}
}

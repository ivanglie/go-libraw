package libraw

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

const sampleRAW = "testdata/RAW_RICOH_GR2.DNG"

func openProcessor(t *testing.T) *Processor {
	t.Helper()
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	t.Cleanup(func() {
		_ = p.Close()
	})
	return p
}

func TestOpenFile(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile(%q) error = %v", sampleRAW, err)
	}
}

func TestOpenFileInvalidPath(t *testing.T) {
	p := openProcessor(t)
	err := p.OpenFile(filepath.Join(t.TempDir(), "does-not-exist.dng"))
	if err == nil {
		t.Fatal("OpenFile() with missing path returned nil error")
	}
	var le Error
	if !errors.As(err, &le) {
		t.Fatalf("OpenFile() error = %T (%v), want libraw.Error", err, err)
	}
}

func TestOpenBuffer(t *testing.T) {
	data, err := os.ReadFile(sampleRAW)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	p := openProcessor(t)
	if err := p.OpenBuffer(data); err != nil {
		t.Fatalf("OpenBuffer() error = %v", err)
	}
}

func TestOpenBufferInvalid(t *testing.T) {
	p := openProcessor(t)

	tests := map[string][]byte{
		"empty":   nil,
		"garbage": []byte("this is not a raw file"),
	}
	for name, data := range tests {
		t.Run(name, func(t *testing.T) {
			err := p.OpenBuffer(data)
			if err == nil {
				t.Fatal("OpenBuffer() with invalid data returned nil error")
			}
			var le Error
			if !errors.As(err, &le) {
				t.Fatalf("OpenBuffer() error = %T (%v), want libraw.Error", err, err)
			}
		})
	}
}

func TestOpenBayer(t *testing.T) {
	const w, h = 16, 16
	data := make([]byte, w*h*2) // 16-bit samples
	for i := range data {
		data[i] = byte(i)
	}

	p := openProcessor(t)
	params := BayerParams{
		RawWidth:     w,
		RawHeight:    h,
		BayerPattern: uint8(LIBRAW_OPENBAYER_RGGB),
	}
	if err := p.OpenBayer(data, params); err != nil {
		t.Fatalf("OpenBayer() error = %v", err)
	}
}

func TestRecycleClearsInput(t *testing.T) {
	p := openProcessor(t)

	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("first OpenFile() error = %v", err)
	}
	if err := p.Recycle(); err != nil {
		t.Fatalf("Recycle() error = %v", err)
	}
	// A recycled handle must be reusable for another input.
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() after Recycle() error = %v", err)
	}
}

func TestRecycleDatastream(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.RecycleDatastream(); err != nil {
		t.Fatalf("RecycleDatastream() error = %v", err)
	}
}

func TestOpenAfterCloseReturnsErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	ops := map[string]func() error{
		"OpenFile":          func() error { return p.OpenFile(sampleRAW) },
		"OpenBuffer":        func() error { return p.OpenBuffer([]byte{0}) },
		"OpenBayer":         func() error { return p.OpenBayer([]byte{0, 0}, BayerParams{RawWidth: 1, RawHeight: 1}) },
		"Recycle":           p.Recycle,
		"RecycleDatastream": p.RecycleDatastream,
	}
	for name, op := range ops {
		t.Run(name, func(t *testing.T) {
			if err := op(); !errors.Is(err, ErrClosed) {
				t.Fatalf("%s after Close() error = %v, want ErrClosed", name, err)
			}
		})
	}
}

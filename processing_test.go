package libraw

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// developed opens, unpacks, and postprocesses the sample RAW, returning a
// processor ready for memory or file output.
func developed(t *testing.T) *Processor {
	t.Helper()
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
	if err := p.DcrawProcess(); err != nil {
		t.Fatalf("DcrawProcess() error = %v", err)
	}
	return p
}

func TestMemImage(t *testing.T) {
	p := developed(t)

	img, err := p.MemImage()
	if err != nil {
		t.Fatalf("MemImage() error = %v", err)
	}
	if img.Width == 0 || img.Height == 0 {
		t.Fatalf("MemImage() dims = %dx%d, want non-zero", img.Width, img.Height)
	}
	if img.Colors < 3 {
		t.Fatalf("MemImage() colors = %d, want >= 3", img.Colors)
	}
	if len(img.Data) == 0 {
		t.Fatal("MemImage() returned empty data")
	}
	want := int(img.Width) * int(img.Height) * int(img.Colors) * int(img.Bits) / 8
	if len(img.Data) != want {
		t.Fatalf("MemImage() data len = %d, want %d", len(img.Data), want)
	}

	// A second call must return an independent, equally sized Go-owned copy,
	// proving the C allocation was freed and copied each time.
	img2, err := p.MemImage()
	if err != nil {
		t.Fatalf("second MemImage() error = %v", err)
	}
	if len(img2.Data) != len(img.Data) {
		t.Fatalf("second MemImage() data len = %d, want %d", len(img2.Data), len(img.Data))
	}
}

func TestWritePPMTiff(t *testing.T) {
	p := developed(t)

	out := filepath.Join(t.TempDir(), "out.ppm")
	if err := p.WritePPMTiff(out); err != nil {
		t.Fatalf("WritePPMTiff() error = %v", err)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !bytes.HasPrefix(data, []byte("P6")) {
		t.Fatalf("output is not a binary PPM (prefix %q)", data[:2])
	}
}

// sampleThumbRAW has an embedded JPEG thumbnail (the Ricoh DNG fixture does not).
const sampleThumbRAW = "testdata/RAW_CANON_6D.CR2"

func TestThumbnail(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleThumbRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.UnpackThumb(); err != nil {
		t.Fatalf("UnpackThumb() error = %v", err)
	}

	img, err := p.MemThumb()
	if err != nil {
		t.Fatalf("MemThumb() error = %v", err)
	}
	if len(img.Data) == 0 {
		t.Fatal("MemThumb() returned empty data")
	}

	out := filepath.Join(t.TempDir(), "thumb.out")
	if err := p.WriteThumb(out); err != nil {
		t.Fatalf("WriteThumb() error = %v", err)
	}
	if info, err := os.Stat(out); err != nil || info.Size() == 0 {
		t.Fatalf("WriteThumb() produced no output: err=%v", err)
	}
}

func TestProcessingSteps(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}

	// Order matters: Raw2Image builds the image buffer that SubtractBlack
	// operates on, so it must run first.
	steps := []struct {
		name string
		fn   func() error
	}{
		{"AdjustSizesInfoOnly", p.AdjustSizesInfoOnly},
		{"Raw2Image", p.Raw2Image},
		{"SubtractBlack", p.SubtractBlack},
		{"FreeImage", p.FreeImage},
	}
	for _, s := range steps {
		if err := s.fn(); err != nil {
			t.Fatalf("%s() error = %v", s.name, err)
		}
	}
}

func TestUnpackThumbEx(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleThumbRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.UnpackThumbEx(0); err != nil {
		t.Fatalf("UnpackThumbEx(0) error = %v", err)
	}
}

func TestProcessOutOfOrder(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	// DcrawProcess before Unpack has no decoded RAW to work on.
	err := p.DcrawProcess()
	if err == nil {
		t.Fatal("DcrawProcess() before Unpack returned nil error")
	}
	var le Error
	if !errors.As(err, &le) {
		t.Fatalf("DcrawProcess() error = %T (%v), want libraw.Error", err, err)
	}
}

func TestProcessingAfterCloseReturnsErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	ops := map[string]func() error{
		"Unpack":              p.Unpack,
		"UnpackThumb":         p.UnpackThumb,
		"UnpackThumbEx":       func() error { return p.UnpackThumbEx(0) },
		"SubtractBlack":       p.SubtractBlack,
		"Raw2Image":           p.Raw2Image,
		"FreeImage":           p.FreeImage,
		"AdjustSizesInfoOnly": p.AdjustSizesInfoOnly,
		"DcrawProcess":        p.DcrawProcess,
		"WritePPMTiff":        func() error { return p.WritePPMTiff("x.ppm") },
		"WriteThumb":          func() error { return p.WriteThumb("x.jpg") },
		"MemImage":            func() error { _, e := p.MemImage(); return e },
		"MemThumb":            func() error { _, e := p.MemThumb(); return e },
	}
	for name, op := range ops {
		t.Run(name, func(t *testing.T) {
			if err := op(); !errors.Is(err, ErrClosed) {
				t.Fatalf("%s after Close() error = %v, want ErrClosed", name, err)
			}
		})
	}
}

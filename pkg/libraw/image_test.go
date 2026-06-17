package libraw

import (
	"errors"
	"testing"
)

func TestColor(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	// Every position in the color filter array maps to an index in [0,3].
	for _, pos := range [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}} {
		c, err := p.Color(pos[0], pos[1])
		if err != nil {
			t.Fatalf("Color%v error = %v", pos, err)
		}
		if c < 0 || c > 3 {
			t.Fatalf("Color%v = %d, want 0..3", pos, c)
		}
	}
}

func TestRawDimensions(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	w, err := p.RawWidth()
	if err != nil {
		t.Fatalf("RawWidth() error = %v", err)
	}
	h, err := p.RawHeight()
	if err != nil {
		t.Fatalf("RawHeight() error = %v", err)
	}
	if w <= 0 || h <= 0 {
		t.Fatalf("raw dimensions = %dx%d, want positive", w, h)
	}
}

func TestRawImage(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}

	meta, err := p.Metadata()
	if err != nil {
		t.Fatalf("Metadata() error = %v", err)
	}

	data, err := p.RawImage()
	if err != nil {
		t.Fatalf("RawImage() error = %v", err)
	}
	want := int(meta.Sizes.RawPitch/2) * int(meta.Sizes.RawHeight)
	if len(data) != want {
		t.Fatalf("RawImage() len = %d, want %d", len(data), want)
	}

	// The copy must outlive the handle.
	first := data[0]
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if data[0] != first || len(data) != want {
		t.Fatal("RawImage() copy changed after Close")
	}
}

func TestRawImageBeforeUnpack(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	_, err := p.RawImage()
	if !errors.Is(err, ErrNoImageData) {
		t.Fatalf("RawImage() before Unpack error = %v, want ErrNoImageData", err)
	}
}

func TestThumbnailData(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleThumbRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.UnpackThumb(); err != nil {
		t.Fatalf("UnpackThumb() error = %v", err)
	}

	data, err := p.ThumbnailData()
	if err != nil {
		t.Fatalf("ThumbnailData() error = %v", err)
	}
	if len(data) == 0 {
		t.Fatal("ThumbnailData() returned empty data")
	}

	// The copy must remain usable after the handle is closed.
	n := len(data)
	first := data[0]
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if len(data) != n || data[0] != first {
		t.Fatal("ThumbnailData() copy changed after Close")
	}
}

func TestThumbnailDataBeforeUnpack(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleThumbRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}

	_, err := p.ThumbnailData()
	if !errors.Is(err, ErrInvalidState) {
		t.Fatalf("ThumbnailData() before UnpackThumb error = %v, want ErrInvalidState", err)
	}
}

func TestFourChannels(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
	if err := p.Raw2Image(); err != nil {
		t.Fatalf("Raw2Image() error = %v", err)
	}
	if err := p.SubtractBlack(); err != nil {
		t.Fatalf("SubtractBlack() error = %v", err)
	}

	meta, err := p.Metadata()
	if err != nil {
		t.Fatalf("Metadata() error = %v", err)
	}
	want := int(meta.Sizes.IHeight) * int(meta.Sizes.IWidth)

	pixels, err := p.FourChannels()
	if err != nil {
		t.Fatalf("FourChannels() error = %v", err)
	}
	if len(pixels) != want {
		t.Fatalf("FourChannels() len = %d, want %d", len(pixels), want)
	}

	// The copy must outlive the handle.
	first := pixels[0]
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if pixels[0] != first || len(pixels) != want {
		t.Fatal("FourChannels() copy changed after Close")
	}
}

func TestFourChannelsBeforeRaw2Image(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	if err := p.Unpack(); err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}

	_, err := p.FourChannels()
	if !errors.Is(err, ErrInvalidState) {
		t.Fatalf("FourChannels() before Raw2Image error = %v, want ErrInvalidState", err)
	}
}

func TestImageAccessAfterCloseReturnsErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	ops := map[string]func() error{
		"Color":         func() error { _, e := p.Color(0, 0); return e },
		"RawWidth":      func() error { _, e := p.RawWidth(); return e },
		"RawHeight":     func() error { _, e := p.RawHeight(); return e },
		"RawImage":      func() error { _, e := p.RawImage(); return e },
		"FourChannels":  func() error { _, e := p.FourChannels(); return e },
		"ThumbnailData": func() error { _, e := p.ThumbnailData(); return e },
	}
	for name, op := range ops {
		t.Run(name, func(t *testing.T) {
			if err := op(); !errors.Is(err, ErrClosed) {
				t.Fatalf("%s after Close() error = %v, want ErrClosed", name, err)
			}
		})
	}
}

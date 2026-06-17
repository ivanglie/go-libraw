//go:build cgo

package libraw

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestFixtureRegressionOpenIdentify(t *testing.T) {
	for _, fixture := range metadataFixtures {
		t.Run(filepath.Base(fixture), func(t *testing.T) {
			p := openProcessor(t)
			if err := p.OpenFile(fixture); err != nil {
				t.Fatalf("OpenFile(%q) error = %v", fixture, err)
			}

			meta, err := p.Metadata()
			if err != nil {
				t.Fatalf("Metadata() error = %v", err)
			}
			if meta.ID.Make == "" || meta.ID.Model == "" {
				t.Fatalf("camera identity is incomplete: make=%q model=%q", meta.ID.Make, meta.ID.Model)
			}
			if meta.Sizes.RawWidth == 0 || meta.Sizes.RawHeight == 0 {
				t.Fatalf("raw dimensions = %dx%d, want non-zero", meta.Sizes.RawWidth, meta.Sizes.RawHeight)
			}

			decoder, err := p.DecoderInfo()
			if err != nil {
				t.Fatalf("DecoderInfo() error = %v", err)
			}
			if decoder.Name == "" {
				t.Fatal("DecoderInfo().Name is empty")
			}
		})
	}
}

func TestFixtureRegressionProcessingSmoke(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping RAW processing smoke test in short mode")
	}

	for _, fixture := range metadataFixtures {
		t.Run(filepath.Base(fixture), func(t *testing.T) {
			p := openProcessor(t)
			if err := p.OpenFile(fixture); err != nil {
				t.Fatalf("OpenFile(%q) error = %v", fixture, err)
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
			if img.Width == 0 || img.Height == 0 {
				t.Fatalf("MemImage() dimensions = %dx%d, want non-zero", img.Width, img.Height)
			}
			if img.Colors == 0 || img.Bits == 0 || len(img.Data) == 0 {
				t.Fatalf("MemImage() payload invalid: colors=%d bits=%d data=%d", img.Colors, img.Bits, len(img.Data))
			}
		})
	}
}

func TestFixtureRegressionThumbnailSmoke(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping thumbnail smoke test in short mode")
	}

	tested := 0
	for _, fixture := range metadataFixtures {
		t.Run(filepath.Base(fixture), func(t *testing.T) {
			p := openProcessor(t)
			if err := p.OpenFile(fixture); err != nil {
				t.Fatalf("OpenFile(%q) error = %v", fixture, err)
			}
			meta, err := p.Metadata()
			if err != nil {
				t.Fatalf("Metadata() error = %v", err)
			}
			if meta.Thumbnail.Width == 0 || meta.Thumbnail.Height == 0 {
				t.Skip("fixture does not advertise an embedded thumbnail")
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
			tested++
		})
	}
	if tested == 0 {
		t.Fatal("no fixture with thumbnail data was tested")
	}
}

func TestInvalidInputRegressionErrorsAreTyped(t *testing.T) {
	tests := map[string]func(*Processor) error{
		"missing file": func(p *Processor) error {
			return p.OpenFile(filepath.Join(t.TempDir(), "missing.raw"))
		},
		"empty buffer": func(p *Processor) error {
			return p.OpenBuffer(nil)
		},
		"garbage buffer": func(p *Processor) error {
			return p.OpenBuffer([]byte("not a raw file"))
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			p := openProcessor(t)
			err := run(p)
			if err == nil {
				t.Fatal("invalid input returned nil error")
			}
			var librawErr Error
			if !errors.As(err, &librawErr) {
				t.Fatalf("invalid input error = %T (%v), want libraw.Error", err, err)
			}
		})
	}
}

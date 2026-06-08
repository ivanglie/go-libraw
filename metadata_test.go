//go:build cgo

package libraw

import (
	"errors"
	"path/filepath"
	"testing"
)

var metadataFixtures = []string{
	"testdata/RAW_CANON_6D.CR2",
	"testdata/RAW_NIKON_D750.NEF",
	"testdata/RAW_RICOH_GR2.DNG",
	"testdata/RAW_SONY_ILCA-77M2.ARW",
}

func TestMetadataForFixtures(t *testing.T) {
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

			if meta.ID.Make == "" {
				t.Fatalf("ID.Make is empty: %+v", meta.ID)
			}
			if meta.ID.Model == "" {
				t.Fatalf("ID.Model is empty: %+v", meta.ID)
			}
			if meta.Sizes.RawWidth == 0 || meta.Sizes.RawHeight == 0 {
				t.Fatalf("raw dimensions = %dx%d, want non-zero", meta.Sizes.RawWidth, meta.Sizes.RawHeight)
			}
			if meta.Sizes.Width == 0 || meta.Sizes.Height == 0 {
				t.Fatalf("image dimensions = %dx%d, want non-zero", meta.Sizes.Width, meta.Sizes.Height)
			}
			if meta.ID.Colors <= 0 {
				t.Fatalf("ID.Colors = %d, want positive", meta.ID.Colors)
			}
			if meta.Color.Maximum == 0 && meta.Color.DataMaximum == 0 {
				t.Fatalf("color maximums are empty: maximum=%d data_maximum=%d", meta.Color.Maximum, meta.Color.DataMaximum)
			}
			if meta.Other.ISOSpeed == 0 {
				t.Fatalf("Other.ISOSpeed = %v, want non-zero", meta.Other.ISOSpeed)
			}
		})
	}
}

func TestMetadataSnapshotSurvivesClose(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	meta, err := p.Metadata()
	if err != nil {
		t.Fatalf("Metadata() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	if meta.ID.Make == "" || meta.Sizes.RawWidth == 0 {
		t.Fatalf("metadata snapshot was not populated before close: %+v", meta)
	}
}

func TestMetadataAfterCloseReturnsErrClosed(t *testing.T) {
	p, err := NewProcessor()
	if err != nil {
		t.Fatalf("NewProcessor() error = %v", err)
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if _, err := p.Metadata(); !errors.Is(err, ErrClosed) {
		t.Fatalf("Metadata() after Close error = %v, want ErrClosed", err)
	}
}

func TestMetadataThumbnailList(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleThumbRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	meta, err := p.Metadata()
	if err != nil {
		t.Fatalf("Metadata() error = %v", err)
	}
	if meta.Thumbs.Count == 0 {
		t.Fatal("Thumbs.Count = 0, want at least one thumbnail entry")
	}
	if len(meta.Thumbs.Items) == 0 {
		t.Fatal("Thumbs.Items is empty")
	}
	for i, item := range meta.Thumbs.Items {
		if item.Width == 0 || item.Height == 0 {
			t.Fatalf("thumbnail %d dimensions = %dx%d, want non-zero", i, item.Width, item.Height)
		}
	}
}

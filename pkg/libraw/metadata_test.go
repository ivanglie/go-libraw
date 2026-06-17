//go:build cgo

package libraw

import (
	"errors"
	"path/filepath"
	"testing"
)

var metadataFixtures = []string{
	"../../testdata/RAW_CANON_R6.CR3",
	"../../testdata/RAW_NIKON_ZFC.NEF",
	"../../testdata/RAW_RICOH_GR3X.DNG",
	"../../testdata/RAW_SONY_ILCE-7M4.ARW",
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
	// LibRaw's thumbnail list can include preview entries it does not assign
	// dimensions to (CR3, for example, exposes multiple preview types). Require
	// at least one sized thumbnail rather than every entry being non-zero.
	sized := 0
	for _, item := range meta.Thumbs.Items {
		if item.Width > 0 && item.Height > 0 {
			sized++
		}
	}
	if sized == 0 {
		t.Fatalf("no thumbnail entry has non-zero dimensions: %+v", meta.Thumbs.Items)
	}
}

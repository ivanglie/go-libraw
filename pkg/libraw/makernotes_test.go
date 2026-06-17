//go:build cgo

package libraw

import (
	"path/filepath"
	"testing"
)

func TestMakerNotesForFixtures(t *testing.T) {
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

			switch filepath.Ext(fixture) {
			case ".CR3":
				if meta.MakerNotes.Canon.SensorWidth == 0 && meta.MakerNotes.Canon.ColorDataVer == 0 {
					t.Fatalf("Canon maker notes look empty: %+v", meta.MakerNotes.Canon)
				}
			case ".NEF":
				if meta.MakerNotes.Nikon.NEFCompression == 0 && meta.MakerNotes.Nikon.SensorWidth == 0 {
					t.Fatalf("Nikon maker notes look empty: %+v", meta.MakerNotes.Nikon)
				}
			case ".ARW":
				if meta.MakerNotes.Sony.CameraType == 0 && meta.MakerNotes.Sony.FileFormat == 0 {
					t.Fatalf("Sony maker notes look empty: %+v", meta.MakerNotes.Sony)
				}
			case ".DNG":
				if meta.MakerNotes.Common.AFCount < 0 {
					t.Fatalf("Common.AFCount = %d, want non-negative", meta.MakerNotes.Common.AFCount)
				}
			}
		})
	}
}

func TestMakerNotesZeroValueForAbsentVendor(t *testing.T) {
	p := openProcessor(t)
	if err := p.OpenFile(sampleRAW); err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	meta, err := p.Metadata()
	if err != nil {
		t.Fatalf("Metadata() error = %v", err)
	}
	if meta.MakerNotes.Hasselblad.Sensor != "" {
		t.Fatalf("Hasselblad.Sensor = %q, want empty for Ricoh fixture", meta.MakerNotes.Hasselblad.Sensor)
	}
}

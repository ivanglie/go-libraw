// Command raw-identify prints a concise metadata summary for a RAW file.
//
// It mirrors LibRaw's upstream raw-identify.cpp sample. Pass a RAW path as the
// first argument, or it defaults to a bundled fixture.
package main

import (
	"fmt"
	"log"
	"os"

	libraw "github.com/ivanglie/go-libraw"
)

func main() {
	path := "testdata/RAW_CANON_6D.CR2"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	p, err := libraw.NewProcessor()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := p.Close(); err != nil {
			log.Printf("close processor: %v", err)
		}
	}()

	if err := p.OpenFile(path); err != nil {
		log.Fatalf("open %s: %v", path, err)
	}
	meta, err := p.Metadata()
	if err != nil {
		log.Fatalf("metadata: %v", err)
	}
	decoder, err := p.DecoderInfo()
	if err != nil {
		log.Fatalf("decoder: %v", err)
	}

	fmt.Printf("File:     %s\n", path)
	fmt.Printf("Camera:   %s %s\n", meta.ID.Make, meta.ID.Model)
	fmt.Printf("Normalized: %s %s\n", meta.ID.NormalizedMake, meta.ID.NormalizedModel)
	fmt.Printf("RawSize:  %dx%d\n", meta.Sizes.RawWidth, meta.Sizes.RawHeight)
	fmt.Printf("OutSize:  %dx%d\n", meta.Sizes.Width, meta.Sizes.Height)
	fmt.Printf("Colors:   %d (%s)\n", meta.ID.Colors, meta.ID.CDesc)
	fmt.Printf("ISO:      %.0f\n", meta.Other.ISOSpeed)
	fmt.Printf("Shutter:  %s\n", shutter(meta.Other.Shutter))
	fmt.Printf("Aperture: f/%.1f\n", meta.Other.Aperture)
	fmt.Printf("FocalLen: %.0f mm\n", meta.Other.FocalLen)
	fmt.Printf("Lens:     %s\n", meta.Lens.Lens)
	fmt.Printf("Decoder:  %s (flags 0x%x)\n", decoder.Name, decoder.Flags)
}

func shutter(s float32) string {
	if s > 0 && s < 1 {
		return fmt.Sprintf("1/%.0f s", 1/s)
	}
	return fmt.Sprintf("%.1f s", s)
}

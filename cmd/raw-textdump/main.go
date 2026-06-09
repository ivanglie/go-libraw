// Command raw-textdump prints a verbose key/value dump of RAW metadata.
//
// It mirrors LibRaw's upstream rawtextdump.cpp sample (metadata portion). Pass a
// RAW path as the first argument, or it defaults to a bundled fixture.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	libraw "github.com/ivanglie/go-libraw"
)

func main() {
	path := "testdata/RAW_NIKON_D750.NEF"
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

	kv := func(key string, value any) { fmt.Printf("%-22s %v\n", key+":", value) }

	kv("file", path)
	kv("libraw", libraw.Version())
	kv("make", meta.ID.Make)
	kv("model", meta.ID.Model)
	kv("normalized_make", meta.ID.NormalizedMake)
	kv("normalized_model", meta.ID.NormalizedModel)
	kv("software", meta.ID.Software)
	kv("raw_count", meta.ID.RawCount)
	kv("dng_version", meta.ID.DNGVersion)
	kv("colors", meta.ID.Colors)
	kv("color_desc", meta.ID.CDesc)
	kv("filters", fmt.Sprintf("0x%08x", meta.ID.Filters))
	kv("raw_size", fmt.Sprintf("%dx%d", meta.Sizes.RawWidth, meta.Sizes.RawHeight))
	kv("output_size", fmt.Sprintf("%dx%d", meta.Sizes.Width, meta.Sizes.Height))
	kv("margins", fmt.Sprintf("top=%d left=%d", meta.Sizes.TopMargin, meta.Sizes.LeftMargin))
	kv("flip", meta.Sizes.Flip)
	kv("iso", meta.Other.ISOSpeed)
	kv("shutter", meta.Other.Shutter)
	kv("aperture", meta.Other.Aperture)
	kv("focal_len", meta.Other.FocalLen)
	kv("timestamp", time.Unix(meta.Other.Timestamp, 0).UTC().Format(time.RFC3339))
	kv("artist", meta.Other.Artist)
	kv("lens", meta.Lens.Lens)
	kv("lens_make", meta.Lens.LensMake)
	kv("focal_35mm", meta.Lens.FocalLengthIn35mmFormat)
	kv("thumb_format", meta.Thumbnail.Format)
	kv("thumb_size", fmt.Sprintf("%dx%d", meta.Thumbnail.Width, meta.Thumbnail.Height))
	kv("thumb_list", meta.Thumbs.Count)
}

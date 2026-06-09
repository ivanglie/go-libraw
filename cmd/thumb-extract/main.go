// Command thumb-extract unpacks and writes the embedded thumbnail of a RAW file.
//
// It mirrors LibRaw's thumbnail extraction samples. Pass a RAW path as the first
// argument, or it defaults to a bundled fixture. Output goes to tmp/examples/.
package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

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
	if err := p.UnpackThumb(); err != nil {
		log.Fatalf("unpack thumb: %v", err)
	}
	data, err := p.ThumbnailData()
	if err != nil {
		log.Fatalf("thumbnail data: %v", err)
	}

	meta, err := p.Metadata()
	if err != nil {
		log.Fatalf("metadata: %v", err)
	}

	ext := ".bin"
	if meta.Thumbnail.Format == libraw.LIBRAW_THUMBNAIL_JPEG {
		ext = ".jpg"
	}
	out := filepath.Join("tmp/examples", base(path)+".thumb"+ext)
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(out, data, 0o644); err != nil {
		log.Fatalf("write %s: %v", out, err)
	}
	log.Printf("thumbnail %dx%d (%d bytes) -> %s", meta.Thumbnail.Width, meta.Thumbnail.Height, len(data), out)
}

func base(path string) string {
	name := filepath.Base(path)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

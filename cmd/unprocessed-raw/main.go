// Command unprocessed-raw dumps the single-channel raw Bayer buffer to a
// 16-bit PGM image under tmp/examples/.
//
// It mirrors LibRaw's upstream unprocessed_raw.cpp sample. Pass a RAW path as
// the first argument, or it defaults to a bundled fixture.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	libraw "github.com/strawberryworks/go-libraw/pkg/libraw"
)

func main() {
	path := "testdata/RAW_RICOH_GR3X.DNG"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	const outDir = "tmp/examples"
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatal(err)
	}

	p, err := libraw.NewProcessor()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = p.Close() }()

	if err := p.OpenFile(path); err != nil {
		log.Fatalf("open %s: %v", path, err)
	}
	if err := p.Unpack(); err != nil {
		log.Fatalf("unpack: %v", err)
	}

	pixels, err := p.RawImage()
	if err != nil {
		log.Fatalf("raw image: %v", err)
	}
	meta, err := p.Metadata()
	if err != nil {
		log.Fatalf("metadata: %v", err)
	}
	w := int(meta.Sizes.RawWidth)
	h := int(meta.Sizes.RawHeight)
	pitch := int(meta.Sizes.RawPitch) / 2 // uint16s per row (may include padding)

	out := filepath.Join(outDir, filepath.Base(path)+".pgm")
	if err := writePGM16(out, w, h, pitch, pixels); err != nil {
		log.Fatalf("write %s: %v", out, err)
	}
	log.Printf("LibRaw %s: %s (%dx%d) -> %s", libraw.Version(), path, w, h, out)
}

// writePGM16 writes pixels as a 16-bit binary PGM (P5). Each row has pitch
// uint16 elements in the source buffer; only the first width are written.
func writePGM16(path string, width, height, pitch int, data []uint16) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	bw := bufio.NewWriterSize(f, 1<<20)
	if _, err := fmt.Fprintf(bw, "P5\n%d %d\n65535\n", width, height); err != nil {
		_ = f.Close()
		return err
	}

	tmp := make([]byte, 2*width)
	for row := 0; row < height; row++ {
		base := row * pitch
		if base+width > len(data) {
			break
		}
		for col, v := range data[base : base+width] {
			tmp[2*col] = byte(v >> 8)
			tmp[2*col+1] = byte(v)
		}
		if _, err := bw.Write(tmp); err != nil {
			_ = f.Close()
			return err
		}
	}
	if err := bw.Flush(); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

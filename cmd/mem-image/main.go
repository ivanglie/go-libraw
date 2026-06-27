// Command mem-image develops a RAW into an in-memory image and writes it as a
// PPM, demonstrating Processor.MemImage.
//
// It mirrors LibRaw's upstream mem_image_sample.cpp. Pass a RAW path as the
// first argument, or it defaults to a bundled fixture. Output goes to
// tmp/examples/.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	libraw "github.com/strawberryworks/go-libraw/pkg/libraw"
)

func main() {
	path := "testdata/RAW_RICOH_GR3X.DNG"
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
	if err := p.Unpack(); err != nil {
		log.Fatalf("unpack: %v", err)
	}
	if err := p.DcrawProcess(); err != nil {
		log.Fatalf("process: %v", err)
	}
	img, err := p.MemImage()
	if err != nil {
		log.Fatalf("mem image: %v", err)
	}
	if img.Colors != 3 || img.Bits != 8 {
		log.Fatalf("unexpected image: colors=%d bits=%d, want a 3x8 bitmap", img.Colors, img.Bits)
	}

	out := filepath.Join("tmp/examples", base(path)+".mem.ppm")
	if err := writePPM(out, img); err != nil {
		log.Fatalf("write %s: %v", out, err)
	}
	log.Printf("MemImage %dx%d (%d bytes) -> %s", img.Width, img.Height, len(img.Data), out)
}

func writePPM(path string, img *libraw.ProcessedImage) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	if _, err := fmt.Fprintf(w, "P6\n%d %d\n255\n", img.Width, img.Height); err != nil {
		_ = f.Close()
		return err
	}
	if _, err := w.Write(img.Data); err != nil {
		_ = f.Close()
		return err
	}
	if err := w.Flush(); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

func base(path string) string {
	name := filepath.Base(path)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

// Command four-channels separates the postprocessing image buffer into its
// four CFA channels and writes each as a 16-bit PGM under tmp/examples/.
//
// It mirrors LibRaw's upstream 4channels.cpp sample. Pass a RAW path as the
// first argument, or it defaults to a bundled fixture.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	libraw "github.com/ivanglie/go-libraw/pkg/libraw"
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
	defer p.Close()

	if err := p.OpenFile(path); err != nil {
		log.Fatalf("open %s: %v", path, err)
	}
	if err := p.Unpack(); err != nil {
		log.Fatalf("unpack: %v", err)
	}
	if err := p.Raw2Image(); err != nil {
		log.Fatalf("raw2image: %v", err)
	}
	if err := p.SubtractBlack(); err != nil {
		log.Fatalf("subtract black: %v", err)
	}

	pixels, err := p.FourChannels()
	if err != nil {
		log.Fatalf("four channels: %v", err)
	}
	meta, err := p.Metadata()
	if err != nil {
		log.Fatalf("metadata: %v", err)
	}
	w := int(meta.Sizes.IWidth)
	h := int(meta.Sizes.IHeight)

	base := filepath.Join(outDir, strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
	labels := [4]string{"R", "Gr", "Gb", "B"}

	for ch := range 4 {
		out := fmt.Sprintf("%s.%s.pgm", base, labels[ch])
		if err := writeChannel(out, w, h, ch, pixels); err != nil {
			log.Fatalf("write channel %s: %v", labels[ch], err)
		}
		log.Printf("channel %s (%dx%d): %s", labels[ch], w, h, out)
	}
}

// writeChannel writes a single channel from a [][4]uint16 buffer as a 16-bit PGM.
func writeChannel(path string, width, height, ch int, pixels [][4]uint16) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriterSize(f, 1<<20)
	fmt.Fprintf(bw, "P5\n%d %d\n65535\n", width, height)

	tmp := make([]byte, 2*width)
	for row := range height {
		for col := range width {
			v := pixels[row*width+col][ch]
			tmp[2*col] = byte(v >> 8)
			tmp[2*col+1] = byte(v)
		}
		if _, err := bw.Write(tmp); err != nil {
			return err
		}
	}
	return bw.Flush()
}

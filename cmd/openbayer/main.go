// Command openbayer develops a synthetic RGGB Bayer buffer into a PPM image
// under tmp/examples/ using LibRaw's OpenBayer path.
//
// It mirrors LibRaw's upstream openbayer_sample.cpp sample. The upstream
// sample reads a raw binary sensor dump from disk; this version generates a
// synthetic 128×128 gradient internally so no external fixture is required.
package main

import (
	"encoding/binary"
	"log"
	"os"
	"path/filepath"

	libraw "github.com/ivanglie/go-libraw/pkg/libraw"
)

const (
	bayerW   = 128
	bayerH   = 128
	bayerMax = 4095 // 12-bit range
)

func main() {
	const outDir = "tmp/examples"
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatal(err)
	}

	data := makeSyntheticBayer(bayerW, bayerH)

	p, err := libraw.NewProcessor()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	params := libraw.BayerParams{
		RawWidth:     bayerW,
		RawHeight:    bayerH,
		BayerPattern: uint8(libraw.LIBRAW_OPENBAYER_RGGB),
	}
	if err := p.OpenBayer(data, params); err != nil {
		log.Fatalf("open bayer: %v", err)
	}
	if err := p.Unpack(); err != nil {
		log.Fatalf("unpack: %v", err)
	}
	if err := p.DcrawProcess(); err != nil {
		log.Fatalf("process: %v", err)
	}

	out := filepath.Join(outDir, "openbayer.ppm")
	if err := p.WritePPMTiff(out); err != nil {
		log.Fatalf("write %s: %v", out, err)
	}
	log.Printf("LibRaw %s: synthetic %dx%d RGGB -> %s", libraw.Version(), bayerW, bayerH, out)
}

// makeSyntheticBayer generates a width×height RGGB Bayer buffer (two bytes per
// pixel, little-endian uint16, 12-bit values). R varies left-to-right, B
// varies top-to-bottom, G is the diagonal midpoint.
func makeSyntheticBayer(width, height int) []byte {
	buf := make([]byte, width*height*2)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			rowEven := row%2 == 0
			colEven := col%2 == 0
			var v uint16
			switch {
			case rowEven && colEven: // R
				v = uint16(col * bayerMax / (width - 1))
			case !rowEven && !colEven: // B
				v = uint16(row * bayerMax / (height - 1))
			default: // G (Gr and Gb)
				v = uint16((col + row) * bayerMax / (width + height - 2))
			}
			binary.LittleEndian.PutUint16(buf[(row*width+col)*2:], v)
		}
	}
	return buf
}

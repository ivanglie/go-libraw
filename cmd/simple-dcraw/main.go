// Command simple-dcraw develops a RAW file into a PPM image under tmp/outputs/
// using LibRaw's dcraw-equivalent pipeline. It mirrors LibRaw's upstream
// simple_dcraw.cpp sample. Pass a RAW path as the first argument, or it
// defaults to a bundled fixture.
package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	libraw "github.com/ivanglie/go-libraw/pkg/libraw"
)

func main() {
	input := "testdata/RAW_CANON_R6.CR3"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	const outDir = "tmp/outputs"
	stem := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
	output := filepath.Join(outDir, stem+".ppm")

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatal(err)
	}

	processor, err := libraw.NewProcessor()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := processor.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := processor.OpenFile(input); err != nil {
		log.Fatalf("open %s: %v", input, err)
	}
	if err := processor.Unpack(); err != nil {
		log.Fatalf("unpack: %v", err)
	}
	if err := processor.DcrawProcess(); err != nil {
		log.Fatalf("process: %v", err)
	}
	if err := processor.WritePPMTiff(output); err != nil {
		log.Fatalf("write %s: %v", output, err)
	}

	log.Printf("LibRaw %s developed %s -> %s", libraw.Version(), input, output)
}

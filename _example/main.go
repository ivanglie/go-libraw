// Command example develops a bundled sample RAW into a PPM image under
// tmp/outputs/ using LibRaw's dcraw-equivalent pipeline. PPM is LibRaw's native
// writer format; the output directory is git-ignored. Run it from the repo root.
package main

import (
	"log"
	"os"
	"path/filepath"

	libraw "github.com/ivanglie/go-libraw"
)

func main() {
	const input = "testdata/RAW_CANON_6D.CR2"
	const outDir = "tmp/outputs"
	output := filepath.Join(outDir, "RAW_CANON_6D.ppm")

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

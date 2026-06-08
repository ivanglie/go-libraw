// Command example develops a bundled sample RAW to a JPEG (with embedded Exif)
// next to the source. Run it from the repository root.
package main

import (
	"log"

	libraw "github.com/ivanglie/go-libraw"
)

func main() {
	p := "testdata/RAW_CANON_6D.CR2"
	processor, err := libraw.NewProcessor()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := processor.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("LibRaw %s ready for %s", libraw.Version(), p)
}

package libraw_test

import (
	"log"

	libraw "github.com/ivanglie/go-libraw"
)

func ExampleProcessor_quickStart() {
	processor, err := libraw.NewProcessor()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := processor.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := processor.OpenFile("input.cr2"); err != nil {
		log.Fatal(err)
	}
	if err := processor.Unpack(); err != nil {
		log.Fatal(err)
	}
	if err := processor.DcrawProcess(); err != nil {
		log.Fatal(err)
	}
	if err := processor.WritePPMTiff("output.ppm"); err != nil {
		log.Fatal(err)
	}
}

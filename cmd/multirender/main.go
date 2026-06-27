// Command multirender develops the same RAW file four times with different
// output parameters and writes each render to a separate PPM under
// tmp/examples/.
//
// It mirrors LibRaw's upstream multirender_test.cpp sample. Pass a RAW path
// as the first argument, or it defaults to a bundled fixture.
package main

import (
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

	// Snapshot params after open so each render starts from the same base.
	// Modifying only the delta fields avoids clobbering internal params that
	// LibRaw sets during parsing (e.g. DNG crop box, flip).
	base0, err := p.OutputParams()
	if err != nil {
		log.Fatalf("output params: %v", err)
	}

	outBase := filepath.Join(outDir, strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))

	type renderCfg struct {
		label string
		apply func(*libraw.OutputParams)
	}
	renders := []renderCfg{
		{"default", func(op *libraw.OutputParams) {}},
		{"half", func(op *libraw.OutputParams) { op.HalfSize = 1 }},
		{"camerawb", func(op *libraw.OutputParams) { op.UseCameraWB = 1 }},
		{"noautobright", func(op *libraw.OutputParams) { op.NoAutoBright = 1 }},
	}

	for i, r := range renders {
		params := base0
		r.apply(&params)
		if err := p.SetOutputParams(params); err != nil {
			log.Fatalf("set params for %s: %v", r.label, err)
		}
		if err := p.DcrawProcess(); err != nil {
			log.Fatalf("process %s: %v", r.label, err)
		}
		out := fmt.Sprintf("%s.%d-%s.ppm", outBase, i+1, r.label)
		if err := p.WritePPMTiff(out); err != nil {
			log.Fatalf("write %s: %v", out, err)
		}
		log.Printf("render %d (%s): %s", i+1, r.label, out)
	}
}

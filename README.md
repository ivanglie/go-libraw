# go-libraw

`go-libraw` provides Go bindings for [LibRaw](https://www.libraw.org/docs),
the RAW image decoding library. It keeps LibRaw's lifecycle visible while
returning Go-owned snapshots and byte slices where that makes memory ownership
safer.

## Requirements

- Go `1.26` or newer, as declared by `go.mod`
- cgo enabled
- LibRaw development headers and library
- a C/C++ toolchain for the target platform

Install LibRaw on common platforms:

```sh
# macOS
brew install libraw

# Debian/Ubuntu
sudo apt-get update
sudo apt-get install -y libraw-dev pkg-config

# Fedora
sudo dnf install LibRaw-devel pkgconf-pkg-config
```

Verify local discovery:

```sh
make libraw-check
```

See [LibRaw Build Setup](docs/libraw-build.md) and the
[Support Matrix](docs/support-matrix.md) for platform details.

## Quick Start

```go
package main

import (
	"log"

	libraw "github.com/ivanglie/go-libraw"
)

func main() {
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
```

Run the bundled example on checked-in fixtures:

```sh
make examples
make clean
```

The sample commands are mapped to upstream LibRaw samples in
[LibRaw Sample Parity Examples](docs/examples.md).

## API Concepts

- `Processor` owns one LibRaw handle. Create it with `NewProcessor`, then call
  `Close` when finished.
- Opening input with `OpenFile`, `OpenBuffer`, or `OpenBayer` prepares metadata
  and decoder state.
- Processing follows LibRaw order: open, optionally set params, `Unpack`,
  `DcrawProcess` or lower-level image operations, then write or copy output.
- `Metadata` returns a Go snapshot of LibRaw metadata and maker-note summaries.
- Raw image, thumbnail, and memory image helpers return Go-owned data.
- LibRaw error codes are returned as Go errors; use `ErrorCode` and `StrError`
  when you need to inspect an underlying LibRaw status.

## Documentation

- [Lifecycle And Processing](docs/lifecycle-processing.md)
- [Memory And Cgo Safety](docs/memory-and-cgo.md)
- [API Coverage Guide](docs/api-coverage.md)
- [LibRaw API Inventory](docs/libraw-api-inventory.md)
- [Metadata Coverage](docs/libraw-metadata-coverage.md)
- [Maker-Notes Coverage](docs/libraw-maker-notes-coverage.md)
- [Output And Raw Params Coverage](docs/libraw-params-coverage.md)

## Upstream Coverage

The generated inventory in [docs/libraw-api-inventory.md](docs/libraw-api-inventory.md)
tracks LibRaw symbols from the fixture headers and marks each as `wrapped`,
`internal`, `deferred`, `unsupported`, or `unmapped`. Run this before changing
coverage-related code:

```sh
make check-api-inventory
```

To regenerate the inventory after updating the coverage map:

```sh
make api-inventory
```

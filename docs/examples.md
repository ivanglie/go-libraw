# LibRaw Sample Parity Examples

These small Go programs mirror representative upstream LibRaw C++ samples and
demonstrate the wrapper end to end. Run them all with `make examples`; generated
files land under `tmp/examples/` and are removed by `make clean`.

Each command takes an optional RAW path as its first argument and otherwise uses
a bundled fixture from `testdata/`.

| Go example | Upstream sample | What it shows |
| --- | --- | --- |
| [`_example`](../_example/main.go) | `simple_dcraw.cpp` | Develop a RAW to a PPM via `Unpack` → `DcrawProcess` → `WritePPMTiff`. |
| [`cmd/raw-identify`](../cmd/raw-identify/main.go) | `raw-identify.cpp` | Open a RAW and print a concise camera/exposure/lens summary. |
| [`cmd/raw-textdump`](../cmd/raw-textdump/main.go) | `rawtextdump.cpp` | Print a verbose key/value dump of the metadata snapshot. |
| [`cmd/mem-image`](../cmd/mem-image/main.go) | `mem_image_sample.cpp` | Develop to an in-memory image (`MemImage`) and write a PPM from the bytes. |
| [`cmd/thumb-extract`](../cmd/thumb-extract/main.go) | thumbnail samples (`dcraw_emu -e`) | Unpack the embedded thumbnail and write its bytes (`ThumbnailData`). |

## Running

```sh
make examples            # run all on bundled fixtures
make clean               # remove tmp/examples (and other generated output)

# or run one directly with your own file:
go run ./cmd/raw-identify /path/to/file.cr2
```

## Not covered

Full `dcraw_emu.cpp` CLI flag parity and a benchmark/multithreaded suite are out
of scope; the examples above favor readability over exhaustive option coverage.

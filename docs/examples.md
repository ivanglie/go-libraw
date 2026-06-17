# LibRaw Sample Parity Examples

These small Go programs mirror representative upstream LibRaw C++ samples and
demonstrate the wrapper end to end. Run them all with `make examples`; generated
files land under `tmp/examples/` and are removed by `make clean`.

Each command takes an optional RAW path as its first argument and otherwise uses
a bundled fixture from `testdata/`.

| Go example | Upstream sample | What it shows |
| --- | --- | --- |
| [`cmd/simple-dcraw`](../cmd/simple-dcraw/main.go) | `simple_dcraw.cpp` | Develop a RAW to a PPM via `Unpack` → `DcrawProcess` → `WritePPMTiff`. |
| [`cmd/raw-identify`](../cmd/raw-identify/main.go) | `raw-identify.cpp` | Open a RAW and print a concise camera/exposure/lens summary. |
| [`cmd/raw-textdump`](../cmd/raw-textdump/main.go) | `rawtextdump.cpp` | Print a verbose key/value dump of the metadata snapshot. |
| [`cmd/mem-image`](../cmd/mem-image/main.go) | `mem_image_sample.cpp` | Develop to an in-memory image (`MemImage`) and write a PPM from the bytes. |
| [`cmd/thumb-extract`](../cmd/thumb-extract/main.go) | thumbnail samples (`dcraw_emu -e`) | Unpack the embedded thumbnail and write its bytes (`ThumbnailData`). |
| [`cmd/unprocessed-raw`](../cmd/unprocessed-raw/main.go) | `unprocessed_raw.cpp` | Dump the raw Bayer sensor buffer as a 16-bit PGM via `RawImage`. |
| [`cmd/openbayer`](../cmd/openbayer/main.go) | `openbayer_sample.cpp` | Open a synthetic RGGB Bayer buffer with `OpenBayer`, then develop and write a PPM. |
| [`cmd/multirender`](../cmd/multirender/main.go) | `multirender_test.cpp` | Produce four renders from a single `Unpack` by changing `OutputParams` between `DcrawProcess` calls. |
| [`cmd/four-channels`](../cmd/four-channels/main.go) | `4channels.cpp` | Separate the `Raw2Image` buffer into four CFA channels and write each as a 16-bit PGM via `FourChannels`. |

## Running

```sh
make examples            # run all on bundled fixtures
make clean               # remove tmp/outputs and tmp/examples

# or run one directly with your own file:
go run ./cmd/simple-dcraw /path/to/file.cr2
```

## Not covered

`dcraw_emu.cpp` (full CLI flag parity) and `postprocessing_benchmark.cpp` are
out of scope; the examples above favor readability over exhaustive option
coverage.

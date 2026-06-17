# Lifecycle And Processing

`go-libraw` follows LibRaw's processor model: one `Processor` owns one LibRaw
handle and processes one opened input at a time. This mirrors the upstream
[LibRaw API overview](https://www.libraw.org/node/29), where a processor can be
reused for consecutive inputs but does not process multiple data sources at the
same time.

## Basic Lifecycle

1. Create a processor with `NewProcessor`.
2. Optionally pass constructor flags with `WithFlags`.
3. Open input with `OpenFile`, `OpenBuffer`, or `OpenBayer`.
4. Read metadata with `Metadata` or configure output/raw-unpack params.
5. Decode with `Unpack` or `UnpackThumb`.
6. Process with `DcrawProcess`, `Raw2Image`, or thumbnail helpers.
7. Write output or copy Go-owned image bytes.
8. Call `Close`.

`Close` is idempotent. All public methods that need a live LibRaw handle return
`ErrClosed` after the processor is closed.

## Opening Input

Use `OpenFile` for normal RAW paths. Use `OpenBuffer` when the RAW file is
already loaded into memory; the bytes are copied into LibRaw-owned memory, so
the caller may reuse or release the input slice after `OpenBuffer` returns.

Use `OpenBayer` for raw Bayer samples plus explicit geometry in `BayerParams`.
`BayerPattern` accepts the generated `LIBRAW_OPENBAYER_*` constants.

Opening a new input replaces the previous input. Use `Recycle` when processing
many files with the same processor and you want to clear the old input before
opening the next one.

## Metadata

`Metadata` can be called after input is opened. It returns a Go snapshot of
LibRaw image params, sizes, color data, thumbnail information, lens data,
shooting information, and maker-note summaries.

Because metadata is copied into Go values, the returned snapshot remains valid
after `Recycle` or `Close`.

## Processing Flow

The common dcraw-style path is:

```go
if err := p.OpenFile(path); err != nil {
	return err
}
if err := p.Unpack(); err != nil {
	return err
}
if err := p.DcrawProcess(); err != nil {
	return err
}
return p.WritePPMTiff(out)
```

For raw data analysis:

- call `Unpack`
- call `RawImage` for a Go-owned copy of the single-channel raw buffer
- use `RawWidth`, `RawHeight`, `Color`, `CamMul`, `PreMul`, and `RGBCam` for
  sensor geometry and color helper data

For embedded thumbnails:

- call `UnpackThumb` or `UnpackThumbEx`
- call `WriteThumb`, `MemThumb`, or `ThumbnailData`

For in-memory processed output:

- call `Unpack`
- call `DcrawProcess`
- call `MemImage`

`SubtractBlack` operates on LibRaw's postprocessing image buffer. Build that
buffer first with `Raw2Image` or a processing path that creates it.

## Parameters

Set output and raw-unpack params before the processing step that consumes them.
Use `WithOutputParams` or `WithRawUnpackParams` at construction time, or call
the corresponding setters on a live processor.

See [Output And Raw Params Coverage](libraw-params-coverage.md) for the LibRaw
field mapping.

## Callbacks

Callbacks can observe progress, data errors, EXIF parsing, and maker-note
parsing:

- `SetProgressHandler`
- `SetDataErrorHandler`
- `SetExifParserHandler`
- `SetMakerNotesHandler`

The progress callback can cancel processing by returning a non-zero value.
Callback panics are recovered and treated as cancellation at the cgo boundary.

## Concurrency

A single `Processor` serializes public calls with an internal lock, but it still
represents one LibRaw handle and one active input. For concurrent processing,
create one processor per worker/input.

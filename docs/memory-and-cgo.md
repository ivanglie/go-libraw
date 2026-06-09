# Memory And Cgo Safety

LibRaw owns several internal buffers during decoding. `go-libraw` exposes those
buffers either as file output operations or as Go-owned copies, so callers do
not need to free C memory directly.

## Processor Ownership

`Processor` owns the underlying LibRaw handle. Call `Close` when finished:

```go
p, err := libraw.NewProcessor()
if err != nil {
	return err
}
defer p.Close()
```

`Close` releases the handle and any callback state registered through the Go
API. It is safe to call more than once.

## Input Buffers

`OpenFile` lets LibRaw read from a path.

`OpenBuffer` and `OpenBayer` copy the supplied Go bytes into memory managed by
LibRaw. The caller may reuse or release the original slice as soon as the open
method returns.

## Go-Owned Results

These APIs return copies owned by Go:

- `Metadata`
- `RawImage`
- `ThumbnailData`
- `MemImage`
- `MemThumb`

The returned values remain valid after `FreeImage`, `Recycle`, or `Close`.

`MemImage` and `MemThumb` use LibRaw's in-memory image producers internally.
The temporary `libraw_processed_image_t` allocation is copied into Go and freed
before the method returns.

## LibRaw-Owned State

These methods operate on LibRaw-owned buffers:

- `Unpack`
- `UnpackThumb`
- `Raw2Image`
- `DcrawProcess`
- `SubtractBlack`
- `FreeImage`
- `Recycle`
- `RecycleDatastream`
- `WritePPMTiff`
- `WriteThumb`

Do not assume internal LibRaw buffers survive `FreeImage`, `Recycle`, or
`Close`. If data must outlive the processor, copy it through the Go-owned result
methods first.

## File Writers

`WritePPMTiff` and `WriteThumb` ask LibRaw to write directly to the target path.
They do not return image bytes. Use `MemImage`, `MemThumb`, `RawImage`, or
`ThumbnailData` when the caller needs bytes in memory.

## Callback Safety

Callback handlers receive Go values only. They do not expose LibRaw pointers or
streams, and they should return quickly because LibRaw is blocked until the
callback returns.

The progress handler may cancel processing by returning a non-zero value.
Panics from callback handlers are recovered at the cgo boundary to avoid
unwinding through C.

## cgo Disabled Builds

The internal bridge has stubs for builds without cgo, but real decoding needs
`CGO_ENABLED=1` and LibRaw development files. Use `make libraw-check` to verify
the active toolchain and linked LibRaw version.

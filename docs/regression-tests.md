# Fixture And Regression Tests

The repository includes small RAW fixtures from Canon, Nikon, Ricoh, and Sony.
They are used to exercise the public wrapper against real LibRaw decoding paths
without downloading large files during CI.

## Targets

```sh
make test-fast      # go test -short ./...
make test-fixtures  # focused real-RAW regression sweep
make test           # race + coverage test suite
make check          # project-level CI equivalent
```

`make test-fixtures` runs the focused fixture regression tests plus the metadata
and maker-note fixture coverage tests. Processing and thumbnail smoke tests are
skipped only when `testing.Short()` is enabled.

## What The Fixture Sweep Covers

- every bundled fixture opens successfully
- camera identity and raw dimensions are populated
- decoder information is available
- all fixtures can unpack, postprocess, and produce non-empty `MemImage` output
- advertised thumbnails can be unpacked and copied with `ThumbnailData`
- invalid file and buffer inputs return typed `libraw.Error` values

The suite intentionally checks structure and non-empty output, not exact color
science or golden image bytes. LibRaw output can differ across library versions,
platform builds, and camera-specific decoder updates.

## Fixture Policy

Do not replace or rewrite existing RAW bytes during normal test maintenance.
New fixtures should be small, redistributable, and documented in
`testdata/README.md`.

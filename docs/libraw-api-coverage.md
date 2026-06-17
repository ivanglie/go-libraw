# LibRaw API Coverage

- Header directory: `testdata/headers/libraw`
- Header version: `0.22.1-Release`
- Total tracked public symbols: `204`
- Explicit coverage entries: `204/204`
- Release gate: `pass` (no unmapped symbols)

## Status Counts

| Status | Count |
| --- | --- |
| `deferred` | `4` |
| `internal` | `3` |
| `unsupported` | `20` |
| `wrapped` | `177` |

## Deferred Symbols

- `function` `libraw_open_wfile`: tracked for a future task
- `function` `libraw_open_wfile_ex`: tracked for a future task
- `struct` `libraw_callbacks_t`: tracked for a future task
- `struct` `libraw_custom_camera_t`: decoder/custom camera work tracked for TASK-012

## Unsupported Symbols

- `function` `libraw_open_file_ex`: removed from default 0.22 build via LIBRAW_NO_IOSTREAMS_DATASTREAM
- `macro` `LIBRAW_CHECK_VERSION`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_COMPILE_CHECK_VERSION`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_COMPILE_CHECK_VERSION_NOTLESS`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_CR3_MEMPOOL`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_FATAL_ERROR`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_IOSPACE_CHECK`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_MAKE_VERSION`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_MEMPOOL_CHECK`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_NO_IOSTREAMS_DATASTREAM`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_OWN_SWAB`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_RUNTIME_CHECK_VERSION_EXACT`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_RUNTIME_CHECK_VERSION_NOTLESS`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_USE_OPENMP`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_USE_STREAMS_DATASTREAM_MAXSIZE`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_VERSION_MAKE`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_WIN32_CALLS`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_WIN32_DLLDEFS`: preprocessor switch or function-like macro not exposed as a Go constant
- `macro` `LIBRAW_WIN32_UNICODEPATHS`: preprocessor switch or function-like macro not exposed as a Go constant
- `version` `LIBRAW_VERSION_TAIL`: non-numeric preprocessor token not exposed as a Go constant

## In-Scope Definition

The release gate covers public C API symbols and public data structures parsed
from the checked-in LibRaw fixture headers. C++-only extension surfaces and
platform/preprocessor-only switches are documented as `unsupported` instead of
being counted as missing Go API.

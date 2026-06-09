# API Coverage Guide

`go-libraw` tracks wrapper coverage against LibRaw fixture headers. The
generated inventory answers the usual question: "I found a LibRaw symbol; where
is the Go equivalent?"

## Where To Look

- [LibRaw API Inventory](libraw-api-inventory.md): generated symbol table for
  functions, enums, macros, structs, and version constants.
- [Metadata Coverage](libraw-metadata-coverage.md): field-level mapping for
  core `libraw_data_t` metadata snapshots.
- [Maker-Notes Coverage](libraw-maker-notes-coverage.md): vendor maker-note
  field coverage and summarized pointer payloads.
- [Output And Raw Params Coverage](libraw-params-coverage.md): mapping for
  `libraw_output_params_t` and `libraw_raw_unpack_params_t`.
- [Sample Parity Examples](examples.md): Go commands corresponding to upstream
  LibRaw sample programs.

## Status Meanings

| Status | Meaning |
| --- | --- |
| `wrapped` | Exposed through the public Go API or generated Go constants. |
| `internal` | Used behind the public API boundary and intentionally not exported. |
| `deferred` | Known upstream symbol tracked for a later workflowr task. |
| `unsupported` | Intentionally not exposed, usually because it is a preprocessor switch, platform-specific entry point, or unavailable in the configured LibRaw build. |
| `unmapped` | Present in headers but missing from the coverage map; this should be fixed or explained before claiming full coverage. |

## Finding A Symbol

1. Search [LibRaw API Inventory](libraw-api-inventory.md) for the LibRaw symbol
   name, for example `libraw_unpack`.
2. Read the status and note column.
3. If the status is `wrapped`, use the named Go API, such as
   `Processor.Unpack`.
4. If the status is `internal`, prefer the public API named by the note.
5. If the status is `deferred`, `unsupported`, or `unmapped`, read the note
   before adding code.

## Regenerating Coverage

Check that generated coverage files are current:

```sh
make check-api-inventory
```

Regenerate after changing the coverage map or fixture headers:

```sh
make api-inventory
```

The inventory is generated from `testdata/headers/libraw` and currently reports
the LibRaw header version shown at the top of
[LibRaw API Inventory](libraw-api-inventory.md).

## Upstream References

LibRaw's own documentation is still the semantic source of truth:

- [LibRaw documentation](https://www.libraw.org/docs)
- [LibRaw C API](https://www.libraw.org/docs/API-C.html)
- [LibRaw API overview](https://www.libraw.org/node/29)

The Go wrapper translates those calls into Go lifecycle, error, and ownership
rules; it does not try to hide that LibRaw is a stateful native decoder.

# Upstream Sync

The offline source of truth is the checked-in LibRaw fixture header directory at
`testdata/headers/libraw`. This keeps normal development reproducible. To audit
against a newer upstream LibRaw release, point the inventory tool at another
header directory.

## Compare Installed Headers

```sh
make api-inventory LIBRAW_HEADERS=/path/to/libraw/include/libraw
make api-coverage LIBRAW_HEADERS=/path/to/libraw/include/libraw
git diff -- docs/libraw-api-inventory.md docs/libraw-api-coverage.md internal/apiinventory/coverage.tsv
```

Added symbols appear in the generated inventory and are added to the coverage
map as `deferred` by `make api-inventory`. Removed symbols disappear from the
generated files. Review the diff, update statuses and notes, then run:

```sh
make release-check
```

## Optional Network Workflow

When network access is available, clone or download
[LibRaw/LibRaw](https://github.com/LibRaw/LibRaw), then use its public header
directory as `LIBRAW_HEADERS`. Network access is not required by
`make release-check`.

## In-Scope Public API

The release coverage gate covers symbols parsed from:

- `libraw.h`
- `libraw_const.h`
- `libraw_types.h`
- `libraw_version.h`

C++-only extension classes, private implementation details, and platform or
preprocessor switches can be marked `unsupported` with an explicit note. Public
C API functions and public data structures should be `wrapped`, `internal`,
`deferred`, or otherwise documented before release.

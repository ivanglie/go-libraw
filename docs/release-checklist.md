# Release Checklist

Use this checklist before tagging the first `go-libraw` release or any release
that claims complete coverage for the tracked LibRaw public C API.

## Required Command

```sh
make release-check
```

The target runs:

- LibRaw linkage verification
- generated API inventory check
- generated API coverage report check
- build, vet, lint, and race tests
- focused fixture regression tests
- sample-parity examples
- cleanup of generated local outputs

## API Coverage Gate

Release readiness is based on the checked-in LibRaw fixture headers and coverage
map:

- [LibRaw API Inventory](libraw-api-inventory.md)
- [LibRaw API Coverage](libraw-api-coverage.md)

The release gate fails if any parsed public symbol is `unmapped`. Symbols marked
`wrapped`, `internal`, `deferred`, or `unsupported` have explicit coverage
decisions and notes. Deferred and unsupported entries must remain visible in the
release notes when they affect user expectations.

## Supported Platforms

Confirm that CI is green for the platforms in the
[Support Matrix](support-matrix.md):

- Ubuntu GitHub-hosted runner
- macOS GitHub-hosted runner

Local release verification should also record:

- `go version`
- `go env GOOS GOARCH CGO_ENABLED`
- `make libraw-check`
- linked LibRaw version

## Documentation

Confirm the following are current:

- [README](../README.md)
- [Versioning Policy](versioning.md)
- [Lifecycle And Processing](lifecycle-processing.md)
- [Memory And Cgo Safety](memory-and-cgo.md)
- [API Coverage Guide](api-coverage.md)
- [Fixture And Regression Tests](regression-tests.md)
- [Sample Parity Examples](examples.md)

## Versioning

Before tagging, choose a standard Go module SemVer tag such as `v0.1.0` and
record the LibRaw baseline separately in the release notes. Do not encode the
LibRaw version in the Git tag.

Release notes should include:

- Go module version
- LibRaw header baseline version
- linked LibRaw runtime version from local verification
- any compatibility notes for supported LibRaw runtime versions

## Licensing And Fixtures

- Keep the project license file in the repository root.
- Do not add large RAW fixtures during release prep.
- Verify any new fixture is redistributable and documented in `testdata/README.md`.

## Upstream Sync

Before release, compare against a freshly downloaded or installed LibRaw header
tree when possible:

```sh
make api-inventory LIBRAW_HEADERS=/path/to/libraw/include/libraw
make api-coverage LIBRAW_HEADERS=/path/to/libraw/include/libraw
git diff -- docs/libraw-api-inventory.md docs/libraw-api-coverage.md internal/apiinventory/coverage.tsv
```

If new symbols appear, classify them in `internal/apiinventory/coverage.tsv`,
regenerate the docs, and rerun `make release-check`.

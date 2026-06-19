# Agent Instructions

## Project

CGo bindings for the [LibRaw](https://www.libraw.org) RAW image processing library.
Linux and macOS only (`amd64`, `arm64`). Requires cgo and LibRaw development headers.

Module: `github.com/ivanglie/go-libraw`  
Public API: `pkg/libraw/`  
CGo bridge: `internal/librawc/`  
Examples: `cmd/`

## Language

All source code, comments, commit messages, and documentation must be in English.

## Writing style

- No "please", "let's", "perhaps".
- No filler phrases.
- Direct: problem → solution.
- Honest feedback.

## Build and test

```sh
make build          # go build ./...
make test           # go test -race ./...
make vet            # go vet ./...
make lint           # golangci-lint run
make examples       # run all sample commands on bundled fixtures
make cover          # coverage of ./internal/... and ./pkg/... (excludes cmd/ and tools/)
make check          # libraw-check + inventory + build + vet + lint + test
```

CGO_ENABLED=1 is required. LibRaw must be installed (`brew install libraw` / `apt-get install libraw-dev`).

Lint with `golangci-lint` **v2.12.2** — the version pinned in `.github/workflows/ci.yml`.
Older local versions miss findings CI enforces (for example `errcheck` in `cmd/`),
so a local green run with a different version is not enough.

## Code layout

- `pkg/libraw/` — public Go API (`Processor`, pipeline methods, params, callbacks, metadata)
- `internal/librawc/` — CGo bridge; each `foo.go` has a `foo_stub.go` with `//go:build !cgo` no-ops
- `cmd/` — standalone example programs, one per upstream LibRaw sample
- `tools/` — repo maintenance tools (not user-installable)
- `testdata/` — CC0 RAW fixtures; `testdata/headers/libraw/` — bundled LibRaw headers for inventory

## Pipeline state machine

`Processor` tracks pipeline progress and rejects out-of-order calls with `ErrInvalidState`.

```
stateInit → (OpenFile/OpenBuffer/OpenBayer) → stateOpened
         → (Unpack)                         → stateUnpacked
         → (Raw2Image)                      → stateImageBuilt
         → (DcrawProcess)                   → stateProcessed
```

`requireState(op, min)` uses `>=` semantics — a state at or beyond `min` passes.
`DcrawProcess` therefore accepts `stateUnpacked` and `stateProcessed` (multirender pattern).
`FreeImage` demotes state back to `stateUnpacked`.
A failed or new `Open*` call resets state to `stateInit`.

## CGo safety rules

- Copy at the boundary: use `unsafe.Slice` + `append([]byte(nil), src...)` or `make`+`copy` — never return slices backed by C memory.
- Check size before casting to `int`: `uint64(img.data_size) > uint64(maxInt)` guard before `int(img.data_size)`.
- Input buffers passed to LibRaw (OpenBuffer, OpenBayer) are copied by LibRaw; the Go slice can be released after the call returns.

## CI vs local LibRaw

CI (`ubuntu-latest`) installs LibRaw from apt — currently **0.21.x**, older than local Homebrew (0.22.x).
Any C symbol added in LibRaw 0.22 must be guarded:

```c
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0, 22, 0)
  // 0.22-only code
#endif
```

Local green is not enough — CI exercises the older version.

## Releases

Each release tag (`v0.x.y`) must state the LibRaw header baseline in the
release description — see `docs/versioning.md`. The MIT license covers only
the Go binding code; distributors linking LibRaw must satisfy LibRaw's
CDDL-1.0 OR LGPL-2.1-or-later — see `THIRD-PARTY-NOTICES.md`.

## API inventory

`docs/libraw-api-inventory.md` tracks every C symbol from `testdata/headers/libraw/`.
Run after any coverage change:

```sh
make check-api-inventory   # verify committed inventory is current
make api-inventory         # regenerate
```

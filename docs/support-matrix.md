# Support Matrix

`go-libraw` is a cgo wrapper around the system LibRaw library. It is a
**Linux- and macOS-only** library by design; Windows is out of scope. Support
depends on Go, cgo, a C/C++ toolchain, and LibRaw development headers being
available on the target machine.

## Continuously Verified

The GitHub Actions workflow runs on pull requests and pushes to `main`.

| Platform | Runner | Architecture | Go | LibRaw source | Checks |
| --- | --- | --- | --- | --- | --- |
| Ubuntu | `ubuntu-latest` | GitHub-hosted `GOARCH` printed in logs | `go.mod` (`1.26`) | `apt` package `libraw-dev` with `pkg-config` | linkage, API inventory, build, vet, lint, race tests, examples, coverage |
| macOS | `macos-latest` | GitHub-hosted `GOARCH` printed in logs | `go.mod` (`1.26`) | Homebrew `libraw` with `pkg-config` | linkage, API inventory, build, vet, lint, race tests, examples, coverage |

Each CI run prints:

- `go version`
- `go env GOOS GOARCH CGO_ENABLED`
- `pkg-config` LibRaw version and flags when available
- the linked LibRaw runtime version from `make libraw-check`

## Supported For Users

| Platform | Architectures | Status | Requirements |
| --- | --- | --- | --- |
| Linux | `amd64`, `arm64` | Supported when LibRaw is discoverable through `pkg-config`. | Go with cgo enabled, C/C++ toolchain, `libraw-dev`/equivalent, `pkg-config`. |
| macOS | `arm64`, `amd64` | Supported with Homebrew LibRaw or `pkg-config`. | Go with cgo enabled, Xcode command line tools, Homebrew `libraw`. |

## Out Of Scope

| Platform | Status | Notes |
| --- | --- | --- |
| Windows | Out of scope. | The library deliberately targets Linux and macOS only; Windows is not built, tested, or advertised. |
| Mobile/WASM | Out of scope. | LibRaw requires native code and filesystem-style integration that is outside the current scope. |

For installation details, see [LibRaw Build Setup](libraw-build.md).

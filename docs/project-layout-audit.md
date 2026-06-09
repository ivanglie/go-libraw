# Project Layout Audit

This audit compares `go-libraw` with the community
[`golang-standards/project-layout`](https://github.com/golang-standards/project-layout)
reference. That repository explicitly describes itself as a collection of common
Go ecosystem layout patterns, not an official Go team standard, and says not
every pattern belongs in every project.

`go-libraw` is a Go library with a small set of support commands, examples,
fixtures, and release tooling. The current layout should stay boring and
import-friendly rather than adopting application/service directories that do
not serve the wrapper.

## Summary

| Area | Current State | Decision |
| --- | --- | --- |
| Root package | Public `libraw` package files live at the module root. | Keep. This is idiomatic for a single-package Go library and avoids a needless `/pkg` import path. |
| `/cmd` | Present for support/example commands. | Keep. Command names match executable intent and `main` packages are small. |
| `/internal` | Present for cgo bridge and API inventory internals. | Keep. It protects non-public implementation code with Go's compiler-enforced `internal` rule. |
| `/docs` | Present for user, API coverage, CI, release, and support docs. | Keep. It matches the layout reference for design and user documentation. |
| `/tools` | Present for checked-in generator tooling. | Keep. It isolates project support tools from the public API. |
| `/testdata` | Present for RAW fixtures and fixture headers. | Keep. Go tooling treats `testdata` specially and it is the right home for package fixtures. |
| `_example` | Present for the minimal library example, plus `cmd/*` sample-parity commands. | Keep for now; document as a deliberate Go-tooling exclusion. Consider a future `/examples` move only if the public docs need it. |
| `/pkg` | Absent. | Keep absent. The module root is already the public library package. |
| `/api`, `/web`, `/configs`, `/init`, `/deployments` | Absent. | Keep absent. These are service/web/deployment patterns and do not apply to this library. |
| `/scripts`, `/build` | Absent. | Keep absent while Makefile and GitHub Actions remain small. Revisit only if build/release logic grows. |
| `/vendor` | Absent. | Keep absent. This is a library using Go modules. |
| `/src` | Absent. | Keep absent. The layout reference explicitly discourages project-level `/src`. |

## Relevant Matches

### `/cmd`

`cmd/libraw-api-inventory`, `cmd/raw-identify`, `cmd/raw-textdump`,
`cmd/mem-image`, and `cmd/thumb-extract` match the `/cmd/<name>` convention.
They are entry points, not reusable library packages. The reusable behavior
stays in the root package or `internal` packages.

### `/internal`

`internal/librawc` contains the private cgo bridge and `internal/apiinventory`
contains release/inventory implementation details. This matches the intent of
`/internal`: code that should not be imported by downstream users.

### `/docs`

The project has user-facing guides, generated API coverage reports, sample
mapping, support matrix, regression test notes, upstream sync notes, and release
readiness docs. This matches the layout reference's `/docs` role.

### `/tools`

`tools/gen-constants` is a project support tool. Keeping it out of the public
package and out of `/cmd` avoids implying it is a user command.

### `/testdata`

RAW fixtures and LibRaw fixture headers live in `testdata`. This is preferred
over a top-level `/test` directory because the fixtures are consumed directly by
Go package tests and Go tooling ignores `testdata` as a package.

## Intentional Deviations

### No `/pkg`

The public package is the module root:

```go
import libraw "github.com/ivanglie/go-libraw"
```

Moving it to `/pkg/libraw` would change the public import path and add nesting
without improving encapsulation. The layout reference also notes that `/pkg` is
not universal and is not needed for smaller projects where extra nesting adds
little value.

### `_example` Instead Of `/examples`

The repository currently uses `_example/main.go` for the smallest runnable
library example. The leading underscore makes Go ignore the directory during
normal `./...` traversal, so the Makefile explicitly build-checks it and
`make examples` runs it.

This is acceptable for now because:

- the example is intentionally tiny and fixture-backed
- `docs/examples.md` maps upstream samples to both `_example` and `cmd/*`
- Makefile targets already make example execution explicit

Recommended follow-up only if desired: create a separate task to evaluate moving
`_example` to `/examples/simple-dcraw` and updating docs/Makefile references.
Do not do that as part of this audit because it is a visible path change.

### No `/test`

The project uses package-level tests and `testdata`. A top-level `/test`
directory would be useful for external black-box apps or larger integration
harnesses, but the current fixture regression suite is clearer next to the
package it validates.

### No `/scripts` Or `/build`

The root `Makefile` and `.github/workflows/ci.yml` are still compact. Splitting
them into `/scripts` or `/build` would add indirection without reducing real
complexity. Revisit this if release packaging, installer generation, or multiple
CI systems are added.

## Intentionally Omitted Application Directories

`go-libraw` is not a web application, service, or deployment artifact. These
directories are intentionally absent:

- `/api`: no OpenAPI, protobuf, or JSON schema surface
- `/web`: no web assets or server-side templates
- `/configs`: no runtime config templates
- `/init`: no system service definitions
- `/deployments`: no Kubernetes, Terraform, Docker Compose, or PaaS manifests
- `/assets`, `/website`: no project website or brand asset bundle
- `/third_party`: no vendored/forked external helper code
- `/vendor`: Go modules and CI dependency installation are sufficient

## Release Readiness Links

The following docs make the current layout auditable for release:

- [README](../README.md)
- [Support Matrix](support-matrix.md)
- [Fixture And Regression Tests](regression-tests.md)
- [API Coverage Guide](api-coverage.md)
- [LibRaw API Coverage](libraw-api-coverage.md)
- [Release Checklist](release-checklist.md)
- [Upstream Sync](upstream-sync.md)

## Recommendation

No restructuring is required before release. The current layout matches the
important Go conventions for a library wrapper: public package at the module
root, private code in `/internal`, command entry points in `/cmd`, support
tooling in `/tools`, docs in `/docs`, and fixtures in `testdata`.

The only arguable mismatch is `_example` versus `/examples`; keep `_example`
until a separate task decides the path change is worth the churn.

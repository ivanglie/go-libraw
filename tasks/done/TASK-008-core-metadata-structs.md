# TASK-008: Wrap Core Metadata Structs

## Goal

Expose complete read access to core LibRaw metadata structs.

## Context

Core metadata lives in `libraw_data_t` fields such as `sizes`, `idata`, `lens`, `shootinginfo`, `color`, `other`, `thumbnail`, `thumbs_list`, `rawdata`, and related nested structs in `libraw_types.h`.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `*.go`
  - `internal/**`
  - `docs/**`
  - `testdata/**`
- Read-only paths:
  - original RAW fixtures.
- Forbidden paths:
  - maker-note vendor-specific structs, except references needed to compile.
- Output artifacts:
  - docs metadata coverage report.
- Default branch: `main`.
- Target branch: `feature/TASK-008-core-metadata-structs`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw and bundled fixtures.

## Scope

- Map core metadata structs to Go types with complete field coverage.
- Add accessors for image dimensions, colors, camera info, lens info, shooting info, GPS/other metadata, thumbnails list, and rawdata summary.
- Add inventory checks for struct field coverage.
- Add fixture tests for Canon, Nikon, Sony, Ricoh, and DNG sample files where available.

## Out Of Scope

- Vendor-specific maker notes.
- Direct mutable access to C-owned metadata.
- Full raw pixel buffer access.

## Acceptance Criteria

- Given each bundled RAW fixture, when opened and metadata is read, then non-empty camera, size, and color metadata are returned where LibRaw provides them.
- Given every core struct field in scope, when coverage is checked, then it is represented in Go or explicitly justified.
- Given metadata is read after `Close`, when an accessor is called, then behavior is documented and safe.

## Test Requirements

- Unit tests: required for C-to-Go conversion helpers.
- Integration tests: required for all bundled RAW fixtures.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: metadata conversion code.
- Coverage metric: line coverage.
- Coverage target: at least `80%`.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `go test ./...`
  - `make check-api-inventory`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Prefer copying metadata into Go-owned values so callers are not exposed to dangling C pointers.

## Clarifications

- Question: Should metadata accessors return snapshots or live views into C data?
- Recommended default: snapshots for safety.
- Answer: return Go-owned snapshots; C-owned pointers are summarized instead of exposed.

## Implementation Outcome

- Added `Processor.Metadata` with Go-owned snapshots for core metadata structs.
- Added coverage for image params, image sizes, color data, lens info, shooting info, GPS/other metadata, thumbnails, and rawdata summary.
- Documented field-level coverage and deferred pointer/vendor-maker-note fields in `docs/libraw-metadata-coverage.md`.
- Added fixture tests for Canon, Nikon, Ricoh DNG, and Sony RAW samples.

## Git And PR

- Branch: `feature/TASK-008-core-metadata-structs`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - metadata
  - api
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: field-by-field mapping is verbose and easy to miss.
- Mitigation: make inventory coverage fail on unmapped fields.

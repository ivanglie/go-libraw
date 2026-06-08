# TASK-004: Wrap Version, Errors, Constants, And Enums

## Goal

Expose LibRaw version helpers, error handling, progress names, constants, and public enum values as idiomatic Go symbols.

## Context

Consumers need stable Go constants and typed errors before higher-level APIs can be pleasant to use. Upstream coverage includes `libraw_version`, `libraw_versionNumber`, `libraw_strerror`, `libraw_strprogress`, macros in `libraw_version.h`, and enums/macros in `libraw_const.h`.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `*.go`
  - `internal/**`
  - `docs/**`
  - `tools/**`
  - `Makefile`
- Read-only paths:
  - `tasks/**`
- Forbidden paths:
  - image processing implementation unrelated to constants/errors.
- Output artifacts:
  - generated constants may live under documented generated paths.
- Default branch: `main`.
- Target branch: `feature/TASK-004-version-errors-constants-enums`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw headers and library.

## Scope

- Wrap `libraw_version`, `libraw_versionNumber`, `libraw_strerror`, and `libraw_strprogress`.
- Add typed Go error values or an error type preserving LibRaw numeric codes.
- Expose public enums from `libraw_const.h` with Go doc comments or generated references.
- Add compile-time or inventory tests proving enum coverage.

## Out Of Scope

- Full lifecycle/image processing.
- Full struct metadata wrappers.

## Acceptance Criteria

- Given a LibRaw error code, when converted to Go error, then the numeric code and LibRaw message are preserved.
- Given a progress enum, when converted to string, then the LibRaw progress text is available.
- Given upstream `libraw_const.h` is inventoried, when coverage is checked, then all public enum constants are mapped or explicitly justified.

## Test Requirements

- Unit tests: required for error formatting and enum conversion.
- Integration tests: required for version and LibRaw string helpers.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: constants/errors package code.
- Coverage metric: line coverage.
- Coverage target: at least `85%` for non-generated code.
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

Generated constants are acceptable. Keep generated files clearly marked and reproducible.

## Clarifications

- Question: Should generated constants use LibRaw C names exactly or Go-style names?
- Recommended default: export Go-style names and preserve C names in comments or mapping metadata.
- Answer:

## Git And PR

- Branch: `feature/TASK-004-version-errors-constants-enums`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - api
  - generated
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: generated names may be awkward or collide.
- Mitigation: add deterministic name normalization and collision tests.


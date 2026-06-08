# TASK-009: Wrap Vendor Maker Notes Metadata

## Goal

Expose complete Go read access to LibRaw maker-note metadata for supported vendors.

## Context

`libraw_makernotes_t` includes vendor-specific structs for Canon, Nikon, Hasselblad, Fuji, Olympus, Sony, Kodak, Panasonic, Pentax, Phase One, Ricoh, Samsung, and common AF metadata.

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
  - processing pipeline behavior unrelated to metadata.
- Output artifacts:
  - docs maker-note coverage report.
- Default branch: `main`.
- Target branch: `feature/TASK-009-maker-notes-metadata`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw and bundled fixtures.

## Scope

- Add Go metadata types for all vendor maker-note structs present in `libraw_types.h`.
- Include nested AF data, crop areas, lens, and high-speed crop structures.
- Add inventory coverage for every maker-note field.
- Add tests using bundled fixtures for vendors represented in `testdata`.

## Out Of Scope

- Semantic interpretation beyond LibRaw's field names and values.
- Adding new RAW fixtures unless needed and approved.

## Acceptance Criteria

- Given Canon, Nikon, Sony, Ricoh, and DNG fixtures, when maker-note metadata is read, then available vendor fields are exposed without panics.
- Given each vendor maker-note struct in upstream headers, when coverage is checked, then every field is mapped or explicitly documented as unavailable.
- Given a fixture from a vendor without data, when the accessor is called, then it returns a zero-value or absent value according to documented API policy.

## Test Requirements

- Unit tests: required for conversion helpers.
- Integration tests: required for represented bundled vendors.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: maker-note conversion code.
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

Keep the Go names close enough to upstream names that users can cross-reference LibRaw documentation.

## Clarifications

- Question: Should absent maker-note groups be pointers, option structs, or zero-value structs?
- Recommended default: use value structs for direct mirrors and document zero-value behavior.
- Answer: use value structs and document zero-value behavior for absent vendor data.

## Implementation Outcome

- Added `Metadata.MakerNotes` with value snapshots for all vendor maker-note groups in `libraw_makernotes_t`.
- Added public aliases for vendor maker-note types.
- Summarized pointer payloads with length/presence fields instead of exposing C-owned pointers.
- Added fixture smoke tests for Canon, Nikon, Ricoh/DNG, and Sony.
- Documented field-level coverage in `docs/libraw-maker-notes-coverage.md`.

## Git And PR

- Branch: `feature/TASK-009-maker-notes-metadata`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - metadata
  - maker-notes
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: some vendor fields are hard to validate with current fixtures.
- Mitigation: use inventory coverage for mapping and fixture tests for smoke behavior.

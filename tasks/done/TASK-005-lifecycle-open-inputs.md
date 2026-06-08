# TASK-005: Wrap Lifecycle And Input Opening APIs

## Goal

Implement complete safe Go coverage for LibRaw handle lifecycle and input opening functions.

## Context

Core LibRaw usage starts with `libraw_init`, `libraw_open_file`, `libraw_open_file_ex`, `libraw_open_buffer`, `libraw_open_bayer`, `libraw_recycle_datastream`, `libraw_recycle`, and `libraw_close`.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `*.go`
  - `internal/**`
  - `_example/**`
  - `docs/**`
  - `testdata/**`
- Read-only paths:
  - original RAW fixture bytes.
- Forbidden paths:
  - callback and processing implementation except stubs needed for tests.
- Output artifacts:
  - temporary files under `tmp/tests/`.
- Default branch: `main`.
- Target branch: `feature/TASK-005-lifecycle-open-inputs`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw library and bundled RAW fixtures.

## Scope

- Add constructor options for LibRaw init flags.
- Add file open, file open with buffer size, memory buffer open, and Bayer buffer open APIs.
- Define ownership rules for Go byte slices passed to C.
- Add recycle and close semantics with clear post-close behavior.

## Out Of Scope

- Windows wide-file APIs unless the project already supports Windows in cgo.
- Full custom datastream subclassing.

## Acceptance Criteria

- Given a valid RAW file in `testdata`, when opened by path, then LibRaw metadata becomes available without leaks.
- Given valid RAW bytes, when opened from memory, then LibRaw can unpack or identify it according to later available helpers.
- Given an invalid path or invalid buffer, when opened, then a typed LibRaw error is returned.
- Given a handle is recycled, when another file is opened, then the second file does not retain stale metadata.

## Test Requirements

- Unit tests: required for lifecycle state transitions.
- Integration tests: required for file and buffer opening using bundled fixtures.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: lifecycle and input code.
- Coverage metric: line coverage.
- Coverage target: at least `80%` for non-cgo Go logic.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `go test ./...`
  - `go test -race ./...`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Document that LibRaw handles are not safe for concurrent mutation unless proven otherwise. Use finalizers only if they do not obscure explicit `Close`.

## Clarifications

- Question: Should `OpenWFile` be supported in initial lifecycle coverage?
- Recommended default: defer to a Windows-specific task unless current platform work proves straightforward.
- Answer: Deferred to a Windows-specific task. `libraw_open_wfile`/`libraw_open_wfile_ex` are Win32-only (`#if defined(_WIN32)`) and not exercised on the current darwin/linux cgo targets.

## Implementation Outcome

- Wrapped `libraw_open_file`, `libraw_open_buffer`, `libraw_open_bayer`, `libraw_recycle`, and `libraw_recycle_datastream` as `Processor.OpenFile`, `OpenBuffer`, `OpenBayer`, `Recycle`, and `RecycleDatastream`.
- `libraw_open_file_ex` is unavailable in the baseline LibRaw 0.22 build: `libraw.h` defines `LIBRAW_NO_IOSTREAMS_DATASTREAM` by default, so the symbol is compiled out and absent from the shared library. Marked `unsupported` in the API coverage map rather than wrapped.
- Buffer ownership: `OpenBuffer`/`OpenBayer` copy input bytes into C memory retained by the handle (LibRaw keeps a pointer rather than copying). The copy is freed on `Recycle`, `Close`, or reuse, keeping ownership on the C side and satisfying cgo pointer rules.

## Git And PR

- Branch: `feature/TASK-005-lifecycle-open-inputs`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - api
  - cgo
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: memory ownership bugs across Go and C.
- Mitigation: copy or pin buffers deliberately and test invalid/post-close behavior.


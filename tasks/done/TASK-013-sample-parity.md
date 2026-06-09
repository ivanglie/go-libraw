# TASK-013: Add LibRaw Sample Parity Examples

## Goal

Provide Go examples that correspond to important upstream LibRaw samples and demonstrate wrapper completeness.

## Context

Upstream samples include `simple_dcraw.cpp`, `dcraw_emu.cpp`, `raw-identify.cpp`, `rawtextdump.cpp`, `mem_image_sample.cpp`, `unprocessed_raw.cpp`, `4channels.cpp`, `openbayer_sample.cpp`, and multithreaded examples.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `_example/**`
  - `cmd/**`
  - `docs/**`
  - `testdata/**`
  - `Makefile`
  - `*.go`
- Read-only paths:
  - original RAW fixtures.
- Forbidden paths:
  - internal API redesign not required by examples.
- Output artifacts:
  - example output files under `tmp/examples/`.
- Default branch: `main`.
- Target branch: `feature/TASK-013-sample-parity`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw and fixtures.

## Scope

- Add Go examples or small commands equivalent to raw identify, simple dcraw, memory image extraction, raw text dump summary, and thumbnail extraction.
- Document which upstream sample each Go example maps to.
- Add Makefile targets for running examples without committing generated outputs.

## Out Of Scope

- Perfect CLI flag parity with `dcraw_emu`.
- Benchmark suite.

## Acceptance Criteria

- Given bundled fixtures, when example targets run, then identify, metadata dump, memory image, and thumbnail examples complete successfully.
- Given docs are read, when a user wants an upstream sample equivalent, then they can find the corresponding Go example.
- Given output artifacts are generated, when `make clean` runs, then generated example outputs are removed.

## Test Requirements

- Unit tests: not required unless helper code is introduced.
- Integration tests: examples must compile.
- E2E tests: required for documented Makefile example targets.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: not applicable.
- Coverage metric: not applicable.
- Coverage target: not applicable.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `go test ./...`
  - `make examples`
  - `make clean`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Keep examples small and readable. Do not hide core API usage behind too much helper code.

## Clarifications

- Question: Should `dcraw_emu` be implemented as a full CLI clone?
- Recommended default: defer full `dcraw_emu` parity; include representative options only.
- Answer: defer full `dcraw_emu` parity. This task adds readable sample-parity examples for the key wrapper flows instead of a complete flag-compatible clone.

## Implementation Outcome

- Added small Go commands for raw identification, metadata text dump, in-memory image extraction, and thumbnail extraction.
- Kept `_example` as the simple dcraw-style library usage example.
- Documented the mapping from upstream LibRaw samples to Go examples in `docs/examples.md`.
- Added `make examples` output under `tmp/examples/` and kept generated files removable with `make clean`.

## Git And PR

- Branch: `feature/TASK-013-sample-parity`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - examples
  - documentation
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: examples may become fragile if they assert exact image bytes.
- Mitigation: assert structural outputs first, with optional golden tests only where stable.

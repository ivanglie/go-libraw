# TASK-002: Add LibRaw Build Discovery

## Goal

Make `go-libraw` reliably discover and link LibRaw across common developer environments.

## Context

The wrapper depends on upstream C/C++ LibRaw. Developers need predictable build behavior via `pkg-config`, common Homebrew/Linux include paths, and clear diagnostics.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `internal/**`
  - `docs/**`
  - `Makefile`
  - `go.mod`
  - `go.sum`
  - `.github/**`
- Read-only paths:
  - `tasks/**`
- Forbidden paths:
  - unrelated package APIs.
- Output artifacts:
  - build diagnostics may be written to `tmp/build/`.
- Default branch: `main`.
- Target branch: `feature/TASK-002-libraw-build-discovery`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: macOS and Linux documentation; at least one local platform verified.

## Scope

- Add cgo flags using `pkg-config: libraw` where possible.
- Document installation for macOS Homebrew, Debian/Ubuntu, Fedora, and source builds.
- Add a build check target to the Makefile.
- Add tests or scripts that verify the linked LibRaw version can be read.

## Out Of Scope

- Windows support.
- Static vendored LibRaw builds.
- Cross-compilation.

## Acceptance Criteria

- Given LibRaw is installed with pkg-config metadata, when `go test ./...` runs, then cgo links without manual flags.
- Given LibRaw is missing, when a developer runs the documented check, then the output identifies the missing dependency.
- Given docs are read by a new contributor, when they follow their platform section, then they know which packages to install.

## Test Requirements

- Unit tests: not required.
- Integration tests: required for linked version check.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: not applicable.
- Coverage metric: not applicable.
- Coverage target: not applicable.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `make build`
  - `go test ./...`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Keep build discovery simple first. Prefer standard cgo `#cgo pkg-config: libraw` over bespoke scripts unless platform evidence requires otherwise.

## Clarifications

- Question: Should Windows be supported in the initial wrapper?
- Recommended default: defer Windows until the Unix cgo path is stable.
- Answer:

## Git And PR

- Branch: `feature/TASK-002-libraw-build-discovery`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - build
  - cgo
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: LibRaw package names differ by distro.
- Mitigation: document verified package names and keep detection diagnostics explicit.


# TASK-014: Add Cross-Platform CI And Dependency Matrix

## Goal

Verify `go-libraw` on supported platforms and document the support matrix.

## Context

A cgo wrapper can fail differently across macOS, Linux distributions, architectures, and LibRaw versions. CI should catch build and test regressions early.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `.github/**`
  - `docs/**`
  - `Makefile`
  - `go.mod`
  - `go.sum`
  - `internal/**`
- Read-only paths:
  - `tasks/**`
- Forbidden paths:
  - public API changes unless required by CI build fixes.
- Output artifacts:
  - CI logs only.
- Default branch: `main`.
- Target branch: `feature/TASK-014-cross-platform-ci`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: GitHub Actions or equivalent CI.

## Scope

- Add CI for Go test, race where practical, lint, and API inventory check.
- Install LibRaw dependencies on macOS and Ubuntu runners.
- Document supported OS, architecture, Go version, and LibRaw version assumptions.
- Add CI caching only where simple and low-risk.

## Out Of Scope

- Windows CI unless earlier tasks established Windows support.
- Publishing releases.

## Acceptance Criteria

- Given a pull request runs CI, when dependencies install successfully, then build, tests, lint, and inventory checks run.
- Given CI fails due to missing LibRaw, when logs are inspected, then the missing package is obvious.
- Given docs are read, when a user checks support status, then supported and unsupported platforms are explicit.

## Test Requirements

- Unit tests: existing suite must pass in CI.
- Integration tests: fixture tests must pass in CI.
- E2E tests: examples may run if runtime is acceptable.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: full module.
- Coverage metric: line coverage where Go supports it.
- Coverage target: report only unless a stricter threshold already exists.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `make check`
  - `make check-api-inventory`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Keep the first CI matrix modest: macOS latest and Ubuntu latest are enough if both pass.

## Clarifications

- Question: Should CI require lint if `golangci-lint` is unavailable or slow?
- Recommended default: yes, install a pinned golangci-lint version in CI.
- Answer: yes. CI uses the official `golangci/golangci-lint-action` with pinned `golangci-lint` version `v2.12.2`.

## Implementation Outcome

- Expanded GitHub Actions CI to run on `ubuntu-latest` and `macos-latest`.
- Added OS-specific LibRaw installation, API inventory check, lint, examples, coverage, and version/discovery logging.
- Added a support matrix documenting CI coverage, user-supported platforms, and unsupported platforms.

## Git And PR

- Branch: `feature/TASK-014-cross-platform-ci`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - ci
  - build
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: CI LibRaw version may differ from developer machines.
- Mitigation: print LibRaw runtime and compile-time versions in CI logs.

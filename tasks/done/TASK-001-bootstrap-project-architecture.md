# TASK-001: Bootstrap Project Architecture

## Goal

Create the initial package architecture for `go-libraw`, including cgo boundaries, internal package layout, documentation skeleton, and a smoke-tested public `libraw` package.

## Context

The repository currently contains only `go.mod`, a Makefile, an example stub, and RAW fixtures. The full wrapper needs a stable architecture before individual LibRaw API surfaces are added.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists; otherwise current workspace root.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `*.go`
  - `internal/**`
  - `cmd/**`
  - `_example/**`
  - `docs/**`
  - `go.mod`
  - `go.sum`
  - `Makefile`
  - `.golangci.yml`
- Read-only paths:
  - `testdata/**` except generated outputs from examples.
- Forbidden paths:
  - `.workflowr/**`
  - `tasks/**`
- Output artifacts:
  - temporary coverage files may be written to `tmp/coverage/`.
- Default branch: `main`.
- Target branch: `feature/TASK-001-bootstrap-project-architecture`.
- Package manager: Go modules.
- Runtime: Go 1.26 as declared by `go.mod`, or the closest installed Go version if 1.26 is unavailable.
- Required environment: a local LibRaw development installation or a documented skip path for cgo smoke tests.

## Scope

- Establish public package naming and exported root types.
- Add an internal cgo bridge package boundary.
- Add a minimal `Processor` or equivalent handle type with lifecycle-safe construction and close semantics.
- Add package documentation explaining ownership, cgo, and LibRaw dependency expectations.
- Update `_example/main.go` to compile against the new public package without performing full image processing yet.

## Out Of Scope

- Full LibRaw API coverage.
- Generated bindings.
- Image processing output parity.
- Vendoring LibRaw source.

## Acceptance Criteria

- Given a consumer imports `github.com/ivanglie/go-libraw`, when `go test ./...` runs with LibRaw available, then the package builds and the smoke test passes.
- Given a handle is closed, when `Close` is called again, then it is safe and deterministic.
- Given LibRaw is not available, when checks are run, then the failure clearly explains the missing dependency or skips only tests that require LibRaw.

## Test Requirements

- Unit tests: required for Go lifecycle behavior.
- Integration tests: cgo smoke test required when LibRaw is installed.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: newly added Go packages.
- Coverage metric: line coverage.
- Coverage target: at least `70%` for newly added non-cgo Go logic.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `go test ./...`
  - `go test ./... -cover`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Prefer a small public API and keep direct cgo calls in `internal/librawc` or an equivalent internal package. Use build tags only if needed to make missing LibRaw diagnostics clean.

## Clarifications

- Question: Should LibRaw be dynamically linked from the system or vendored into the module?
- Recommended default: dynamically link first; create a later task for optional vendoring/static builds.
- Answer:

## Git And PR

- Branch: `feature/TASK-001-bootstrap-project-architecture`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - golang
  - cgo
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: local environments may not have LibRaw headers installed.
- Mitigation: provide explicit dependency detection and actionable errors.

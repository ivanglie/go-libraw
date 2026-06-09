# TASK-015: Write Complete API Documentation

## Goal

Document installation, API concepts, lifecycle, processing flow, metadata coverage, memory ownership, examples, and upstream coverage status.

## Context

A 100% wrapper needs documentation that explains not only how to call methods, but also what is intentionally low-level, where LibRaw owns memory, and how coverage is validated.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `README.md`
  - `docs/**`
  - `*.go`
  - `_example/**`
- Read-only paths:
  - `tasks/**`
- Forbidden paths:
  - behavioral implementation changes except doc examples that must compile.
- Output artifacts:
  - generated docs coverage report if produced by existing tooling.
- Default branch: `main`.
- Target branch: `feature/TASK-015-api-documentation`.
- Package manager: Go modules.
- Runtime: Go.
- Required environment: no external services beyond LibRaw for example verification.

## Scope

- Add README quick start and dependency installation.
- Add lifecycle and processing guide.
- Add metadata and maker-note coverage guide.
- Add memory ownership and cgo safety guide.
- Add upstream coverage table linked to generated inventory.
- Ensure doc examples compile where feasible.

## Out Of Scope

- Marketing site.
- Generated godoc hosting.

## Acceptance Criteria

- Given a new user has LibRaw installed, when they follow README quick start, then they can open and process a RAW file.
- Given a user needs a specific LibRaw symbol, when they read coverage docs, then they can see its Go equivalent or documented exclusion.
- Given a user reads memory docs, when handling images or thumbnails, then ownership and lifetime are clear.

## Test Requirements

- Unit tests: doc examples should compile if written as Go examples.
- Integration tests: quick start example should be covered by existing example tests where practical.
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
  - `go test ./...`
  - `make check-api-inventory`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Prefer concise docs with links to upstream LibRaw documentation. Keep generated coverage docs separate from hand-written guides.

## Clarifications

- Question: Should docs include Russian text?
- Recommended default: English public docs, because Go package docs and upstream LibRaw are English.
- Answer: use English public docs. User-facing task discussion can stay in Russian, but package README and docs should match Go and LibRaw documentation conventions.

## Implementation Outcome

- Added a README with install commands, quick start, examples, and coverage links.
- Added lifecycle/processing, memory/cgo safety, and API coverage guides.
- Added a compile-only Go example for the README quick-start flow.
- Linked handwritten docs to generated LibRaw inventory and coverage reports.

## Git And PR

- Branch: `feature/TASK-015-api-documentation`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - documentation
  - api-coverage
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: docs drift from generated inventory.
- Mitigation: link docs to generated coverage reports and verify examples in tests.

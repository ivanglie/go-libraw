# TASK-017: Add Upstream Sync And Release Readiness Checks

## Goal

Create maintenance checks that keep `go-libraw` aligned with upstream LibRaw and make the first release auditable.

## Context

LibRaw evolves. A 100% wrapper should detect upstream changes, refresh the inventory, and clearly report whether the Go wrapper is still complete.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `cmd/**`
  - `tools/**`
  - `docs/**`
  - `.github/**`
  - `Makefile`
  - `README.md`
- Read-only paths:
  - source wrapper APIs unless small fixes are required by release checks.
- Forbidden paths:
  - unrelated feature work.
- Output artifacts:
  - `docs/release-checklist.md`
  - `docs/libraw-api-coverage.md`
- Default branch: `main`.
- Target branch: `feature/TASK-017-upstream-sync-and-release-readiness`.
- Package manager: Go modules.
- Runtime: Go.
- Required environment: network access for optional upstream version comparison.

## Scope

- Add a command or documented process to compare local coverage with latest upstream LibRaw.
- Add release checklist covering supported platforms, LibRaw versions, API coverage, tests, docs, licenses, and examples.
- Add CI or Makefile target that fails when API coverage is below 100% for in-scope public C API.
- Document exclusions for private/internal LibRaw APIs.

## Out Of Scope

- Publishing tags or GitHub releases.
- Automatically opening upstream sync PRs unless CI infrastructure already supports it.

## Acceptance Criteria

- Given upstream headers change, when sync check is run, then added/removed/changed symbols are reported.
- Given release readiness is checked, when any in-scope symbol is unmapped, then the release checklist shows the wrapper is not 100% complete.
- Given all tasks are complete, when `make release-check` runs, then API coverage, tests, examples, and docs checks pass.

## Test Requirements

- Unit tests: required for sync diff logic.
- Integration tests: optional network-backed latest-upstream check; must be skippable offline.
- E2E tests: release-check target required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: sync tooling.
- Coverage metric: line coverage.
- Coverage target: at least `80%`.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `go test ./...`
  - `make release-check`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

The source of truth should be checked-in inventory plus a documented upstream version. Network checks should enrich the report, not make offline development impossible.

## Clarifications

- Question: Should "100%" include C++-only classes such as custom datastream subclasses?
- Recommended default: include public C API and data structures first; document C++-only surfaces as explicit exclusions unless a later design chooses to support them.
- Answer: define release coverage around parsed public C API symbols and public data structures from the tracked LibRaw headers. C++-only extension surfaces and platform/preprocessor-only switches are explicit documented exclusions.

## Implementation Outcome

- Added generated API coverage summary report support to the inventory tool.
- Added `make api-coverage`, `make check-api-coverage`, and `make release-check`.
- Added release checklist and upstream sync documentation.
- Linked release/upstream docs from the README.

## Git And PR

- Branch: `feature/TASK-017-upstream-sync-and-release-readiness`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - release
  - api-coverage
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: "100%" can be interpreted differently.
- Mitigation: define in-scope public API explicitly and make exclusions visible in release docs.

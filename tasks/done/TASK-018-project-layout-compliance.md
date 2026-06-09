# TASK-018: Audit Go Project Layout Compliance

## Goal

Check `go-libraw` against the common layout patterns described by `golang-standards/project-layout` and document intentional matches, gaps, and deviations.

## Context

The repository now has a stable package, commands, internal tooling, docs, examples, CI, fixtures, and release checks. Before release, it should be clear whether the structure follows common Go project conventions and where this library intentionally differs.

Reference: https://github.com/golang-standards/project-layout

Important nuance: `golang-standards/project-layout` is a community layout reference, not an official Go team standard. Treat it as an audit checklist, not as a mandatory restructuring rulebook.

## Execution Context

- Workspace: current workspace.
- Repository: current repository if one exists.
- Project root: current workspace root.
- Create new project if missing: no.
- Allowed paths:
  - `README.md`
  - `docs/**`
  - `Makefile`
  - `.github/**`
  - `cmd/**`
  - `internal/**`
  - `_example/**`
  - `testdata/**`
- Read-only paths:
  - public wrapper API files unless a tiny layout-related comment/doc fix is required.
  - original RAW fixture bytes.
- Forbidden paths:
  - broad package restructuring without explicit follow-up task approval.
  - behavioral implementation changes unrelated to layout/readiness.
- Output artifacts:
  - `docs/project-layout-audit.md`
- Default branch: `main`.
- Target branch: `feature/TASK-018-project-layout-compliance`
- Package manager: Go modules.
- Runtime: Go with cgo enabled for full checks.
- Required environment: LibRaw and bundled fixtures for project checks.

## Scope

- Compare current repository directories against relevant `project-layout` sections such as `/cmd`, `/internal`, `/pkg`, `/api`, `/web`, `/configs`, `/scripts`, `/build`, `/deployments`, `/test`, `/docs`, `/tools`, `/examples`, and `/testdata`.
- Document which directories are present and why.
- Document directories intentionally omitted because `go-libraw` is a library, not a service or web application.
- Verify `_example` versus common `/examples` guidance and recommend whether to keep, rename, or document the current convention.
- Check README/docs links so users can find commands, examples, fixtures, CI, release checks, and support matrix.
- Add small Makefile/docs changes only if they clarify layout or release readiness.

## Out Of Scope

- Moving public package files into `/pkg`.
- Moving RAW fixtures or changing fixture bytes.
- Creating service/web/deployment scaffolding that the library does not need.
- Enforcing `project-layout` as an official Go standard.

## Acceptance Criteria

- Given a maintainer reads `docs/project-layout-audit.md`, when they compare the repository to `golang-standards/project-layout`, then they can see each relevant directory decision and rationale.
- Given the audit finds a missing high-value convention, when the change is low risk, then docs/Makefile links are updated or a follow-up task is recommended.
- Given the project is a Go library wrapper, when irrelevant service/app directories are absent, then the audit documents them as intentionally omitted.
- Given checks run, when layout documentation changes are complete, then core build/test/release commands still pass.

## Test Requirements

- Unit tests: not required unless Go code changes are introduced.
- Integration tests: existing suite must pass if layout or Makefile changes are made.
- E2E tests: `make release-check` should pass when feasible.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required if Go files change.
- Coverage scope: not applicable.
- Coverage metric: not applicable.
- Coverage target: not applicable.
- Required commands:
  - `make check`
  - `make release-check`
  - `git diff --check`

## Language Requirements

- Language: Markdown for audit documentation; Go only if compile-checked examples or tooling are adjusted.
- Style: concise, factual, and explicit about intentional deviations.
- Dependency policy: do not add dependencies.
- Standard tools: project Makefile targets and Go tooling.

## Implementation Notes

Prefer an audit document over mechanical restructuring. The current repository already uses several common conventions: root package files for the library API, `/cmd` for command binaries, `/internal` for private tooling/cgo bridge, `/docs` for guides, `/testdata` for fixtures, and `_example` for a minimal example excluded from normal package discovery.

If `_example` is judged confusing compared with `/examples`, recommend a follow-up task rather than moving it automatically, because existing docs and Makefile targets currently reference it.

## Clarifications

- Question: Should this task force the repository to exactly match `golang-standards/project-layout`?
- Recommended default: no. Produce a compliance audit and only make low-risk documentation or Makefile improvements.
- Answer: no. This task audits against the community reference, documents intentional matches and deviations, and avoids broad restructuring.

## Implementation Outcome

- Added `docs/project-layout-audit.md` comparing the current repository layout with relevant `golang-standards/project-layout` directories.
- Documented intentional omissions for service/web/deployment directories that do not apply to a Go library wrapper.
- Documented why the root package, `/cmd`, `/internal`, `/docs`, `/tools`, and `testdata` should remain as-is.
- Linked the audit from the README documentation list.

## Git And PR

- Branch: `feature/TASK-018-project-layout-compliance`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - documentation
  - maintenance
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: blindly following a generic layout could make a small Go library more complex.
- Mitigation: document intentional omissions and avoid restructuring unless the benefit is clear and separately approved.

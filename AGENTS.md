# Workflowr Development Protocol

Use this protocol for every implementation task unless the task file says otherwise.

## Default Flow

1. Read the task file and any referenced docs.
2. Inspect the repository before deciding on an implementation.
3. Run the git safety preflight when the project is in a git repository.
4. Produce a short plan with scope, risks, and verification steps.
5. Implement the smallest change that satisfies the acceptance criteria.
6. Add or update tests according to the task requirements.
7. Run required checks: lint, typecheck, tests, coverage, or the closest available equivalents.
8. Review the diff for regressions, unrelated changes, and missing tests.
9. Commit changes using the configured commit style.
10. Prepare a pull request description from `.workflowr/pr-template.md`.

## Operating Rules

- Preserve unrelated user changes.
- Prefer existing project patterns over new abstractions.
- Keep changes scoped to the task.
- Do not skip tests unless the task explicitly permits it or the user approves.
- If coverage requirements cannot be measured, state why and provide the closest evidence.
- If external access or credentials are required, ask before proceeding.

## Execution Context

Before implementation, identify:

1. Workspace root.
2. Repository root.
3. Project root.
4. Allowed edit scope.
5. Forbidden paths.
6. Required commands and working directories.

If the task does not define a project root:

- Use the current repository root when a repository exists.
- Use the current workspace root only for new-project tasks.
- Do not create a new project unless the task explicitly says so.
- If multiple candidate projects exist and the target is unclear, ask a clarification question.

If the task defines allowed or forbidden paths:

- Treat allowed paths as the maximum edit scope.
- Treat forbidden paths as read-only unless the user explicitly approves changes.
- Ask before editing outside the declared scope.

If the task defines output artifacts:

- Write generated reports, coverage files, screenshots, or build artifacts only to the declared output paths.
- Do not commit generated artifacts unless the task explicitly requires them.

## Git Safety

Before editing in a git repository:

1. Run `git status --short --branch`.
2. Identify the current branch and compare it with the task target branch.
3. Identify existing uncommitted changes.
4. Treat pre-existing changes as user-owned unless the task states otherwise.
5. Do not overwrite, revert, reformat, stage, or commit user-owned changes.

Before committing:

1. Review `git diff` for all changed files.
2. Stage only files changed for the current task.
3. Verify the commit includes no secrets, debug artifacts, local paths, or unrelated generated files.
4. Use the task commit style.
5. If the working tree contains unrelated changes, mention them separately and leave them unstaged.

Branch rules:

- Use the task target branch when provided.
- If the target branch does not exist, create it from the default branch when it is safe to do so.
- If the current branch has uncommitted user-owned changes and switching branches would be risky, ask for confirmation.
- Do not force push, reset hard, clean, or delete branches unless explicitly requested.

## Coverage Policy

When coverage is required:

- Identify the coverage scope before implementation.
- Prefer task-defined scope over project defaults.
- If no scope is provided, report coverage for changed packages or changed modules.
- State the metric used, such as line, branch, function, or statement coverage.
- If the tool cannot measure the requested metric, use the closest available metric and explain the gap.
- Coverage must meet the required threshold before delivery unless the task explicitly allows an exception.

## Lint And Static Analysis Policy

When linting, formatting, or static analysis is required:

- Prefer project-defined commands over generic defaults.
- If no project-defined command exists, use the closest standard tool for the language and ecosystem.
- Run formatting before linting when the ecosystem separates those steps.
- Treat formatter, linter, and static analysis failures as required check failures.
- Do not introduce a new linter, formatter, or config unless the task explicitly asks for it.
- If the required tool is missing and cannot be installed without external access, block the task or ask for approval.

## Language Conventions

Follow idiomatic conventions for the language used by the changed code.

For Go code:

- Prefer idiomatic Go over clever abstractions.
- Prefer the Go standard library before adding external dependencies.
- Do not add external dependencies unless they are already used by the project or explicitly required by the task.
- Keep package APIs small and cohesive.
- Use `gofmt` for formatting.
- Use `go vet` for standard static analysis unless the project defines a stricter command.
- Keep errors explicit and useful.
- Prefer table-driven tests when they improve clarity.

## Failure Policy

Required checks must pass before delivery.

If a required check fails:

1. Try to fix the failure if it is in scope.
2. If the failure is unrelated to the task, document evidence and keep the task changes isolated.
3. If the failure cannot be resolved, mark the task as blocked instead of done.
4. Do not open a PR or claim completion with failed required checks unless the task or user explicitly allows degraded delivery.

## Task Lifecycle

Task files move through these status directories:

- `tasks/inbox/`: tasks waiting to be started.
- `tasks/in-progress/`: tasks currently being planned or implemented.
- `tasks/blocked/`: tasks that cannot continue without user input, external access, failing dependency resolution, or an explicit exception.
- `tasks/done/`: tasks that satisfy the Definition Of Done.

Lifecycle rules:

- Move a task from `tasks/inbox/` to `tasks/in-progress/` when implementation starts.
- Keep the same task filename when moving between status directories.
- Move a task to `tasks/blocked/` only after writing a blocker note.
- Move a task to `tasks/done/` only after the Definition Of Done is satisfied.
- Do not delete task files during normal workflow.
- If the agent cannot move files, state the intended status transition in the final report.

When blocking a task:

1. Create a blocker note from `.workflowr/blocker-note.md`.
2. Place it next to the blocked task or in the task file under `## Blocker Note`.
3. Include unknowns, evidence, attempted options, and the recommended next step.

## Queue Continuation

When a task completes successfully and its PR is merged:

1. Move the completed task to `tasks/done/`.
2. Check `tasks/inbox/` for the next task.
3. If a next task exists, start it by following the normal Task Intake flow.
4. If multiple tasks exist, choose the lexicographically first filename unless the user or task metadata defines priority.
5. If no next task exists, report that the queue is empty.
6. Do not start the next task if the completed task is not merged, is blocked, or has failed required checks.

## Task Intake

Task files should live in `tasks/inbox/` until started.

When asked to implement a task:

1. Move conceptually from intake to planning.
2. If the task is ambiguous but has a safe interpretation, proceed with that interpretation and record the assumption.
3. If ambiguity could cause wasted work or risky behavior, ask one concise question.

## Clarification Flow

If the task is unclear:

1. Read the repository context before asking questions.
2. Prefer making a documented assumption when the risk is low.
3. Ask only when ambiguity can change implementation, data model, API, UX, security, tests, or delivery scope.
4. Ask 1-3 specific questions at most.
5. Include a recommended default for each question when possible.
6. Do not implement risky assumptions without confirmation.
7. If blocked, write a short blocker note with unknowns, options considered, and the recommended next step.

## Git And PR

- Branch naming default: `feature/<task-id-slug>`.
- Commit style default: Conventional Commits.
- One commit is preferred for small tasks; multiple commits are fine when they map to clear implementation phases.
- PRs should include summary, testing, coverage result, risks, and review notes.

## PR Creation Policy

Use this policy when the task asks the agent to open a PR, not only prepare a PR draft.

Defaults:

- Provider: GitHub.
- Tool: GitHub CLI (`gh`).
- PR mode: draft.
- Base branch: task PR target, otherwise default branch.
- Head branch: task target branch.
- Title format: `<TASK-ID>: <short summary>`.
- Body template: `.workflowr/pr-template.md`.
- Push branch automatically: yes, only after required checks pass.

Rules:

- Do not create or push a PR until required checks and coverage pass.
- Do not create a PR if the task is blocked.
- Do not merge a PR until all required checks pass successfully.
- Treat pending, skipped, cancelled, or failed required checks as not mergeable unless the repository policy explicitly allows them.
- If `gh` is missing, not authenticated, or the remote is not GitHub, prepare the PR body locally and report the exact command the user can run.
- If labels, reviewers, or assignees are defined by the task, include them in the `gh pr create` command when supported.
- If no reviewers or labels are defined, omit them.
- Never force push unless explicitly requested.
- Prefer draft PRs unless the task explicitly requests a ready-for-review PR.

## Definition Of Done

A task is done when:

- Acceptance criteria are satisfied.
- Required tests/checks have passed.
- Coverage requirements are met.
- The diff has been reviewed.
- Any PR intended for merge has all required checks passing.
- Commits and PR draft are ready, or the PR has been opened when requested.

If any required check or coverage target cannot be satisfied, the task is blocked unless an explicit exception is granted.

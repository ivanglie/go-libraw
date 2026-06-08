# TASK-012: Wrap Camera, Capability, Decoder, And Helper APIs

## Goal

Expose remaining public helper functions for camera lists, capabilities, decoder metadata, dimensions, color matrices, and raw inset crop adjustment.

## Context

Public helper APIs include `libraw_cameraList`, `libraw_cameraCount`, `libraw_capabilities`, `libraw_unpack_function_name`, `libraw_get_decoder_info`, `libraw_get_raw_height`, `libraw_get_raw_width`, `libraw_get_iheight`, `libraw_get_iwidth`, `libraw_get_cam_mul`, `libraw_get_pre_mul`, `libraw_get_rgb_cam`, `libraw_get_color_maximum`, `libraw_get_iparams`, `libraw_get_lensinfo`, `libraw_get_imgother`, and `libraw_adjust_to_raw_inset_crop`.

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
  - callback implementation.
- Output artifacts:
  - none.
- Default branch: `main`.
- Target branch: `feature/TASK-012-camera-capabilities-decoder-helpers`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw and bundled fixtures.

## Scope

- Add camera list and count APIs.
- Add runtime capabilities API.
- Add decoder info and unpack function name APIs.
- Add dimension, color multiplier, RGB camera matrix, maximum color, and inset crop helper methods.
- Ensure inventory coverage marks these public helpers complete.

## Out Of Scope

- Core metadata struct mapping already handled by metadata tasks.
- Processing pipeline behavior beyond setup needed for helper calls.

## Acceptance Criteria

- Given LibRaw is linked, when camera list is requested, then a non-empty Go slice is returned and its length matches camera count.
- Given a fixture is opened, when decoder info and unpack function name are requested, then the wrapper returns LibRaw-provided values.
- Given color helper indexes are invalid, when called, then documented validation behavior occurs.

## Test Requirements

- Unit tests: required for index validation and slice conversion.
- Integration tests: required for camera list and fixture helper calls.
- E2E tests: not required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: helper wrappers.
- Coverage metric: line coverage.
- Coverage target: at least `85%` for non-generated Go code.
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

Prefer returning copies of C string arrays. Keep helper methods grouped clearly in docs.

## Clarifications

- Question: Should camera list be cached?
- Recommended default: no cache initially; LibRaw owns the static list and calls are cheap enough for clarity.
- Answer: No cache. CameraList rebuilds the Go slice from LibRaw's static list on each call; documented as such.

## Implementation Outcome

- Package helpers: `CameraList`, `CameraCount`, `Capabilities`.
- Processor helpers: `UnpackFunctionName`, `DecoderInfo`, `IWidth`, `IHeight`, `CamMul`, `PreMul`, `RGBCam`, `ColorMaximum`, `AdjustToRawInsetCrop`.
- Index validation: `CamMul`/`PreMul` require 0..3; `RGBCam` requires row 0..2, col 0..3 (reuses validateIndex, returns errors).
- `libraw_adjust_to_raw_inset_crop` is LibRaw 0.22+; guarded in the cgo preamble and reported as `ErrUnsupported` on older libraries (CI runs 0.21.2). All other helpers exist in 0.21.2 and need no guard (verified against the upstream 0.21.2 header).
- `libraw_get_iparams`/`get_lensinfo`/`get_imgother` are not re-wrapped (metadata struct mapping is out of scope); their data is already exposed via `Processor.Metadata()`. Marked wrapped-via-metadata in the coverage map.

## Git And PR

- Branch: `feature/TASK-012-camera-capabilities-decoder-helpers`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - api
  - helpers
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: helper APIs overlap with metadata snapshots.
- Mitigation: document direct helper methods separately from snapshot metadata.


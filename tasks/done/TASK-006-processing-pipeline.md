# TASK-006: Wrap Processing Pipeline APIs

## Goal

Expose LibRaw unpacking, thumbnail extraction, raw-to-image conversion, dcraw processing, and writer functions.

## Context

The main processing C API includes `libraw_unpack`, `libraw_unpack_thumb`, `libraw_unpack_thumb_ex`, `libraw_subtract_black`, `libraw_raw2image`, `libraw_free_image`, `libraw_adjust_sizes_info_only`, `libraw_dcraw_process`, `libraw_dcraw_ppm_tiff_writer`, `libraw_dcraw_thumb_writer`, `libraw_dcraw_make_mem_image`, `libraw_dcraw_make_mem_thumb`, and `libraw_dcraw_clear_mem`.

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
  - `Makefile`
- Read-only paths:
  - original RAW fixtures.
- Forbidden paths:
  - broad metadata model work except fields needed by tests.
- Output artifacts:
  - generated test images under `tmp/outputs/` or cleaned `testdata/*.jpg` only when Makefile target allows it.
- Default branch: `main`.
- Target branch: `feature/TASK-006-processing-pipeline`.
- Package manager: Go modules.
- Runtime: Go with cgo enabled.
- Required environment: LibRaw and bundled RAW fixtures.

## Scope

- Add typed processing methods for unpack, thumbnail unpack, raw-to-image, and dcraw processing.
- Add memory image and memory thumbnail wrappers with safe freeing.
- Add file writer methods for PPM/TIFF and thumbnails.
- Update `_example` to develop the sample RAW into an output image.

## Out Of Scope

- Full output parameter coverage.
- Full image pixel access abstractions beyond what writer and memory image APIs need.

## Acceptance Criteria

- Given `testdata/RAW_CANON_6D.CR2`, when processed through the documented example, then an output image file is produced.
- Given a memory image is requested, when it is closed/freed, then repeated test runs do not leak obvious C allocations.
- Given processing is called before required previous steps, when LibRaw returns an error, then the Go API returns a typed error.

## Test Requirements

- Unit tests: required for memory wrapper ownership.
- Integration tests: required for processing at least one bundled fixture.
- E2E tests: example build/run required.
- Formatting: `gofmt` must be applied to changed Go files when Go files are changed.
- Lint: use `make lint` when available; otherwise use `go vet ./...`.
- Static analysis: `go vet ./...` is required for Go changes.
- Coverage scope: processing wrapper logic.
- Coverage metric: line coverage.
- Coverage target: at least `75%` for non-generated Go code.
- Required commands:
  - `gofmt` on changed Go files
  - `go vet ./...`
  - `go test ./...`
  - `go test -race ./...`
  - `go run ./_example`

## Language Requirements

- Language: Go for implementation and tooling; Markdown for workflowr docs and task files.
- Style: idiomatic Go, simple package boundaries, explicit errors, table-driven tests when helpful.
- Dependency policy: prefer the Go standard library; do not add external dependencies unless explicitly required by the task.
- Standard tools: `gofmt`, `go vet`, `go test`, and project Makefile targets.

## Implementation Notes

Avoid committing generated images. Add cleanup hooks or write to `tmp/outputs/` by default.

## Clarifications

- Question: Should example output be JPEG, TIFF, or PPM?
- Recommended default: use the closest LibRaw-supported writer path first; document format explicitly.
- Answer: PPM. It is LibRaw's native `dcraw_ppm_tiff_writer` output (no extra encoder/dependency). The example writes `tmp/outputs/RAW_CANON_6D.ppm`, a git-ignored path, avoiding committed generated images.

## Implementation Outcome

- Wrapped unpack (`Unpack`, `UnpackThumb`, `UnpackThumbEx`), `SubtractBlack`, `Raw2Image`, `FreeImage`, `AdjustSizesInfoOnly`, `DcrawProcess`, writers (`WritePPMTiff`, `WriteThumb`), and in-memory renderers (`MemImage`, `MemThumb`).
- `MemImage`/`MemThumb` copy `libraw_processed_image_t` into Go-owned bytes and free the C allocation via `libraw_dcraw_clear_mem` before returning, so callers have nothing to release.
- `_example` now develops `RAW_CANON_6D.CR2` end to end into `tmp/outputs/RAW_CANON_6D.ppm`; `make clean` removes `tmp/outputs`.
- Note: `libraw_dcraw_make_mem_image`/`_thumb` are not in the generated API inventory because their declarations span two header lines (the regex matches single-line `DllDef ... name(`); they are still wrapped. Tracked as an inventory limitation, not a coverage gap.

## Git And PR

- Branch: `feature/TASK-006-processing-pipeline`
- Commit style: Conventional Commits
- PR target: `main`
- PR provider: GitHub
- PR tool: GitHub CLI (`gh`)
- PR mode: draft
- Push branch automatically: yes, after required checks pass
- PR labels:
  - processing
  - cgo
  - agent-generated
- PR reviewers:
- PR assignees:
- Review requirements: self-review before PR

## Risks

- Risk: processing tests are slow or platform-sensitive.
- Mitigation: keep one fast integration fixture required and mark heavier fixture matrix separately.


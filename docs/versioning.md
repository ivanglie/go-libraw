# Versioning Policy

## Go module version

Git tags follow standard semver: `v0.1.0`, `v0.2.0`, `v1.0.0`.

- **Patch** — bug fixes, docs, CI, build scripts; no public API change.
- **Minor** — new public APIs or a newer LibRaw baseline; backwards-compatible.
- **Major** — breaking API change; requires `/v2` module path suffix.

Before `v1.0.0` the public API may still change in minor releases.

## LibRaw baseline

Each release notes the LibRaw header baseline in the release description:

```
Git tag:      v0.1.0
Release name: go-libraw v0.1.0 for LibRaw 0.22.1
```

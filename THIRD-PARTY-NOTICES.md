# Third-Party Notices

`go-libraw` is licensed under the MIT License (see [LICENSE](LICENSE)). That
license covers **only the Go binding code in this repository**. The components
listed below are owned by their respective authors and remain subject to their
own licenses. Distributing software that links LibRaw — or that bundles the
files noted here — carries the obligations described in those licenses.

## 1) LibRaw

- **Upstream:** https://github.com/LibRaw/LibRaw
- **License:** CDDL-1.0 OR LGPL-2.1-or-later (upstream dual licensing — the
  distributor may choose either arm)
- **License texts:**
  - [licenses/LibRaw-LICENSE.CDDL](licenses/LibRaw-LICENSE.CDDL)
  - [licenses/LibRaw-LICENSE.LGPL](licenses/LibRaw-LICENSE.LGPL)

`go-libraw` does not vendor LibRaw; it links against a LibRaw installed on the
build/host system via cgo. LibRaw is not redistributed by this repository, but
software built from it links LibRaw and must satisfy LibRaw's license.

### Distribution notes

- **Dynamic linking** (the default) satisfies LGPL-2.1's relinking requirement
  on its own: ship the LibRaw notice and the license text, and keep LibRaw
  replaceable as a shared library.
- **Static linking** (the `libraw_static` build tag) is simplest under the
  **CDDL-1.0** arm, which does not impose LGPL's relinking obligation. Retain
  the CDDL notice and make the corresponding LibRaw source available.

### Bundled LibRaw headers

For offline API-coverage tooling and tests, this repository checks in a copy of
LibRaw's public headers under `testdata/headers/libraw/` (`libraw.h`,
`libraw_const.h`, `libraw_types.h`, `libraw_version.h`). These files are part of
LibRaw, carry their original `Copyright 2008-2025 LibRaw LLC` notices, and
remain under LibRaw's CDDL-1.0 / LGPL-2.1-or-later dual license — **not** the
MIT license of this repository.

## 2) RAW test fixtures

The sample RAW files under `testdata/` are third-party camera files used only
for testing. See [testdata/README.md](testdata/README.md) for their provenance
and usage terms.

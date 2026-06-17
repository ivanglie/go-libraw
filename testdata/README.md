# Test Data

The RAW sample files here are used solely for automated testing. They come from
[raw.pixls.us](https://raw.pixls.us) and are released under
[CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/) (public domain
dedication) — free to use, redistribute, and modify without restriction or
attribution.

| File | Camera | Format | Vendor coverage |
| --- | --- | --- | --- |
| `RAW_CANON_R6.CR3` | Canon EOS R6 | CR3 | Canon maker notes; embedded thumbnail |
| `RAW_NIKON_ZFC.NEF` | Nikon Z fc | NEF (lossless compressed) | Nikon maker notes |
| `RAW_SONY_ILCE-7M4.ARW` | Sony A7 IV | ARW | Sony maker notes |
| `RAW_RICOH_GR3X.DNG` | Ricoh GR IIIx | DNG | Ricoh / DNG zero-value policy |

Bodies are chosen to stay within the LibRaw version linked in CI: very recent
cameras may not yet be supported by the CI LibRaw even when they work against a
newer local build. The `headers/` directory holds LibRaw's public headers and is
covered by LibRaw's own license, not CC0 (see [THIRD-PARTY-NOTICES.md](../THIRD-PARTY-NOTICES.md)).

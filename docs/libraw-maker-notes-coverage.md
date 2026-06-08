# LibRaw Maker Notes Coverage

This document records field-level policy for maker-note metadata exposed through
`Processor.Metadata().MakerNotes`.

## Policy

- Vendor maker-note groups are value structs.
- If a RAW file does not contain a vendor group, LibRaw leaves the Go snapshot at
  the zero value.
- Scalar fields, fixed strings, fixed arrays, crop areas, lens metadata, and
  high-speed crop metadata are copied into Go-owned values.
- C-owned pointer payloads are summarized by length and presence flags.
- No semantic interpretation is added beyond LibRaw field names and values.

## Covered Structs

| Upstream struct | Go type | Status | Notes |
| --- | --- | --- | --- |
| `libraw_makernotes_t` | `MakerNotes` | wrapped | All vendor groups copied as value structs. |
| `libraw_canon_makernotes_t` | `CanonMakerNotes` | wrapped | Includes crop `Area` fields and fixed arrays. |
| `libraw_nikon_makernotes_t` | `NikonMakerNotes` | wrapped | `BurstTable_0x0056` pointer summarized by length and `HasBurstTable0056`. |
| `libraw_hasselblad_makernotes_t` | `HasselbladMakerNotes` | wrapped | Full scalar/string/matrix coverage. |
| `libraw_fuji_info_t` | `FujiMakerNotes` | wrapped | Full scalar/string/fixed-array coverage. |
| `libraw_olympus_makernotes_t` | `OlympusMakerNotes` | wrapped | Decoder tags are grouped into `DecoderTags` in upstream order. |
| `libraw_sony_info_t` | `SonyMakerNotes` | wrapped | Full scalar/string/fixed-array coverage for the current header baseline. |
| `libraw_kodak_makernotes_t` | `KodakMakerNotes` | wrapped | Full scalar/matrix coverage. |
| `libraw_panasonic_makernotes_t` | `PanasonicMakerNotes` | wrapped | Full scalar/fixed-array coverage. |
| `libraw_pentax_makernotes_t` | `PentaxMakerNotes` | wrapped | Full scalar/fixed-array coverage. |
| `libraw_p1_makernotes_t` | `PhaseOneMakerNotes` | wrapped | Full string coverage. |
| `libraw_ricoh_makernotes_t` | `RicohMakerNotes` | wrapped | Full scalar/fixed-array coverage. |
| `libraw_samsung_makernotes_t` | `SamsungMakerNotes` | wrapped | Full scalar/string/fixed-array coverage. |
| `libraw_metadata_common_t` | `MetadataCommon` | wrapped | AF data list is bounded by `LIBRAW_AFDATA_MAXCOUNT`. |
| `libraw_afinfo_item_t` | `AFInfoItem` | summarized | `AFInfoData` pointer summarized by `Length` and `HasData`. |
| `libraw_area_t` | `Area` | wrapped | Used by Canon maker notes. |
| `libraw_sensor_highspeed_crop_t` | `SensorHighSpeedCrop` | wrapped | Used by Nikon maker notes. |

## Fixture Coverage

| Fixture | Vendor path tested |
| --- | --- |
| `testdata/RAW_CANON_6D.CR2` | Canon maker notes |
| `testdata/RAW_NIKON_D750.NEF` | Nikon maker notes |
| `testdata/RAW_RICOH_GR2.DNG` | Ricoh/DNG zero-value policy |
| `testdata/RAW_SONY_ILCA-77M2.ARW` | Sony maker notes |

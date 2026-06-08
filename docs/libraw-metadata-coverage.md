# LibRaw Core Metadata Coverage

This document records field-level coverage for the core metadata snapshot added
by `Processor.Metadata`.

## Policy

- Scalar values and fixed arrays are copied into Go-owned values.
- C-owned pointers are not exposed directly.
- Raw pixel buffers are represented as presence flags in `RawDataSummary`; byte
  access is intentionally deferred to the raw/image memory task.
- Vendor-specific maker note structs remain out of scope for this task.

## Covered Structs

| Upstream struct | Go type | Status | Notes |
| --- | --- | --- | --- |
| `libraw_iparams_t` | `ImageParams` | wrapped | String, scalar, fixed-array fields copied; `xmpdata` summarized by `XMLLen` and `HasXML`. |
| `libraw_raw_inset_crop_t` | `RawInsetCrop` | wrapped | Full scalar coverage. |
| `libraw_image_sizes_t` | `ImageSizes` | wrapped | Full scalar and fixed-array coverage. |
| `libraw_internal_output_params_t` | `InternalOutputParams` | wrapped | Full scalar coverage inside `RawDataSummary`. |
| `libraw_dng_color_t` | `DNGColor` | wrapped | Full scalar and matrix coverage. |
| `libraw_dng_rawopcode_t` | `DNGRawOpcode` | summarized | Payload pointer is summarized by `Len` and `HasData`. |
| `libraw_dng_levels_t` | `DNGLevels` | wrapped | Full scalar and fixed-array coverage; raw opcodes summarized. |
| `libraw_P1_color_t` | `P1Color` | wrapped | Full fixed-array coverage. |
| `struct ph1_t` | `PhaseOneData` | wrapped | Full scalar coverage through `ColorData`. |
| `libraw_colordata_t` | `ColorData` | wrapped | Full scalar and fixed-array coverage; ICC profile pointer summarized by `ProfileLength` and `HasProfile`. |
| `libraw_thumbnail_t` | `Thumbnail` | summarized | Metadata copied; thumbnail data pointer summarized by `HasData`. |
| `libraw_thumbnail_item_t` | `ThumbnailItem` | wrapped | Full scalar coverage. |
| `libraw_thumbnail_list_t` | `ThumbnailList` | wrapped | Count and bounded item snapshots copied. |
| `libraw_gps_info_t` | `GPSInfo` | wrapped | Full scalar and fixed-array coverage. |
| `libraw_imgother_t` | `ImageOther` | wrapped | Full scalar and fixed-array coverage. |
| `libraw_nikonlens_t` | `NikonLens` | wrapped | Full scalar coverage. |
| `libraw_dnglens_t` | `DNGLens` | wrapped | Full scalar coverage. |
| `libraw_makernotes_lens_t` | `MakerNotesLens` | wrapped | Lens metadata only; not vendor maker-note structs. |
| `libraw_lensinfo_t` | `LensInfo` | wrapped | Full scalar/string coverage including nested lens structs. |
| `libraw_shootinginfo_t` | `ShootingInfo` | wrapped | Full scalar/string coverage. |
| `libraw_rawdata_t` | `RawDataSummary` | summarized | Raw pixel pointers summarized; embedded metadata snapshots copied. |
| `libraw_data_t` | `Metadata` | summarized | Core fields copied; params/rawparams are covered by TASK-007; vendor maker notes are out of scope. |
| `libraw_processed_image_t` | `ProcessedImage` | wrapped | Covered by memory image APIs from TASK-006. |

## Out Of Scope Or Deferred

| Upstream struct | Status | Reason |
| --- | --- | --- |
| `libraw_callbacks_t` | deferred | Callback APIs are covered by TASK-011. |
| `libraw_custom_camera_t` | deferred | Decoder helper/custom camera work is covered by TASK-012. |
| `libraw_area_t` | deferred | Currently only needed inside vendor-specific maker notes, which are out of scope here. |
| `libraw_metadata_common_t` | deferred | Lives inside vendor maker notes and is covered by TASK-009. |
| vendor maker-note structs | deferred | TASK-009 owns vendor maker-note metadata. |

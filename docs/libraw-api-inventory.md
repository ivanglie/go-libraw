# LibRaw API Inventory

- Header directory: `testdata/headers/libraw`
- Header version: `0.22.1-Release`
- Total symbols: `204`

Statuses:

- `wrapped`: covered by the current Go API.
- `internal`: used behind the public Go API boundary.
- `deferred`: in scope for later workflowr tasks.
- `unsupported`: intentionally not planned.
- `unmapped`: present upstream but missing from coverage map.

## Functions

| Symbol | Header | Status | Note |
| --- | --- | --- | --- |
| `libraw_COLOR` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_adjust_sizes_info_only` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_adjust_to_raw_inset_crop` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_cameraCount` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_cameraList` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_capabilities` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_close` | `libraw.h` | `internal` | used by Processor.Close |
| `libraw_dcraw_clear_mem` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dcraw_ppm_tiff_writer` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dcraw_process` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dcraw_thumb_writer` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_free_image` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_cam_mul` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_color_maximum` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_decoder_info` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_iheight` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_imgother` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_iparams` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_iwidth` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_lensinfo` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_pre_mul` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_raw_height` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_raw_width` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_get_rgb_cam` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_init` | `libraw.h` | `internal` | used by NewProcessor |
| `libraw_open_bayer` | `libraw.h` | `wrapped` | wrapped by Processor.OpenBayer |
| `libraw_open_buffer` | `libraw.h` | `wrapped` | wrapped by Processor.OpenBuffer |
| `libraw_open_file` | `libraw.h` | `wrapped` | wrapped by Processor.OpenFile |
| `libraw_open_file_ex` | `libraw.h` | `unsupported` | removed from default 0.22 build via LIBRAW_NO_IOSTREAMS_DATASTREAM |
| `libraw_open_wfile` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_open_wfile_ex` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_raw2image` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_recycle` | `libraw.h` | `wrapped` | wrapped by Processor.Recycle |
| `libraw_recycle_datastream` | `libraw.h` | `wrapped` | wrapped by Processor.RecycleDatastream |
| `libraw_set_adjust_maximum_thr` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_bright` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_dataerror_handler` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_demosaic` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_exifparser_handler` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_fbdd_noiserd` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_gamma` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_highlight` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_makernotes_handler` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_no_auto_bright` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_output_bps` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_output_color` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_output_tif` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_progress_handler` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_set_user_mul` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_strerror` | `libraw.h` | `wrapped` | exposed as StrError and ErrorCode.String |
| `libraw_strprogress` | `libraw.h` | `wrapped` | exposed as StrProgress and Progress.String |
| `libraw_subtract_black` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_unpack` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_unpack_function_name` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_unpack_thumb` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_unpack_thumb_ex` | `libraw.h` | `deferred` | tracked for a future workflowr task |
| `libraw_version` | `libraw.h` | `wrapped` | exposed as Version |
| `libraw_versionNumber` | `libraw.h` | `wrapped` | exposed as VersionNumber |

## Enums

| Symbol | Header | Status | Note |
| --- | --- | --- | --- |
| `LIBRAW_SONY_FOCUSMODEmodes` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRawImageAspects` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_As_Shot_WB_Applied_codes` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_Canon_RecordModes` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_ExifTagTypes` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_HasselbladFormatCodes` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_KodakSensors` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_MultiExposure_related` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_Sony_0x2010_Type` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_Sony_0x9050_Type` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_camera_formats` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_camera_mounts` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_cameramaker_index` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_colorspace` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_constructor_flags` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_decoder_flags` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_dng_processing` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_dngfields_marks` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_errors` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_exceptions` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_image_formats` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_internal_thumbnail_formats` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_lens_focal_types` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_minolta_bayerpatterns` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_minolta_storagemethods` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_open_flags` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_openbayer_patterns` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_output_flags` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_processing_options` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_progress` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_rawspecial_t` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_rawspeed_bits_t` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_runtime_capabilities` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_sony_cameratypes` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_thumbnail_formats` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_warnings` | `libraw_const.h` | `wrapped` | generated constants expose enum members |
| `LibRaw_whitebalance_code` | `libraw_const.h` | `wrapped` | generated constants expose enum members |

## Macros

| Symbol | Header | Status | Note |
| --- | --- | --- | --- |
| `LIBRAW_AFDATA_MAXCOUNT` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_AHD_TILE` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_CBLACK_SIZE` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_CHECK_VERSION` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_COMPILE_CHECK_VERSION` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_COMPILE_CHECK_VERSION_NOTLESS` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_CR3_MEMPOOL` | `libraw_const.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_CRXTRACKS_MAXCOUNT` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_DEFAULT_ADJUST_MAXIMUM_THRESHOLD` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_DEFAULT_AUTO_BRIGHTNESS_THRESHOLD` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_binary` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_complex` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_double` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_float` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_ifd` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_ifd64` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int16s` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int16u` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int32s` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int32u` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int64s` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int64u` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int8s` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_int8u` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_rational64s` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_rational64u` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_string` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_undef` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_EXIFTOOLTAGTYPE_unicode` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_FATAL_ERROR` | `libraw_const.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_IFD_MAXCOUNT` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_IOSPACE_CHECK` | `libraw_const.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_LENS_NOT_SET` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAKE_VERSION` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_MAX_ALLOC_MB_DEFAULT` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAX_CR3_RAW_FILE_SIZE` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAX_DNG_RAW_FILE_SIZE` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAX_METADATA_BLOCKS` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAX_NONDNG_RAW_FILE_SIZE` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAX_PROFILE_SIZE_MB` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MAX_THUMBNAIL_MB` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_MEMPOOL_CHECK` | `libraw_const.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_NO_IOSTREAMS_DATASTREAM` | `libraw.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_OWN_SWAB` | `libraw_const.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_PROGRESS_THUMB_MASK` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_RUNTIME_CHECK_VERSION_EXACT` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_RUNTIME_CHECK_VERSION_NOTLESS` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_SHLIB_AGE` | `libraw_version.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_SHLIB_CURRENT` | `libraw_version.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_SHLIB_REVISION` | `libraw_version.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_THUMBNAIL_MAXCOUNT` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_USE_OPENMP` | `libraw_types.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_USE_STREAMS_DATASTREAM_MAXSIZE` | `libraw.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_VERSION_MAKE` | `libraw_version.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_WIN32_CALLS` | `libraw.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_WIN32_DLLDEFS` | `libraw.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_WIN32_UNICODEPATHS` | `libraw.h` | `unsupported` | preprocessor switch or function-like macro not exposed as a Go constant |
| `LIBRAW_X3F_ALLOC_LIMIT_MB` | `libraw_const.h` | `wrapped` | generated from value-like macro |
| `LIBRAW_XTRANS` | `libraw_const.h` | `wrapped` | generated from value-like macro |

## Structs

| Symbol | Header | Status | Note |
| --- | --- | --- | --- |
| `libraw_P1_color_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_afinfo_item_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_area_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_callbacks_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_canon_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_colordata_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_custom_camera_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_data_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_decoder_info_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dng_color_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dng_levels_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dng_rawopcode_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_dnglens_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_fuji_info_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_gps_info_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_hasselblad_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_image_sizes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_imgother_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_internal_output_params_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_iparams_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_kodak_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_lensinfo_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_makernotes_lens_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_metadata_common_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_nikon_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_nikonlens_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_olympus_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_output_params_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_p1_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_panasonic_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_pentax_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_processed_image_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_raw_inset_crop_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_raw_unpack_params_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_rawdata_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_ricoh_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_samsung_makernotes_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_sensor_highspeed_crop_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_shootinginfo_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_sony_info_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_thumbnail_item_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_thumbnail_list_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |
| `libraw_thumbnail_t` | `libraw_types.h` | `deferred` | tracked for a future workflowr task |

## Versions

| Symbol | Header | Status | Note |
| --- | --- | --- | --- |
| `LIBRAW_MAJOR_VERSION` | `libraw_version.h` | `wrapped` | generated numeric version constant |
| `LIBRAW_MINOR_VERSION` | `libraw_version.h` | `wrapped` | generated numeric version constant |
| `LIBRAW_PATCH_VERSION` | `libraw_version.h` | `wrapped` | generated numeric version constant |
| `LIBRAW_VERSION` | `libraw_version.h` | `wrapped` | generated numeric version constant |
| `LIBRAW_VERSION_STR` | `libraw_version.h` | `wrapped` | exposed at runtime via Version |
| `LIBRAW_VERSION_TAIL` | `libraw_version.h` | `unsupported` | non-numeric preprocessor token not exposed as a Go constant |


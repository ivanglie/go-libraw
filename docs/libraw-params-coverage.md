# LibRaw Parameter Field Coverage

This document records field-level coverage for `libraw_output_params_t` and
`libraw_raw_unpack_params_t` from `testdata/headers/libraw/libraw_types.h`.

## libraw_output_params_t

| Upstream field | Go field | Status | Notes |
| --- | --- | --- | --- |
| `greybox[4]` | `OutputParams.Greybox` | wrapped | Full get/set round-trip. |
| `cropbox[4]` | `OutputParams.Cropbox` | wrapped | Full get/set round-trip. |
| `aber[4]` | `OutputParams.Aber` | wrapped | Full get/set round-trip. |
| `gamm[6]` | `OutputParams.Gamm` | wrapped | Full get/set plus `Processor.SetGamma` with index validation. |
| `user_mul[4]` | `OutputParams.UserMul` | wrapped | Full get/set plus `Processor.SetUserMul` with index validation. |
| `bright` | `OutputParams.Bright` | wrapped | Full get/set plus `Processor.SetBright`. |
| `threshold` | `OutputParams.Threshold` | wrapped | Full get/set round-trip. |
| `half_size` | `OutputParams.HalfSize` | wrapped | Full get/set round-trip. |
| `four_color_rgb` | `OutputParams.FourColorRGB` | wrapped | Full get/set round-trip. |
| `highlight` | `OutputParams.Highlight` | wrapped | Full get/set plus `Processor.SetHighlight`. |
| `use_auto_wb` | `OutputParams.UseAutoWB` | wrapped | Full get/set round-trip. |
| `use_camera_wb` | `OutputParams.UseCameraWB` | wrapped | Full get/set round-trip. |
| `use_camera_matrix` | `OutputParams.UseCameraMatrix` | wrapped | Full get/set round-trip. |
| `output_color` | `OutputParams.OutputColor` | wrapped | Full get/set plus `Processor.SetOutputColor`. |
| `output_profile` | `OutputParams.OutputProfile` | wrapped | Copied into C memory retained by the handle. |
| `camera_profile` | `OutputParams.CameraProfile` | wrapped | Copied into C memory retained by the handle. |
| `bad_pixels` | `OutputParams.BadPixels` | wrapped | Copied into C memory retained by the handle. |
| `dark_frame` | `OutputParams.DarkFrame` | wrapped | Copied into C memory retained by the handle. |
| `output_bps` | `OutputParams.OutputBPS` | wrapped | Full get/set plus `Processor.SetOutputBPS` validation for 8 or 16. |
| `output_tiff` | `OutputParams.OutputTIFF` | wrapped | Full get/set plus `Processor.SetOutputTIFF`. |
| `output_flags` | `OutputParams.OutputFlags` | wrapped | Full get/set round-trip. |
| `user_flip` | `OutputParams.UserFlip` | wrapped | Full get/set round-trip. |
| `user_qual` | `OutputParams.UserQual` | wrapped | Full get/set plus `Processor.SetDemosaic`. |
| `user_black` | `OutputParams.UserBlack` | wrapped | Full get/set round-trip. |
| `user_cblack[4]` | `OutputParams.UserCblack` | wrapped | Full get/set round-trip. |
| `user_sat` | `OutputParams.UserSat` | wrapped | Full get/set round-trip. |
| `med_passes` | `OutputParams.MedPasses` | wrapped | Full get/set round-trip. |
| `auto_bright_thr` | `OutputParams.AutoBrightThr` | wrapped | Full get/set round-trip. |
| `adjust_maximum_thr` | `OutputParams.AdjustMaximumThr` | wrapped | Full get/set plus `Processor.SetAdjustMaximumThreshold`. |
| `no_auto_bright` | `OutputParams.NoAutoBright` | wrapped | Full get/set plus `Processor.SetNoAutoBright`. |
| `use_fuji_rotate` | `OutputParams.UseFujiRotate` | wrapped | Full get/set round-trip. |
| `use_p1_correction` | `OutputParams.UseP1Correction` | wrapped | Full get/set round-trip. |
| `green_matching` | `OutputParams.GreenMatching` | wrapped | Full get/set round-trip. |
| `dcb_iterations` | `OutputParams.DCBIterations` | wrapped | Full get/set round-trip. |
| `dcb_enhance_fl` | `OutputParams.DCBEnhanceFL` | wrapped | Full get/set round-trip. |
| `fbdd_noiserd` | `OutputParams.FBDDNoiseRD` | wrapped | Full get/set plus `Processor.SetFBDDNoiseReduction`. |
| `exp_correc` | `OutputParams.ExpCorrec` | wrapped | Full get/set round-trip. |
| `exp_shift` | `OutputParams.ExpShift` | wrapped | Full get/set round-trip. |
| `exp_preser` | `OutputParams.ExpPreser` | wrapped | Full get/set round-trip. |
| `no_auto_scale` | `OutputParams.NoAutoScale` | wrapped | Full get/set round-trip. |
| `no_interpolation` | `OutputParams.NoInterpolation` | wrapped | Full get/set round-trip. |

## libraw_raw_unpack_params_t

| Upstream field | Go field | Status | Notes |
| --- | --- | --- | --- |
| `use_rawspeed` | `RawUnpackParams.UseRawSpeed` | wrapped | Full get/set round-trip. |
| `use_dngsdk` | `RawUnpackParams.UseDNGSDK` | wrapped | Full get/set round-trip. |
| `options` | `RawUnpackParams.Options` | wrapped | Full get/set round-trip. |
| `shot_select` | `RawUnpackParams.ShotSelect` | wrapped | Full get/set round-trip. |
| `specials` | `RawUnpackParams.Specials` | wrapped | Full get/set round-trip. |
| `max_raw_memory_mb` | `RawUnpackParams.MaxRawMemoryMB` | wrapped | Full get/set round-trip. |
| `sony_arw2_posterization_thr` | `RawUnpackParams.SonyARW2PosterizationTh` | wrapped | Full get/set round-trip. |
| `coolscan_nef_gamma` | `RawUnpackParams.CoolscanNEFGamma` | wrapped | Full get/set round-trip. |
| `p4shot_order[5]` | `RawUnpackParams.P4ShotOrder` | wrapped | Up to 4 bytes plus NUL; public setter rejects longer values. |
| `custom_camera_strings` | none | unsupported | Requires a dedicated owned string-list API and is intentionally not mirrored. |

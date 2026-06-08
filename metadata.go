//go:build cgo

package libraw

import "github.com/ivanglie/go-libraw/internal/librawc"

// Metadata is a Go-owned snapshot of core LibRaw metadata.
type Metadata = librawc.Metadata

// ProcessState reports LibRaw processing flags and warnings.
type ProcessState = librawc.ProcessState

// ImageParams mirrors libraw_iparams_t.
type ImageParams = librawc.ImageParams

// RawInsetCrop mirrors libraw_raw_inset_crop_t.
type RawInsetCrop = librawc.RawInsetCrop

// ImageSizes mirrors libraw_image_sizes_t.
type ImageSizes = librawc.ImageSizes

// InternalOutputParams mirrors libraw_internal_output_params_t.
type InternalOutputParams = librawc.InternalOutputParams

// DNGColor mirrors libraw_dng_color_t.
type DNGColor = librawc.DNGColor

// DNGRawOpcode summarizes libraw_dng_rawopcode_t without exposing C-owned data.
type DNGRawOpcode = librawc.DNGRawOpcode

// DNGLevels mirrors libraw_dng_levels_t, summarizing raw opcode payloads.
type DNGLevels = librawc.DNGLevels

// PhaseOneData mirrors struct ph1_t.
type PhaseOneData = librawc.PhaseOneData

// P1Color mirrors libraw_P1_color_t.
type P1Color = librawc.P1Color

// ColorData mirrors libraw_colordata_t, summarizing profile payloads.
type ColorData = librawc.ColorData

// Thumbnail mirrors libraw_thumbnail_t without exposing the thumbnail pointer.
type Thumbnail = librawc.Thumbnail

// ThumbnailItem mirrors libraw_thumbnail_item_t.
type ThumbnailItem = librawc.ThumbnailItem

// ThumbnailList mirrors libraw_thumbnail_list_t.
type ThumbnailList = librawc.ThumbnailList

// GPSInfo mirrors libraw_gps_info_t.
type GPSInfo = librawc.GPSInfo

// ImageOther mirrors libraw_imgother_t.
type ImageOther = librawc.ImageOther

// NikonLens mirrors libraw_nikonlens_t.
type NikonLens = librawc.NikonLens

// DNGLens mirrors libraw_dnglens_t.
type DNGLens = librawc.DNGLens

// MakerNotesLens mirrors libraw_makernotes_lens_t.
type MakerNotesLens = librawc.MakerNotesLens

// LensInfo mirrors libraw_lensinfo_t.
type LensInfo = librawc.LensInfo

// ShootingInfo mirrors libraw_shootinginfo_t.
type ShootingInfo = librawc.ShootingInfo

// RawDataSummary mirrors libraw_rawdata_t without exposing raw image pointers.
type RawDataSummary = librawc.RawDataSummary

// Metadata returns a Go-owned snapshot of the current LibRaw metadata.
func (p *Processor) Metadata() (Metadata, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || p.handle == nil {
		return Metadata{}, ErrClosed
	}
	return p.handle.GetMetadata(), nil
}

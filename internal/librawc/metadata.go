//go:build cgo

package librawc

/*
#include <libraw/libraw.h>

static unsigned go_libraw_dng_rawopcode_len(libraw_dng_levels_t *p, int i) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->rawopcodes[i].len;
#else
	(void)p;
	(void)i;
	return 0;
#endif
}

static int go_libraw_dng_rawopcode_has_data(libraw_dng_levels_t *p, int i) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->rawopcodes[i].data != 0;
#else
	(void)p;
	(void)i;
	return 0;
#endif
}
*/
import "C"

// Metadata is a Go-owned snapshot of core LibRaw metadata.
type Metadata struct {
	Sizes     ImageSizes
	ID        ImageParams
	Lens      LensInfo
	Shooting  ShootingInfo
	Color     ColorData
	Other     ImageOther
	Thumbnail Thumbnail
	Thumbs    ThumbnailList
	RawData   RawDataSummary
	Process   ProcessState
}

// ProcessState reports LibRaw processing flags and warnings.
type ProcessState struct {
	ProgressFlags   uint32
	ProcessWarnings uint32
}

// ImageParams mirrors libraw_iparams_t.
type ImageParams struct {
	Guard           string
	Make            string
	Model           string
	Software        string
	NormalizedMake  string
	NormalizedModel string
	MakerIndex      uint32
	RawCount        uint32
	DNGVersion      uint32
	IsFoveon        uint32
	Colors          int
	Filters         uint32
	XTrans          [6][6]int8
	XTransAbs       [6][6]int8
	CDesc           string
	XMLLen          uint32
	HasXML          bool
}

// RawInsetCrop mirrors libraw_raw_inset_crop_t.
type RawInsetCrop struct {
	Left   uint16
	Top    uint16
	Width  uint16
	Height uint16
}

// ImageSizes mirrors libraw_image_sizes_t.
type ImageSizes struct {
	RawHeight     uint16
	RawWidth      uint16
	Height        uint16
	Width         uint16
	TopMargin     uint16
	LeftMargin    uint16
	IHeight       uint16
	IWidth        uint16
	RawPitch      uint32
	PixelAspect   float64
	Flip          int
	Mask          [8][4]int
	RawAspect     uint16
	RawInsetCrops [2]RawInsetCrop
}

// InternalOutputParams mirrors libraw_internal_output_params_t.
type InternalOutputParams struct {
	MixGreen  uint32
	RawColor  uint32
	ZeroIsBad uint32
	Shrink    uint16
	FujiWidth uint16
}

// DNGColor mirrors libraw_dng_color_t.
type DNGColor struct {
	ParsedFields  uint32
	Illuminant    uint16
	Calibration   [4][4]float32
	ColorMatrix   [4][3]float32
	ForwardMatrix [3][4]float32
}

// DNGRawOpcode summarizes libraw_dng_rawopcode_t without exposing C-owned data.
type DNGRawOpcode struct {
	Len     uint32
	HasData bool
}

// DNGLevels mirrors libraw_dng_levels_t, summarizing raw opcode payloads.
type DNGLevels struct {
	ParsedFields        uint32
	DNGCBlack           [6]uint32
	DNGBlack            uint32
	DNGFCBlack          [6]float32
	DNGFBlack           float32
	DNGWhiteLevel       [4]uint32
	DefaultCrop         [4]uint16
	UserCrop            [4]float32
	PreviewColorspace   uint32
	AnalogBalance       [4]float32
	AsShotNeutral       [4]float32
	BaselineExposure    float32
	LinearResponseLimit float32
	RawOpcodes          [3]DNGRawOpcode
}

// PhaseOneData mirrors struct ph1_t.
type PhaseOneData struct {
	Format   int
	KeyOff   int
	Tag21A   int
	TBlack   int
	SplitCol int
	BlackCol int
	SplitRow int
	BlackRow int
	Tag210   float32
}

// P1Color mirrors libraw_P1_color_t.
type P1Color struct {
	ROMMCam [9]float32
}

// ColorData mirrors libraw_colordata_t, summarizing profile payloads.
type ColorData struct {
	Curve                [0x10000]uint16
	CBlack               [6]uint32
	Black                uint32
	DataMaximum          uint32
	Maximum              uint32
	LinearMax            [4]uint32
	FMaximum             float32
	FNorm                float32
	White                [8][8]uint16
	CamMul               [4]float32
	PreMul               [4]float32
	CMatrix              [3][4]float32
	CCM                  [3][4]float32
	RGBCam               [3][4]float32
	CamXYZ               [4][3]float32
	PhaseOneData         PhaseOneData
	FlashUsed            float32
	CanonEV              float32
	Model2               string
	UniqueCameraModel    string
	LocalizedCameraModel string
	ImageUniqueID        string
	RawDataUniqueID      string
	OriginalRawFileName  string
	ProfileLength        uint32
	HasProfile           bool
	BlackStat            [8]uint32
	DNGColor             [2]DNGColor
	DNGLevels            DNGLevels
	WBCoeffs             [256][4]int
	WBCTCoeffs           [64][5]float32
	AsShotWBApplied      int
	P1Color              [2]P1Color
	RawBPS               uint32
	ExifColorSpace       int
}

// Thumbnail mirrors libraw_thumbnail_t without exposing the thumbnail pointer.
type Thumbnail struct {
	Format  int
	Width   uint16
	Height  uint16
	Length  uint32
	Colors  int
	HasData bool
}

// ThumbnailItem mirrors libraw_thumbnail_item_t.
type ThumbnailItem struct {
	Format int
	Width  uint16
	Height uint16
	Flip   uint16
	Length uint32
	Misc   uint32
	Offset int64
}

// ThumbnailList mirrors libraw_thumbnail_list_t.
type ThumbnailList struct {
	Count int
	Items []ThumbnailItem
}

// GPSInfo mirrors libraw_gps_info_t.
type GPSInfo struct {
	Latitude  [3]float32
	Longitude [3]float32
	Timestamp [3]float32
	Altitude  float32
	AltRef    int8
	LatRef    int8
	LongRef   int8
	Status    int8
	Parsed    int8
}

// ImageOther mirrors libraw_imgother_t.
type ImageOther struct {
	ISOSpeed      float32
	Shutter       float32
	Aperture      float32
	FocalLen      float32
	Timestamp     int64
	ShotOrder     uint32
	GPSData       [32]uint32
	ParsedGPS     GPSInfo
	Description   string
	Artist        string
	AnalogBalance [4]float32
}

// NikonLens mirrors libraw_nikonlens_t.
type NikonLens struct {
	EffectiveMaxAp float32
	LensIDNumber   uint8
	LensFStops     uint8
	MCUVersion     uint8
	LensType       uint8
}

// DNGLens mirrors libraw_dnglens_t.
type DNGLens struct {
	MinFocal       float32
	MaxFocal       float32
	MaxAp4MinFocal float32
	MaxAp4MaxFocal float32
}

// MakerNotesLens mirrors libraw_makernotes_lens_t.
type MakerNotesLens struct {
	LensID                  uint64
	Lens                    string
	LensFormat              uint16
	LensMount               uint16
	CamID                   uint64
	CameraFormat            uint16
	CameraMount             uint16
	Body                    string
	FocalType               int16
	LensFeaturesPre         string
	LensFeaturesSuf         string
	MinFocal                float32
	MaxFocal                float32
	MaxAp4MinFocal          float32
	MaxAp4MaxFocal          float32
	MinAp4MinFocal          float32
	MinAp4MaxFocal          float32
	MaxAp                   float32
	MinAp                   float32
	CurFocal                float32
	CurAp                   float32
	MaxAp4CurFocal          float32
	MinAp4CurFocal          float32
	MinFocusDistance        float32
	FocusRangeIndex         float32
	LensFStops              float32
	TeleconverterID         uint64
	Teleconverter           string
	AdapterID               uint64
	Adapter                 string
	AttachmentID            uint64
	Attachment              string
	FocalUnits              uint16
	FocalLengthIn35mmFormat float32
}

// LensInfo mirrors libraw_lensinfo_t.
type LensInfo struct {
	MinFocal                float32
	MaxFocal                float32
	MaxAp4MinFocal          float32
	MaxAp4MaxFocal          float32
	EXIFMaxAp               float32
	LensMake                string
	Lens                    string
	LensSerial              string
	InternalLensSerial      string
	FocalLengthIn35mmFormat uint16
	Nikon                   NikonLens
	DNG                     DNGLens
	MakerNotes              MakerNotesLens
}

// ShootingInfo mirrors libraw_shootinginfo_t.
type ShootingInfo struct {
	DriveMode          int16
	FocusMode          int16
	MeteringMode       int16
	AFPoint            int16
	ExposureMode       int16
	ExposureProgram    int16
	ImageStabilization int16
	BodySerial         string
	InternalBodySerial string
}

// RawDataSummary mirrors libraw_rawdata_t without exposing raw image pointers.
type RawDataSummary struct {
	HasRawAlloc    bool
	HasRawImage    bool
	HasColor4Image bool
	HasColor3Image bool
	HasFloatImage  bool
	HasFloat3Image bool
	HasFloat4Image bool
	HasPH1CBlack   bool
	HasPH1RBlack   bool
	ImageParams    ImageParams
	Sizes          ImageSizes
	InternalOutput InternalOutputParams
	Color          ColorData
}

// GetMetadata copies core metadata from the LibRaw handle.
func (h *Handle) GetMetadata() Metadata {
	d := h.ptr
	return Metadata{
		Sizes:     convertImageSizes(&d.sizes),
		ID:        convertImageParams(&d.idata),
		Lens:      convertLensInfo(&d.lens),
		Shooting:  convertShootingInfo(&d.shootinginfo),
		Color:     convertColorData(&d.color),
		Other:     convertImageOther(&d.other),
		Thumbnail: convertThumbnail(&d.thumbnail),
		Thumbs:    convertThumbnailList(&d.thumbs_list),
		RawData:   convertRawDataSummary(&d.rawdata),
		Process:   ProcessState{ProgressFlags: uint32(d.progress_flags), ProcessWarnings: uint32(d.process_warnings)},
	}
}

func convertImageParams(p *C.libraw_iparams_t) ImageParams {
	out := ImageParams{
		Guard:           cString(&p.guard[0]),
		Make:            cString(&p.make[0]),
		Model:           cString(&p.model[0]),
		Software:        cString(&p.software[0]),
		NormalizedMake:  cString(&p.normalized_make[0]),
		NormalizedModel: cString(&p.normalized_model[0]),
		MakerIndex:      uint32(p.maker_index),
		RawCount:        uint32(p.raw_count),
		DNGVersion:      uint32(p.dng_version),
		IsFoveon:        uint32(p.is_foveon),
		Colors:          int(p.colors),
		Filters:         uint32(p.filters),
		CDesc:           cString(&p.cdesc[0]),
		XMLLen:          uint32(p.xmplen),
		HasXML:          p.xmpdata != nil,
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			out.XTrans[i][j] = int8(p.xtrans[i][j])
			out.XTransAbs[i][j] = int8(p.xtrans_abs[i][j])
		}
	}
	return out
}

func convertImageSizes(s *C.libraw_image_sizes_t) ImageSizes {
	out := ImageSizes{
		RawHeight:   uint16(s.raw_height),
		RawWidth:    uint16(s.raw_width),
		Height:      uint16(s.height),
		Width:       uint16(s.width),
		TopMargin:   uint16(s.top_margin),
		LeftMargin:  uint16(s.left_margin),
		IHeight:     uint16(s.iheight),
		IWidth:      uint16(s.iwidth),
		RawPitch:    uint32(s.raw_pitch),
		PixelAspect: float64(s.pixel_aspect),
		Flip:        int(s.flip),
		RawAspect:   uint16(s.raw_aspect),
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			out.Mask[i][j] = int(s.mask[i][j])
		}
	}
	for i := 0; i < 2; i++ {
		out.RawInsetCrops[i] = RawInsetCrop{
			Left:   uint16(s.raw_inset_crops[i].cleft),
			Top:    uint16(s.raw_inset_crops[i].ctop),
			Width:  uint16(s.raw_inset_crops[i].cwidth),
			Height: uint16(s.raw_inset_crops[i].cheight),
		}
	}
	return out
}

func convertInternalOutputParams(p *C.libraw_internal_output_params_t) InternalOutputParams {
	return InternalOutputParams{
		MixGreen:  uint32(p.mix_green),
		RawColor:  uint32(p.raw_color),
		ZeroIsBad: uint32(p.zero_is_bad),
		Shrink:    uint16(p.shrink),
		FujiWidth: uint16(p.fuji_width),
	}
}

func convertDNGColor(c *C.libraw_dng_color_t) DNGColor {
	out := DNGColor{ParsedFields: uint32(c.parsedfields), Illuminant: uint16(c.illuminant)}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			out.Calibration[i][j] = float32(c.calibration[i][j])
		}
		for j := 0; j < 3; j++ {
			out.ColorMatrix[i][j] = float32(c.colormatrix[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			out.ForwardMatrix[i][j] = float32(c.forwardmatrix[i][j])
		}
	}
	return out
}

func convertDNGLevels(l *C.libraw_dng_levels_t) DNGLevels {
	out := DNGLevels{
		ParsedFields:        uint32(l.parsedfields),
		DNGBlack:            uint32(l.dng_black),
		DNGFBlack:           float32(l.dng_fblack),
		PreviewColorspace:   uint32(l.preview_colorspace),
		BaselineExposure:    float32(l.baseline_exposure),
		LinearResponseLimit: float32(l.LinearResponseLimit),
	}
	for i := 0; i < 6; i++ {
		out.DNGCBlack[i] = uint32(l.dng_cblack[i])
		out.DNGFCBlack[i] = float32(l.dng_fcblack[i])
	}
	for i := 0; i < 4; i++ {
		out.DNGWhiteLevel[i] = uint32(l.dng_whitelevel[i])
		out.DefaultCrop[i] = uint16(l.default_crop[i])
		out.UserCrop[i] = float32(l.user_crop[i])
		out.AnalogBalance[i] = float32(l.analogbalance[i])
		out.AsShotNeutral[i] = float32(l.asshotneutral[i])
	}
	for i := 0; i < 3; i++ {
		out.RawOpcodes[i] = DNGRawOpcode{
			Len:     uint32(C.go_libraw_dng_rawopcode_len(l, C.int(i))),
			HasData: C.go_libraw_dng_rawopcode_has_data(l, C.int(i)) != 0,
		}
	}
	return out
}

func convertColorData(c *C.libraw_colordata_t) ColorData {
	out := ColorData{
		Black:                uint32(c.black),
		DataMaximum:          uint32(c.data_maximum),
		Maximum:              uint32(c.maximum),
		FMaximum:             float32(c.fmaximum),
		FNorm:                float32(c.fnorm),
		FlashUsed:            float32(c.flash_used),
		CanonEV:              float32(c.canon_ev),
		Model2:               cString(&c.model2[0]),
		UniqueCameraModel:    cString(&c.UniqueCameraModel[0]),
		LocalizedCameraModel: cString(&c.LocalizedCameraModel[0]),
		ImageUniqueID:        cString(&c.ImageUniqueID[0]),
		RawDataUniqueID:      cString(&c.RawDataUniqueID[0]),
		OriginalRawFileName:  cString(&c.OriginalRawFileName[0]),
		ProfileLength:        uint32(c.profile_length),
		HasProfile:           c.profile != nil,
		DNGLevels:            convertDNGLevels(&c.dng_levels),
		AsShotWBApplied:      int(c.as_shot_wb_applied),
		RawBPS:               uint32(c.raw_bps),
		ExifColorSpace:       int(c.ExifColorSpace),
	}
	for i := 0; i < 0x10000; i++ {
		out.Curve[i] = uint16(c.curve[i])
	}
	for i := 0; i < 6; i++ {
		out.CBlack[i] = uint32(c.cblack[i])
	}
	for i := 0; i < 4; i++ {
		out.LinearMax[i] = uint32(c.linear_max[i])
		out.CamMul[i] = float32(c.cam_mul[i])
		out.PreMul[i] = float32(c.pre_mul[i])
	}
	for i := 0; i < 8; i++ {
		out.BlackStat[i] = uint32(c.black_stat[i])
		for j := 0; j < 8; j++ {
			out.White[i][j] = uint16(c.white[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			out.CMatrix[i][j] = float32(c.cmatrix[i][j])
			out.CCM[i][j] = float32(c.ccm[i][j])
			out.RGBCam[i][j] = float32(c.rgb_cam[i][j])
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			out.CamXYZ[i][j] = float32(c.cam_xyz[i][j])
		}
	}
	out.PhaseOneData = PhaseOneData{
		Format:   int(c.phase_one_data.format),
		KeyOff:   int(c.phase_one_data.key_off),
		Tag21A:   int(c.phase_one_data.tag_21a),
		TBlack:   int(c.phase_one_data.t_black),
		SplitCol: int(c.phase_one_data.split_col),
		BlackCol: int(c.phase_one_data.black_col),
		SplitRow: int(c.phase_one_data.split_row),
		BlackRow: int(c.phase_one_data.black_row),
		Tag210:   float32(c.phase_one_data.tag_210),
	}
	for i := 0; i < 2; i++ {
		out.DNGColor[i] = convertDNGColor(&c.dng_color[i])
		for j := 0; j < 9; j++ {
			out.P1Color[i].ROMMCam[j] = float32(c.P1_color[i].romm_cam[j])
		}
	}
	for i := 0; i < 256; i++ {
		for j := 0; j < 4; j++ {
			out.WBCoeffs[i][j] = int(c.WB_Coeffs[i][j])
		}
	}
	for i := 0; i < 64; i++ {
		for j := 0; j < 5; j++ {
			out.WBCTCoeffs[i][j] = float32(c.WBCT_Coeffs[i][j])
		}
	}
	return out
}

func convertThumbnail(t *C.libraw_thumbnail_t) Thumbnail {
	return Thumbnail{
		Format:  int(t.tformat),
		Width:   uint16(t.twidth),
		Height:  uint16(t.theight),
		Length:  uint32(t.tlength),
		Colors:  int(t.tcolors),
		HasData: t.thumb != nil,
	}
}

func convertThumbnailList(l *C.libraw_thumbnail_list_t) ThumbnailList {
	count := int(l.thumbcount)
	if count < 0 {
		count = 0
	}
	if max := int(C.LIBRAW_THUMBNAIL_MAXCOUNT); count > max {
		count = max
	}
	out := ThumbnailList{Count: int(l.thumbcount), Items: make([]ThumbnailItem, 0, count)}
	for i := 0; i < count; i++ {
		item := l.thumblist[i]
		out.Items = append(out.Items, ThumbnailItem{
			Format: int(item.tformat),
			Width:  uint16(item.twidth),
			Height: uint16(item.theight),
			Flip:   uint16(item.tflip),
			Length: uint32(item.tlength),
			Misc:   uint32(item.tmisc),
			Offset: int64(item.toffset),
		})
	}
	return out
}

func convertGPSInfo(g *C.libraw_gps_info_t) GPSInfo {
	out := GPSInfo{
		Altitude: float32(g.altitude),
		AltRef:   int8(g.altref),
		LatRef:   int8(g.latref),
		LongRef:  int8(g.longref),
		Status:   int8(g.gpsstatus),
		Parsed:   int8(g.gpsparsed),
	}
	for i := 0; i < 3; i++ {
		out.Latitude[i] = float32(g.latitude[i])
		out.Longitude[i] = float32(g.longitude[i])
		out.Timestamp[i] = float32(g.gpstimestamp[i])
	}
	return out
}

func convertImageOther(o *C.libraw_imgother_t) ImageOther {
	out := ImageOther{
		ISOSpeed:    float32(o.iso_speed),
		Shutter:     float32(o.shutter),
		Aperture:    float32(o.aperture),
		FocalLen:    float32(o.focal_len),
		Timestamp:   int64(o.timestamp),
		ShotOrder:   uint32(o.shot_order),
		ParsedGPS:   convertGPSInfo(&o.parsed_gps),
		Description: cString(&o.desc[0]),
		Artist:      cString(&o.artist[0]),
	}
	for i := 0; i < 32; i++ {
		out.GPSData[i] = uint32(o.gpsdata[i])
	}
	for i := 0; i < 4; i++ {
		out.AnalogBalance[i] = float32(o.analogbalance[i])
	}
	return out
}

func convertLensInfo(l *C.libraw_lensinfo_t) LensInfo {
	return LensInfo{
		MinFocal:                float32(l.MinFocal),
		MaxFocal:                float32(l.MaxFocal),
		MaxAp4MinFocal:          float32(l.MaxAp4MinFocal),
		MaxAp4MaxFocal:          float32(l.MaxAp4MaxFocal),
		EXIFMaxAp:               float32(l.EXIF_MaxAp),
		LensMake:                cString(&l.LensMake[0]),
		Lens:                    cString(&l.Lens[0]),
		LensSerial:              cString(&l.LensSerial[0]),
		InternalLensSerial:      cString(&l.InternalLensSerial[0]),
		FocalLengthIn35mmFormat: uint16(l.FocalLengthIn35mmFormat),
		Nikon: NikonLens{
			EffectiveMaxAp: float32(l.nikon.EffectiveMaxAp),
			LensIDNumber:   uint8(l.nikon.LensIDNumber),
			LensFStops:     uint8(l.nikon.LensFStops),
			MCUVersion:     uint8(l.nikon.MCUVersion),
			LensType:       uint8(l.nikon.LensType),
		},
		DNG: DNGLens{
			MinFocal:       float32(l.dng.MinFocal),
			MaxFocal:       float32(l.dng.MaxFocal),
			MaxAp4MinFocal: float32(l.dng.MaxAp4MinFocal),
			MaxAp4MaxFocal: float32(l.dng.MaxAp4MaxFocal),
		},
		MakerNotes: convertMakerNotesLens(&l.makernotes),
	}
}

func convertMakerNotesLens(l *C.libraw_makernotes_lens_t) MakerNotesLens {
	return MakerNotesLens{
		LensID:                  uint64(l.LensID),
		Lens:                    cString(&l.Lens[0]),
		LensFormat:              uint16(l.LensFormat),
		LensMount:               uint16(l.LensMount),
		CamID:                   uint64(l.CamID),
		CameraFormat:            uint16(l.CameraFormat),
		CameraMount:             uint16(l.CameraMount),
		Body:                    cString(&l.body[0]),
		FocalType:               int16(l.FocalType),
		LensFeaturesPre:         cString(&l.LensFeatures_pre[0]),
		LensFeaturesSuf:         cString(&l.LensFeatures_suf[0]),
		MinFocal:                float32(l.MinFocal),
		MaxFocal:                float32(l.MaxFocal),
		MaxAp4MinFocal:          float32(l.MaxAp4MinFocal),
		MaxAp4MaxFocal:          float32(l.MaxAp4MaxFocal),
		MinAp4MinFocal:          float32(l.MinAp4MinFocal),
		MinAp4MaxFocal:          float32(l.MinAp4MaxFocal),
		MaxAp:                   float32(l.MaxAp),
		MinAp:                   float32(l.MinAp),
		CurFocal:                float32(l.CurFocal),
		CurAp:                   float32(l.CurAp),
		MaxAp4CurFocal:          float32(l.MaxAp4CurFocal),
		MinAp4CurFocal:          float32(l.MinAp4CurFocal),
		MinFocusDistance:        float32(l.MinFocusDistance),
		FocusRangeIndex:         float32(l.FocusRangeIndex),
		LensFStops:              float32(l.LensFStops),
		TeleconverterID:         uint64(l.TeleconverterID),
		Teleconverter:           cString(&l.Teleconverter[0]),
		AdapterID:               uint64(l.AdapterID),
		Adapter:                 cString(&l.Adapter[0]),
		AttachmentID:            uint64(l.AttachmentID),
		Attachment:              cString(&l.Attachment[0]),
		FocalUnits:              uint16(l.FocalUnits),
		FocalLengthIn35mmFormat: float32(l.FocalLengthIn35mmFormat),
	}
}

func convertShootingInfo(s *C.libraw_shootinginfo_t) ShootingInfo {
	return ShootingInfo{
		DriveMode:          int16(s.DriveMode),
		FocusMode:          int16(s.FocusMode),
		MeteringMode:       int16(s.MeteringMode),
		AFPoint:            int16(s.AFPoint),
		ExposureMode:       int16(s.ExposureMode),
		ExposureProgram:    int16(s.ExposureProgram),
		ImageStabilization: int16(s.ImageStabilization),
		BodySerial:         cString(&s.BodySerial[0]),
		InternalBodySerial: cString(&s.InternalBodySerial[0]),
	}
}

func convertRawDataSummary(r *C.libraw_rawdata_t) RawDataSummary {
	return RawDataSummary{
		HasRawAlloc:    r.raw_alloc != nil,
		HasRawImage:    r.raw_image != nil,
		HasColor4Image: r.color4_image != nil,
		HasColor3Image: r.color3_image != nil,
		HasFloatImage:  r.float_image != nil,
		HasFloat3Image: r.float3_image != nil,
		HasFloat4Image: r.float4_image != nil,
		HasPH1CBlack:   r.ph1_cblack != nil,
		HasPH1RBlack:   r.ph1_rblack != nil,
		ImageParams:    convertImageParams(&r.iparams),
		Sizes:          convertImageSizes(&r.sizes),
		InternalOutput: convertInternalOutputParams(&r.ioparams),
		Color:          convertColorData(&r.color),
	}
}

func cString(p *C.char) string {
	return C.GoString(p)
}

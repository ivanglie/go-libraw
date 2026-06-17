//go:build cgo

package librawc

/*
#include <libraw/libraw.h>

static short go_libraw_canon_auto_rotate_mode(libraw_canon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->AutoRotateMode;
#else
	(void)p;
	return 0;
#endif
}

static const char *go_libraw_fuji_model(libraw_fuji_info_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->FujiModel;
#else
	(void)p;
	return "";
#endif
}

static const char *go_libraw_fuji_model2(libraw_fuji_info_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->FujiModel2;
#else
	(void)p;
	return "";
#endif
}

static unsigned go_libraw_nikon_picture_control_version(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->PictureControlVersion;
#else
	(void)p;
	return 0;
#endif
}

static const char *go_libraw_nikon_picture_control_name(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->PictureControlName;
#else
	(void)p;
	return "";
#endif
}

static const char *go_libraw_nikon_picture_control_base(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->PictureControlBase;
#else
	(void)p;
	return "";
#endif
}

static const char *go_libraw_nikon_shot_info_firmware(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->ShotInfoFirmware;
#else
	(void)p;
	return "";
#endif
}

static unsigned go_libraw_nikon_burst_table_len(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->BurstTable_0x0056_len;
#else
	(void)p;
	return 0;
#endif
}

static int go_libraw_nikon_burst_table_has_data(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->BurstTable_0x0056 != 0;
#else
	(void)p;
	return 0;
#endif
}

static unsigned short go_libraw_nikon_burst_table_ver(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->BurstTable_0x0056_ver;
#else
	(void)p;
	return 0;
#endif
}

static unsigned short go_libraw_nikon_burst_table_gid(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->BurstTable_0x0056_gid;
#else
	(void)p;
	return 0;
#endif
}

static unsigned char go_libraw_nikon_burst_table_fnum(libraw_nikon_makernotes_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->BurstTable_0x0056_fnum;
#else
	(void)p;
	return 0;
#endif
}

static unsigned go_libraw_olympus_decoder_tag(libraw_olympus_makernotes_t *p, int i) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	switch(i) {
	case 0: return p->tagX640;
	case 1: return p->tagX641;
	case 2: return p->tagX642;
	case 3: return p->tagX643;
	case 4: return p->tagX644;
	case 5: return p->tagX645;
	case 6: return p->tagX646;
	case 7: return p->tagX647;
	case 8: return p->tagX648;
	case 9: return p->tagX649;
	case 10: return p->tagX650;
	case 11: return p->tagX651;
	case 12: return p->tagX652;
	case 13: return p->tagX653;
	default: return 0;
	}
#else
	(void)p;
	(void)i;
	return 0;
#endif
}

static unsigned short go_libraw_sony_len_group9050(libraw_sony_info_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->len_group9050;
#else
	(void)p;
	return 0;
#endif
}

static float go_libraw_sony_aspect_ratio(libraw_sony_info_t *p) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->AspectRatio;
#else
	(void)p;
	return 0;
#endif
}

static unsigned char go_libraw_pentax_dynamic_range_expansion(libraw_pentax_makernotes_t *p, int i) {
#if LIBRAW_VERSION >= LIBRAW_MAKE_VERSION(0,22,0)
	return p->DynamicRangeExpansion[i];
#else
	(void)p;
	(void)i;
	return 0;
#endif
}
*/
import "C"

// MakerNotes mirrors libraw_makernotes_t.
type MakerNotes struct {
	Canon      CanonMakerNotes
	Nikon      NikonMakerNotes
	Hasselblad HasselbladMakerNotes
	Fuji       FujiMakerNotes
	Olympus    OlympusMakerNotes
	Sony       SonyMakerNotes
	Kodak      KodakMakerNotes
	Panasonic  PanasonicMakerNotes
	Pentax     PentaxMakerNotes
	PhaseOne   PhaseOneMakerNotes
	Ricoh      RicohMakerNotes
	Samsung    SamsungMakerNotes
	Common     MetadataCommon
}

// Area mirrors libraw_area_t.
type Area struct {
	Top    int16
	Left   int16
	Bottom int16
	Right  int16
}

// SensorHighSpeedCrop mirrors libraw_sensor_highspeed_crop_t.
type SensorHighSpeedCrop struct {
	Left   uint16
	Top    uint16
	Width  uint16
	Height uint16
}

// AFInfoItem mirrors libraw_afinfo_item_t without exposing the C data pointer.
type AFInfoItem struct {
	Tag     uint32
	Order   int16
	Version uint32
	Length  uint32
	HasData bool
}

// MetadataCommon mirrors libraw_metadata_common_t.
type MetadataCommon struct {
	FlashEC                  float32
	FlashGN                  float32
	CameraTemperature        float32
	SensorTemperature        float32
	SensorTemperature2       float32
	LensTemperature          float32
	AmbientTemperature       float32
	BatteryTemperature       float32
	EXIFAmbientTemperature   float32
	EXIFHumidity             float32
	EXIFPressure             float32
	EXIFWaterDepth           float32
	EXIFAcceleration         float32
	EXIFCameraElevationAngle float32
	RealISO                  float32
	EXIFExposureIndex        float32
	ColorSpace               uint16
	Firmware                 string
	ExposureCalibrationShift float32
	AFData                   []AFInfoItem
	AFCount                  int
}

// CanonMakerNotes mirrors libraw_canon_makernotes_t.
type CanonMakerNotes struct {
	ColorDataVer          int
	ColorDataSubVer       int
	SpecularWhiteLevel    int
	NormalWhiteLevel      int
	ChannelBlackLevel     [4]int
	AverageBlackLevel     int
	Multishot             [4]uint32
	MeteringMode          int16
	SpotMeteringMode      int16
	FlashMeteringMode     uint8
	FlashExposureLock     int16
	ExposureMode          int16
	AESetting             int16
	ImageStabilization    int16
	FlashMode             int16
	FlashActivity         int16
	FlashBits             int16
	ManualFlashOutput     int16
	FlashOutput           int16
	FlashGuideNumber      int16
	ContinuousDrive       int16
	SensorWidth           int16
	SensorHeight          int16
	AFMicroAdjMode        int
	AFMicroAdjValue       float32
	MakernotesFlip        int16
	AutoRotateMode        int16
	RecordMode            int16
	SRAWQuality           int16
	WBI                   uint32
	RFLensID              int16
	AutoLightingOptimizer int
	HighlightTonePriority int
	Quality               int16
	CanonLog              int
	DefaultCropAbsolute   Area
	RecommendedImageArea  Area
	LeftOpticalBlack      Area
	UpperOpticalBlack     Area
	ActiveArea            Area
	ISOGain               [2]int16
}

// HasselbladMakerNotes mirrors libraw_hasselblad_makernotes_t.
type HasselbladMakerNotes struct {
	BaseISO                  int
	Gain                     float64
	Sensor                   string
	SensorUnit               string
	HostBody                 string
	SensorCode               int
	SensorSubCode            int
	CoatingCode              int
	Uncropped                int
	CaptureSequenceInitiator string
	SensorUnitConnector      string
	Format                   int
	NIFDCM                   [2]int
	RecommendedCrop          [2]int
	MNColorMatrix            [4][3]float64
}

// FujiMakerNotes mirrors libraw_fuji_info_t.
type FujiMakerNotes struct {
	ExpoMidPointShift       float32
	DynamicRange            uint16
	FilmMode                uint16
	DynamicRangeSetting     uint16
	DevelopmentDynamicRange uint16
	AutoDynamicRange        uint16
	DRangePriority          uint16
	DRangePriorityAuto      uint16
	DRangePriorityFixed     uint16
	FujiModel               string
	FujiModel2              string
	BrightnessCompensation  float32
	FocusMode               uint16
	AFMode                  uint16
	FocusPixel              [2]uint16
	PrioritySettings        uint16
	FocusSettings           uint32
	AFCSettings             uint32
	FocusWarning            uint16
	ImageStabilization      [3]uint16
	FlashMode               uint16
	WBPreset                uint16
	ShutterType             uint16
	ExrMode                 uint16
	Macro                   uint16
	Rating                  uint32
	CropMode                uint16
	SerialSignature         string
	SensorID                string
	RAFVersion              string
	RAFDataGeneration       int
	RAFDataVersion          uint16
	IsTSNERDTS              int
	DriveMode               int16
	BlackLevel              [9]uint16
	RAFDataImageSizeTable   [32]uint32
	AutoBracketing          int
	SequenceNumber          int
	SeriesLength            int
	PixelShiftOffset        [2]float32
	ImageCount              int
}

// NikonMakerNotes mirrors libraw_nikon_makernotes_t.
type NikonMakerNotes struct {
	ExposureBracketValue            float64
	ActiveDLighting                 uint16
	ShootingMode                    uint16
	ImageStabilization              [7]uint8
	VibrationReduction              uint8
	VRMode                          uint8
	FlashSetting                    string
	FlashType                       string
	FlashExposureCompensation       [4]uint8
	ExternalFlashExposureComp       [4]uint8
	FlashExposureBracketValue       [4]uint8
	FlashMode                       uint8
	FlashExposureCompensation2      int8
	FlashExposureCompensation3      int8
	FlashExposureCompensation4      int8
	FlashSource                     uint8
	FlashFirmware                   [2]uint8
	ExternalFlashFlags              uint8
	FlashControlCommanderMode       uint8
	FlashOutputAndCompensation      uint8
	FlashFocalLength                uint8
	FlashGNDistance                 uint8
	FlashGroupControlMode           [4]uint8
	FlashGroupOutputAndCompensation [4]uint8
	FlashColorFilter                uint8
	NEFCompression                  uint16
	ExposureMode                    int
	ExposureProgram                 int
	NMEShots                        int
	MEGainOn                        int
	MEWB                            [4]float64
	AFFineTune                      uint8
	AFFineTuneIndex                 uint8
	AFFineTuneAdj                   int8
	LensDataVersion                 uint32
	FlashInfoVersion                uint32
	ColorBalanceVersion             uint32
	Key                             uint8
	NEFBitDepth                     [4]uint16
	HighSpeedCropFormat             uint16
	SensorHighSpeedCrop             SensorHighSpeedCrop
	SensorWidth                     uint16
	SensorHeight                    uint16
	ActiveDLighting2                uint16
	PictureControlVersion           uint32
	PictureControlName              string
	PictureControlBase              string
	ShotInfoVersion                 uint32
	ShotInfoFirmware                string
	BurstTable0056Len               uint32
	HasBurstTable0056               bool
	BurstTable0056Ver               uint16
	BurstTable0056GID               uint16
	BurstTable0056FNum              uint8
	MakernotesFlip                  int16
	RollAngle                       float64
	PitchAngle                      float64
	YawAngle                        float64
}

// OlympusMakerNotes mirrors libraw_olympus_makernotes_t.
type OlympusMakerNotes struct {
	CameraType2       string
	ValidBits         uint16
	DecoderTags       [14]uint32
	SensorCalibration [2]int
	DriveMode         [5]uint16
	ColorSpace        uint16
	FocusMode         [2]uint16
	AutoFocus         uint16
	AFPoint           uint16
	AFAreas           [64]uint32
	AFPointSelected   [5]float64
	AFResult          uint16
	AFFineTune        uint8
	AFFineTuneAdj     [3]int16
	SpecialMode       [3]uint32
	ZoomStepCount     uint16
	FocusStepCount    uint16
	FocusStepInfinity uint16
	FocusStepNear     uint16
	FocusDistance     float64
	AspectFrame       [4]uint16
	StackedImage      [2]uint32
	IsLiveND          uint8
	LiveNDFactor      uint32
	PanoramaMode      uint16
	PanoramaFrameNum  uint16
}

// SonyMakerNotes mirrors libraw_sony_info_t.
type SonyMakerNotes struct {
	CameraType                    uint16
	Sony0x9400Version             uint8
	Sony0x9400ReleaseMode2        uint8
	Sony0x9400SequenceImageNumber uint32
	Sony0x9400SequenceLength1     uint8
	Sony0x9400SequenceFileNumber  uint32
	Sony0x9400SequenceLength2     uint8
	AFAreaModeSetting             uint8
	AFAreaMode                    uint16
	FlexibleSpotPosition          [2]uint16
	AFPointSelected               uint8
	AFPointSelected0x201e         uint8
	NAFPointsUsed                 int16
	AFPointsUsed                  [10]uint8
	AFTracking                    uint8
	AFType                        uint8
	FocusLocation                 [4]uint16
	FocusPosition                 uint16
	AFMicroAdjValue               int8
	AFMicroAdjOn                  int8
	AFMicroAdjRegisteredLenses    uint8
	VariableLowPassFilter         uint16
	LongExposureNoiseReduction    uint32
	HighISONoiseReduction         uint16
	HDR                           [2]uint16
	Group2010                     uint16
	Group9050                     uint16
	LenGroup9050                  uint16
	RealISOOffset                 uint16
	MeteringModeOffset            uint16
	ExposureProgramOffset         uint16
	ReleaseMode2Offset            uint16
	MinoltaCamID                  uint32
	Firmware                      float32
	ImageCount3Offset             uint16
	ImageCount3                   uint32
	ElectronicFrontCurtainShutter uint32
	MeteringMode2                 uint16
	SonyDateTime                  string
	ShotNumberSincePowerUp        uint32
	PixelShiftGroupPrefix         uint16
	PixelShiftGroupID             uint32
	NShotsInPixelShiftGroup       int8
	NumInPixelShiftGroup          int8
	PRDImageHeight                uint16
	PRDImageWidth                 uint16
	PRDTotalBPS                   uint16
	PRDActiveBPS                  uint16
	PRDStorageMethod              uint16
	PRDBayerPattern               uint16
	SonyRawFileType               uint16
	RawFileType                   uint16
	RawSizeType                   uint16
	Quality                       uint32
	FileFormat                    uint16
	MetaVersion                   string
	AspectRatio                   float32
}

// KodakMakerNotes mirrors libraw_kodak_makernotes_t.
type KodakMakerNotes struct {
	BlackLevelTop      uint16
	BlackLevelBottom   uint16
	OffsetLeft         int16
	OffsetTop          int16
	ClipBlack          uint16
	ClipWhite          uint16
	ROMMCamDaylight    [3][3]float32
	ROMMCamTungsten    [3][3]float32
	ROMMCamFluorescent [3][3]float32
	ROMMCamFlash       [3][3]float32
	ROMMCamCustom      [3][3]float32
	ROMMCamAuto        [3][3]float32
	Val018Percent      uint16
	Val100Percent      uint16
	Val170Percent      uint16
	MakerNoteKodak8a   int16
	ISOCalibrationGain float32
	AnalogISO          float32
}

// PanasonicMakerNotes mirrors libraw_panasonic_makernotes_t.
type PanasonicMakerNotes struct {
	Compression       uint16
	BlackLevelDim     uint16
	BlackLevel        [8]float32
	Multishot         uint32
	Gamma             float32
	HighISOMultiplier [3]int
	FocusStepNear     int16
	FocusStepCount    int16
	ZoomPosition      uint32
	LensManufacturer  uint32
}

// PentaxMakerNotes mirrors libraw_pentax_makernotes_t.
type PentaxMakerNotes struct {
	DriveMode              [4]uint8
	FocusMode              [2]uint16
	AFPointSelected        [2]uint16
	AFPointSelectedArea    uint16
	AFPointsInFocusVersion int
	AFPointsInFocus        uint32
	FocusPosition          uint16
	DynamicRangeExpansion  [4]uint8
	AFAdjustment           int16
	AFPointMode            uint8
	MultiExposure          uint8
	Quality                uint16
}

// PhaseOneMakerNotes mirrors libraw_p1_makernotes_t.
type PhaseOneMakerNotes struct {
	Software       string
	SystemType     string
	FirmwareString string
	SystemModel    string
}

// RicohMakerNotes mirrors libraw_ricoh_makernotes_t.
type RicohMakerNotes struct {
	AFStatus           uint16
	AFAreaXPosition    [2]uint32
	AFAreaYPosition    [2]uint32
	AFAreaMode         uint16
	SensorWidth        uint32
	SensorHeight       uint32
	CroppedImageWidth  uint32
	CroppedImageHeight uint32
	WideAdapter        uint16
	CropMode           uint16
	NDFilter           uint16
	AutoBracketing     uint16
	MacroMode          uint16
	FlashMode          uint16
	FlashExposureComp  float64
	ManualFlashOutput  float64
}

// SamsungMakerNotes mirrors libraw_samsung_makernotes_t.
type SamsungMakerNotes struct {
	ImageSizeFull [4]uint32
	ImageSizeCrop [4]uint32
	ColorSpace    [2]int
	Key           [11]uint32
	DigitalGain   float64
	DeviceType    int
	LensFirmware  string
}

func convertMakerNotes(m *C.libraw_makernotes_t) MakerNotes {
	return MakerNotes{
		Canon:      convertCanonMakerNotes(&m.canon),
		Nikon:      convertNikonMakerNotes(&m.nikon),
		Hasselblad: convertHasselbladMakerNotes(&m.hasselblad),
		Fuji:       convertFujiMakerNotes(&m.fuji),
		Olympus:    convertOlympusMakerNotes(&m.olympus),
		Sony:       convertSonyMakerNotes(&m.sony),
		Kodak:      convertKodakMakerNotes(&m.kodak),
		Panasonic:  convertPanasonicMakerNotes(&m.panasonic),
		Pentax:     convertPentaxMakerNotes(&m.pentax),
		PhaseOne:   convertPhaseOneMakerNotes(&m.phaseone),
		Ricoh:      convertRicohMakerNotes(&m.ricoh),
		Samsung:    convertSamsungMakerNotes(&m.samsung),
		Common:     convertMetadataCommon(&m.common),
	}
}

func convertArea(a *C.libraw_area_t) Area {
	return Area{Top: int16(a.t), Left: int16(a.l), Bottom: int16(a.b), Right: int16(a.r)}
}

func convertSensorHighSpeedCrop(c *C.libraw_sensor_highspeed_crop_t) SensorHighSpeedCrop {
	return SensorHighSpeedCrop{Left: uint16(c.cleft), Top: uint16(c.ctop), Width: uint16(c.cwidth), Height: uint16(c.cheight)}
}

func convertMetadataCommon(c *C.libraw_metadata_common_t) MetadataCommon {
	count := int(c.afcount)
	if count < 0 {
		count = 0
	}
	if max := int(C.LIBRAW_AFDATA_MAXCOUNT); count > max {
		count = max
	}
	out := MetadataCommon{
		FlashEC: float32(c.FlashEC), FlashGN: float32(c.FlashGN),
		CameraTemperature: float32(c.CameraTemperature), SensorTemperature: float32(c.SensorTemperature),
		SensorTemperature2: float32(c.SensorTemperature2), LensTemperature: float32(c.LensTemperature),
		AmbientTemperature: float32(c.AmbientTemperature), BatteryTemperature: float32(c.BatteryTemperature),
		EXIFAmbientTemperature: float32(c.exifAmbientTemperature), EXIFHumidity: float32(c.exifHumidity),
		EXIFPressure: float32(c.exifPressure), EXIFWaterDepth: float32(c.exifWaterDepth),
		EXIFAcceleration: float32(c.exifAcceleration), EXIFCameraElevationAngle: float32(c.exifCameraElevationAngle),
		RealISO: float32(c.real_ISO), EXIFExposureIndex: float32(c.exifExposureIndex),
		ColorSpace: uint16(c.ColorSpace), Firmware: cString(&c.firmware[0]),
		ExposureCalibrationShift: float32(c.ExposureCalibrationShift), AFCount: int(c.afcount),
		AFData: make([]AFInfoItem, 0, count),
	}
	for i := 0; i < count; i++ {
		item := c.afdata[i]
		out.AFData = append(out.AFData, AFInfoItem{
			Tag: uint32(item.AFInfoData_tag), Order: int16(item.AFInfoData_order),
			Version: uint32(item.AFInfoData_version), Length: uint32(item.AFInfoData_length),
			HasData: item.AFInfoData != nil,
		})
	}
	return out
}

func convertCanonMakerNotes(c *C.libraw_canon_makernotes_t) CanonMakerNotes {
	out := CanonMakerNotes{
		ColorDataVer: int(c.ColorDataVer), ColorDataSubVer: int(c.ColorDataSubVer),
		SpecularWhiteLevel: int(c.SpecularWhiteLevel), NormalWhiteLevel: int(c.NormalWhiteLevel),
		AverageBlackLevel: int(c.AverageBlackLevel), MeteringMode: int16(c.MeteringMode),
		SpotMeteringMode: int16(c.SpotMeteringMode), FlashMeteringMode: uint8(c.FlashMeteringMode),
		FlashExposureLock: int16(c.FlashExposureLock), ExposureMode: int16(c.ExposureMode),
		AESetting: int16(c.AESetting), ImageStabilization: int16(c.ImageStabilization),
		FlashMode: int16(c.FlashMode), FlashActivity: int16(c.FlashActivity), FlashBits: int16(c.FlashBits),
		ManualFlashOutput: int16(c.ManualFlashOutput), FlashOutput: int16(c.FlashOutput),
		FlashGuideNumber: int16(c.FlashGuideNumber), ContinuousDrive: int16(c.ContinuousDrive),
		SensorWidth: int16(c.SensorWidth), SensorHeight: int16(c.SensorHeight),
		AFMicroAdjMode: int(c.AFMicroAdjMode), AFMicroAdjValue: float32(c.AFMicroAdjValue),
		MakernotesFlip: int16(c.MakernotesFlip), AutoRotateMode: int16(C.go_libraw_canon_auto_rotate_mode(c)),
		RecordMode: int16(c.RecordMode), SRAWQuality: int16(c.SRAWQuality), WBI: uint32(c.wbi),
		RFLensID: int16(c.RF_lensID), AutoLightingOptimizer: int(c.AutoLightingOptimizer),
		HighlightTonePriority: int(c.HighlightTonePriority), Quality: int16(c.Quality), CanonLog: int(c.CanonLog),
		DefaultCropAbsolute: convertArea(&c.DefaultCropAbsolute), RecommendedImageArea: convertArea(&c.RecommendedImageArea),
		LeftOpticalBlack: convertArea(&c.LeftOpticalBlack), UpperOpticalBlack: convertArea(&c.UpperOpticalBlack),
		ActiveArea: convertArea(&c.ActiveArea),
	}
	for i := 0; i < 4; i++ {
		out.ChannelBlackLevel[i] = int(c.ChannelBlackLevel[i])
		out.Multishot[i] = uint32(c.multishot[i])
	}
	for i := 0; i < 2; i++ {
		out.ISOGain[i] = int16(c.ISOgain[i])
	}
	return out
}

func convertHasselbladMakerNotes(h *C.libraw_hasselblad_makernotes_t) HasselbladMakerNotes {
	out := HasselbladMakerNotes{
		BaseISO: int(h.BaseISO), Gain: float64(h.Gain), Sensor: cString(&h.Sensor[0]),
		SensorUnit: cString(&h.SensorUnit[0]), HostBody: cString(&h.HostBody[0]),
		SensorCode: int(h.SensorCode), SensorSubCode: int(h.SensorSubCode), CoatingCode: int(h.CoatingCode),
		Uncropped: int(h.uncropped), CaptureSequenceInitiator: cString(&h.CaptureSequenceInitiator[0]),
		SensorUnitConnector: cString(&h.SensorUnitConnector[0]), Format: int(h.format),
	}
	for i := 0; i < 2; i++ {
		out.NIFDCM[i] = int(h.nIFD_CM[i])
		out.RecommendedCrop[i] = int(h.RecommendedCrop[i])
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			out.MNColorMatrix[i][j] = float64(h.mnColorMatrix[i][j])
		}
	}
	return out
}

func convertFujiMakerNotes(f *C.libraw_fuji_info_t) FujiMakerNotes {
	out := FujiMakerNotes{
		ExpoMidPointShift: float32(f.ExpoMidPointShift), DynamicRange: uint16(f.DynamicRange),
		FilmMode: uint16(f.FilmMode), DynamicRangeSetting: uint16(f.DynamicRangeSetting),
		DevelopmentDynamicRange: uint16(f.DevelopmentDynamicRange), AutoDynamicRange: uint16(f.AutoDynamicRange),
		DRangePriority: uint16(f.DRangePriority), DRangePriorityAuto: uint16(f.DRangePriorityAuto),
		DRangePriorityFixed: uint16(f.DRangePriorityFixed), FujiModel: C.GoString(C.go_libraw_fuji_model(f)),
		FujiModel2: C.GoString(C.go_libraw_fuji_model2(f)), BrightnessCompensation: float32(f.BrightnessCompensation),
		FocusMode: uint16(f.FocusMode), AFMode: uint16(f.AFMode), PrioritySettings: uint16(f.PrioritySettings),
		FocusSettings: uint32(f.FocusSettings), AFCSettings: uint32(f.AF_C_Settings), FocusWarning: uint16(f.FocusWarning),
		FlashMode: uint16(f.FlashMode), WBPreset: uint16(f.WB_Preset), ShutterType: uint16(f.ShutterType),
		ExrMode: uint16(f.ExrMode), Macro: uint16(f.Macro), Rating: uint32(f.Rating), CropMode: uint16(f.CropMode),
		SerialSignature: cString(&f.SerialSignature[0]), SensorID: cString(&f.SensorID[0]), RAFVersion: cString(&f.RAFVersion[0]),
		RAFDataGeneration: int(f.RAFDataGeneration), RAFDataVersion: uint16(f.RAFDataVersion), IsTSNERDTS: int(f.isTSNERDTS),
		DriveMode: int16(f.DriveMode), AutoBracketing: int(f.AutoBracketing), SequenceNumber: int(f.SequenceNumber),
		SeriesLength: int(f.SeriesLength), ImageCount: int(f.ImageCount),
	}
	for i := 0; i < 2; i++ {
		out.FocusPixel[i] = uint16(f.FocusPixel[i])
		out.PixelShiftOffset[i] = float32(f.PixelShiftOffset[i])
	}
	for i := 0; i < 3; i++ {
		out.ImageStabilization[i] = uint16(f.ImageStabilization[i])
	}
	for i := 0; i < 9; i++ {
		out.BlackLevel[i] = uint16(f.BlackLevel[i])
	}
	for i := 0; i < 32; i++ {
		out.RAFDataImageSizeTable[i] = uint32(f.RAFData_ImageSizeTable[i])
	}
	return out
}

func convertNikonMakerNotes(n *C.libraw_nikon_makernotes_t) NikonMakerNotes {
	out := NikonMakerNotes{
		ExposureBracketValue: float64(n.ExposureBracketValue), ActiveDLighting: uint16(n.ActiveDLighting),
		ShootingMode: uint16(n.ShootingMode), VibrationReduction: uint8(n.VibrationReduction), VRMode: uint8(n.VRMode),
		FlashSetting: cString(&n.FlashSetting[0]), FlashType: cString(&n.FlashType[0]), FlashMode: uint8(n.FlashMode),
		FlashExposureCompensation2: int8(n.FlashExposureCompensation2), FlashExposureCompensation3: int8(n.FlashExposureCompensation3),
		FlashExposureCompensation4: int8(n.FlashExposureCompensation4), FlashSource: uint8(n.FlashSource),
		ExternalFlashFlags: uint8(n.ExternalFlashFlags), FlashControlCommanderMode: uint8(n.FlashControlCommanderMode),
		FlashOutputAndCompensation: uint8(n.FlashOutputAndCompensation), FlashFocalLength: uint8(n.FlashFocalLength),
		FlashGNDistance: uint8(n.FlashGNDistance), FlashColorFilter: uint8(n.FlashColorFilter), NEFCompression: uint16(n.NEFCompression),
		ExposureMode: int(n.ExposureMode), ExposureProgram: int(n.ExposureProgram), NMEShots: int(n.nMEshots), MEGainOn: int(n.MEgainOn),
		AFFineTune: uint8(n.AFFineTune), AFFineTuneIndex: uint8(n.AFFineTuneIndex), AFFineTuneAdj: int8(n.AFFineTuneAdj),
		LensDataVersion: uint32(n.LensDataVersion), FlashInfoVersion: uint32(n.FlashInfoVersion), ColorBalanceVersion: uint32(n.ColorBalanceVersion),
		Key: uint8(n.key), HighSpeedCropFormat: uint16(n.HighSpeedCropFormat), SensorHighSpeedCrop: convertSensorHighSpeedCrop(&n.SensorHighSpeedCrop),
		SensorWidth: uint16(n.SensorWidth), SensorHeight: uint16(n.SensorHeight), ActiveDLighting2: uint16(n.Active_D_Lighting),
		PictureControlVersion: uint32(C.go_libraw_nikon_picture_control_version(n)), PictureControlName: C.GoString(C.go_libraw_nikon_picture_control_name(n)),
		PictureControlBase: C.GoString(C.go_libraw_nikon_picture_control_base(n)), ShotInfoVersion: uint32(n.ShotInfoVersion),
		ShotInfoFirmware: C.GoString(C.go_libraw_nikon_shot_info_firmware(n)), BurstTable0056Len: uint32(C.go_libraw_nikon_burst_table_len(n)),
		HasBurstTable0056: C.go_libraw_nikon_burst_table_has_data(n) != 0, BurstTable0056Ver: uint16(C.go_libraw_nikon_burst_table_ver(n)),
		BurstTable0056GID: uint16(C.go_libraw_nikon_burst_table_gid(n)), BurstTable0056FNum: uint8(C.go_libraw_nikon_burst_table_fnum(n)),
		MakernotesFlip: int16(n.MakernotesFlip), RollAngle: float64(n.RollAngle), PitchAngle: float64(n.PitchAngle), YawAngle: float64(n.YawAngle),
	}
	for i := 0; i < 7; i++ {
		out.ImageStabilization[i] = uint8(n.ImageStabilization[i])
	}
	for i := 0; i < 4; i++ {
		out.FlashExposureCompensation[i] = uint8(n.FlashExposureCompensation[i])
		out.ExternalFlashExposureComp[i] = uint8(n.ExternalFlashExposureComp[i])
		out.FlashExposureBracketValue[i] = uint8(n.FlashExposureBracketValue[i])
		out.FlashGroupControlMode[i] = uint8(n.FlashGroupControlMode[i])
		out.FlashGroupOutputAndCompensation[i] = uint8(n.FlashGroupOutputAndCompensation[i])
		out.MEWB[i] = float64(n.ME_WB[i])
		out.NEFBitDepth[i] = uint16(n.NEFBitDepth[i])
	}
	for i := 0; i < 2; i++ {
		out.FlashFirmware[i] = uint8(n.FlashFirmware[i])
	}
	return out
}

func convertOlympusMakerNotes(o *C.libraw_olympus_makernotes_t) OlympusMakerNotes {
	out := OlympusMakerNotes{
		CameraType2: cString(&o.CameraType2[0]), ValidBits: uint16(o.ValidBits),
		ColorSpace: uint16(o.ColorSpace), AutoFocus: uint16(o.AutoFocus), AFPoint: uint16(o.AFPoint),
		AFResult: uint16(o.AFResult), AFFineTune: uint8(o.AFFineTune), ZoomStepCount: uint16(o.ZoomStepCount),
		FocusStepCount: uint16(o.FocusStepCount), FocusStepInfinity: uint16(o.FocusStepInfinity), FocusStepNear: uint16(o.FocusStepNear),
		FocusDistance: float64(o.FocusDistance), IsLiveND: uint8(o.isLiveND), LiveNDFactor: uint32(o.LiveNDfactor),
		PanoramaMode: uint16(o.Panorama_mode), PanoramaFrameNum: uint16(o.Panorama_frameNum),
	}
	for i := 0; i < 14; i++ {
		out.DecoderTags[i] = uint32(C.go_libraw_olympus_decoder_tag(o, C.int(i)))
	}
	for i := 0; i < 2; i++ {
		out.SensorCalibration[i] = int(o.SensorCalibration[i])
		out.FocusMode[i] = uint16(o.FocusMode[i])
		out.StackedImage[i] = uint32(o.StackedImage[i])
	}
	for i := 0; i < 5; i++ {
		out.DriveMode[i] = uint16(o.DriveMode[i])
		out.AFPointSelected[i] = float64(o.AFPointSelected[i])
	}
	for i := 0; i < 64; i++ {
		out.AFAreas[i] = uint32(o.AFAreas[i])
	}
	for i := 0; i < 3; i++ {
		out.AFFineTuneAdj[i] = int16(o.AFFineTuneAdj[i])
		out.SpecialMode[i] = uint32(o.SpecialMode[i])
	}
	for i := 0; i < 4; i++ {
		out.AspectFrame[i] = uint16(o.AspectFrame[i])
	}
	return out
}

func convertSonyMakerNotes(s *C.libraw_sony_info_t) SonyMakerNotes {
	out := SonyMakerNotes{
		CameraType: uint16(s.CameraType), Sony0x9400Version: uint8(s.Sony0x9400_version),
		Sony0x9400ReleaseMode2: uint8(s.Sony0x9400_ReleaseMode2), Sony0x9400SequenceImageNumber: uint32(s.Sony0x9400_SequenceImageNumber),
		Sony0x9400SequenceLength1: uint8(s.Sony0x9400_SequenceLength1), Sony0x9400SequenceFileNumber: uint32(s.Sony0x9400_SequenceFileNumber),
		Sony0x9400SequenceLength2: uint8(s.Sony0x9400_SequenceLength2), AFAreaModeSetting: uint8(s.AFAreaModeSetting),
		AFAreaMode: uint16(s.AFAreaMode), AFPointSelected: uint8(s.AFPointSelected), AFPointSelected0x201e: uint8(s.AFPointSelected_0x201e),
		NAFPointsUsed: int16(s.nAFPointsUsed), AFTracking: uint8(s.AFTracking), AFType: uint8(s.AFType), FocusPosition: uint16(s.FocusPosition),
		AFMicroAdjValue: int8(s.AFMicroAdjValue), AFMicroAdjOn: int8(s.AFMicroAdjOn), AFMicroAdjRegisteredLenses: uint8(s.AFMicroAdjRegisteredLenses),
		VariableLowPassFilter: uint16(s.VariableLowPassFilter), LongExposureNoiseReduction: uint32(s.LongExposureNoiseReduction),
		HighISONoiseReduction: uint16(s.HighISONoiseReduction), Group2010: uint16(s.group2010), Group9050: uint16(s.group9050), LenGroup9050: uint16(C.go_libraw_sony_len_group9050(s)),
		RealISOOffset: uint16(s.real_iso_offset), MeteringModeOffset: uint16(s.MeteringMode_offset), ExposureProgramOffset: uint16(s.ExposureProgram_offset),
		ReleaseMode2Offset: uint16(s.ReleaseMode2_offset), MinoltaCamID: uint32(s.MinoltaCamID), Firmware: float32(s.firmware),
		ImageCount3Offset: uint16(s.ImageCount3_offset), ImageCount3: uint32(s.ImageCount3), ElectronicFrontCurtainShutter: uint32(s.ElectronicFrontCurtainShutter),
		MeteringMode2: uint16(s.MeteringMode2), SonyDateTime: cString(&s.SonyDateTime[0]), ShotNumberSincePowerUp: uint32(s.ShotNumberSincePowerUp),
		PixelShiftGroupPrefix: uint16(s.PixelShiftGroupPrefix), PixelShiftGroupID: uint32(s.PixelShiftGroupID), NShotsInPixelShiftGroup: int8(s.nShotsInPixelShiftGroup),
		NumInPixelShiftGroup: int8(s.numInPixelShiftGroup), PRDImageHeight: uint16(s.prd_ImageHeight), PRDImageWidth: uint16(s.prd_ImageWidth),
		PRDTotalBPS: uint16(s.prd_Total_bps), PRDActiveBPS: uint16(s.prd_Active_bps), PRDStorageMethod: uint16(s.prd_StorageMethod), PRDBayerPattern: uint16(s.prd_BayerPattern),
		SonyRawFileType: uint16(s.SonyRawFileType), RawFileType: uint16(s.RAWFileType), RawSizeType: uint16(s.RawSizeType), Quality: uint32(s.Quality),
		FileFormat: uint16(s.FileFormat), MetaVersion: cString(&s.MetaVersion[0]), AspectRatio: float32(C.go_libraw_sony_aspect_ratio(s)),
	}
	for i := 0; i < 2; i++ {
		out.FlexibleSpotPosition[i] = uint16(s.FlexibleSpotPosition[i])
		out.HDR[i] = uint16(s.HDR[i])
	}
	for i := 0; i < 10; i++ {
		out.AFPointsUsed[i] = uint8(s.AFPointsUsed[i])
	}
	for i := 0; i < 4; i++ {
		out.FocusLocation[i] = uint16(s.FocusLocation[i])
	}
	return out
}

func convertKodakMakerNotes(k *C.libraw_kodak_makernotes_t) KodakMakerNotes {
	out := KodakMakerNotes{BlackLevelTop: uint16(k.BlackLevelTop), BlackLevelBottom: uint16(k.BlackLevelBottom), OffsetLeft: int16(k.offset_left), OffsetTop: int16(k.offset_top), ClipBlack: uint16(k.clipBlack), ClipWhite: uint16(k.clipWhite), Val018Percent: uint16(k.val018percent), Val100Percent: uint16(k.val100percent), Val170Percent: uint16(k.val170percent), MakerNoteKodak8a: int16(k.MakerNoteKodak8a), ISOCalibrationGain: float32(k.ISOCalibrationGain), AnalogISO: float32(k.AnalogISO)}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out.ROMMCamDaylight[i][j] = float32(k.romm_camDaylight[i][j])
			out.ROMMCamTungsten[i][j] = float32(k.romm_camTungsten[i][j])
			out.ROMMCamFluorescent[i][j] = float32(k.romm_camFluorescent[i][j])
			out.ROMMCamFlash[i][j] = float32(k.romm_camFlash[i][j])
			out.ROMMCamCustom[i][j] = float32(k.romm_camCustom[i][j])
			out.ROMMCamAuto[i][j] = float32(k.romm_camAuto[i][j])
		}
	}
	return out
}

func convertPanasonicMakerNotes(p *C.libraw_panasonic_makernotes_t) PanasonicMakerNotes {
	out := PanasonicMakerNotes{Compression: uint16(p.Compression), BlackLevelDim: uint16(p.BlackLevelDim), Multishot: uint32(p.Multishot), Gamma: float32(p.gamma), FocusStepNear: int16(p.FocusStepNear), FocusStepCount: int16(p.FocusStepCount), ZoomPosition: uint32(p.ZoomPosition), LensManufacturer: uint32(p.LensManufacturer)}
	for i := 0; i < 8; i++ {
		out.BlackLevel[i] = float32(p.BlackLevel[i])
	}
	for i := 0; i < 3; i++ {
		out.HighISOMultiplier[i] = int(p.HighISOMultiplier[i])
	}
	return out
}

func convertPentaxMakerNotes(p *C.libraw_pentax_makernotes_t) PentaxMakerNotes {
	out := PentaxMakerNotes{AFPointSelectedArea: uint16(p.AFPointSelected_Area), AFPointsInFocusVersion: int(p.AFPointsInFocus_version), AFPointsInFocus: uint32(p.AFPointsInFocus), FocusPosition: uint16(p.FocusPosition), AFAdjustment: int16(p.AFAdjustment), AFPointMode: uint8(p.AFPointMode), MultiExposure: uint8(p.MultiExposure), Quality: uint16(p.Quality)}
	for i := 0; i < 4; i++ {
		out.DriveMode[i] = uint8(p.DriveMode[i])
		out.DynamicRangeExpansion[i] = uint8(C.go_libraw_pentax_dynamic_range_expansion(p, C.int(i)))
	}
	for i := 0; i < 2; i++ {
		out.FocusMode[i] = uint16(p.FocusMode[i])
		out.AFPointSelected[i] = uint16(p.AFPointSelected[i])
	}
	return out
}

func convertPhaseOneMakerNotes(p *C.libraw_p1_makernotes_t) PhaseOneMakerNotes {
	return PhaseOneMakerNotes{Software: cString(&p.Software[0]), SystemType: cString(&p.SystemType[0]), FirmwareString: cString(&p.FirmwareString[0]), SystemModel: cString(&p.SystemModel[0])}
}

func convertRicohMakerNotes(r *C.libraw_ricoh_makernotes_t) RicohMakerNotes {
	out := RicohMakerNotes{AFStatus: uint16(r.AFStatus), AFAreaMode: uint16(r.AFAreaMode), SensorWidth: uint32(r.SensorWidth), SensorHeight: uint32(r.SensorHeight), CroppedImageWidth: uint32(r.CroppedImageWidth), CroppedImageHeight: uint32(r.CroppedImageHeight), WideAdapter: uint16(r.WideAdapter), CropMode: uint16(r.CropMode), NDFilter: uint16(r.NDFilter), AutoBracketing: uint16(r.AutoBracketing), MacroMode: uint16(r.MacroMode), FlashMode: uint16(r.FlashMode), FlashExposureComp: float64(r.FlashExposureComp), ManualFlashOutput: float64(r.ManualFlashOutput)}
	for i := 0; i < 2; i++ {
		out.AFAreaXPosition[i] = uint32(r.AFAreaXPosition[i])
		out.AFAreaYPosition[i] = uint32(r.AFAreaYPosition[i])
	}
	return out
}

func convertSamsungMakerNotes(s *C.libraw_samsung_makernotes_t) SamsungMakerNotes {
	out := SamsungMakerNotes{DigitalGain: float64(s.DigitalGain), DeviceType: int(s.DeviceType), LensFirmware: cString(&s.LensFirmware[0])}
	for i := 0; i < 4; i++ {
		out.ImageSizeFull[i] = uint32(s.ImageSizeFull[i])
		out.ImageSizeCrop[i] = uint32(s.ImageSizeCrop[i])
	}
	for i := 0; i < 2; i++ {
		out.ColorSpace[i] = int(s.ColorSpace[i])
	}
	for i := 0; i < 11; i++ {
		out.Key[i] = uint32(s.key[i])
	}
	return out
}

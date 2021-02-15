package webcam

// #include "v4l2-binding.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

var formatToString = map[uint32]string{

	C.V4L2_PIX_FMT_RGB332:   "V4L2_PIX_FMT_RGB332",
	C.V4L2_PIX_FMT_RGB444:   "V4L2_PIX_FMT_RGB444",
	C.V4L2_PIX_FMT_ARGB444:  "V4L2_PIX_FMT_ARGB444",
	C.V4L2_PIX_FMT_XRGB444:  "V4L2_PIX_FMT_XRGB444",
	C.V4L2_PIX_FMT_RGB555:   "V4L2_PIX_FMT_RGB555",
	C.V4L2_PIX_FMT_ARGB555:  "V4L2_PIX_FMT_ARGB555",
	C.V4L2_PIX_FMT_XRGB555:  "V4L2_PIX_FMT_XRGB555",
	C.V4L2_PIX_FMT_RGB565:   "V4L2_PIX_FMT_RGB565",
	C.V4L2_PIX_FMT_RGB555X:  "V4L2_PIX_FMT_RGB555X",
	C.V4L2_PIX_FMT_ARGB555X: "V4L2_PIX_FMT_ARGB555X",
	C.V4L2_PIX_FMT_XRGB555X: "V4L2_PIX_FMT_XRGB555X",
	C.V4L2_PIX_FMT_RGB565X:  "V4L2_PIX_FMT_RGB565X",
	C.V4L2_PIX_FMT_BGR666:   "V4L2_PIX_FMT_BGR666",
	C.V4L2_PIX_FMT_BGR24:    "V4L2_PIX_FMT_BGR24",
	C.V4L2_PIX_FMT_RGB24:    "V4L2_PIX_FMT_RGB24",
	C.V4L2_PIX_FMT_BGR32:    "V4L2_PIX_FMT_BGR32",
	C.V4L2_PIX_FMT_ABGR32:   "V4L2_PIX_FMT_ABGR32",
	C.V4L2_PIX_FMT_XBGR32:   "V4L2_PIX_FMT_XBGR32",
	C.V4L2_PIX_FMT_RGB32:    "V4L2_PIX_FMT_RGB32",
	C.V4L2_PIX_FMT_ARGB32:   "V4L2_PIX_FMT_ARGB32",
	C.V4L2_PIX_FMT_XRGB32:   "V4L2_PIX_FMT_XRGB32",
	C.V4L2_PIX_FMT_GREY:     "V4L2_PIX_FMT_GREY",
	C.V4L2_PIX_FMT_Y4:       "V4L2_PIX_FMT_Y4",
	C.V4L2_PIX_FMT_Y6:       "V4L2_PIX_FMT_Y6",
	C.V4L2_PIX_FMT_Y10:      "V4L2_PIX_FMT_Y10",
	C.V4L2_PIX_FMT_Y12:      "V4L2_PIX_FMT_Y12",
	C.V4L2_PIX_FMT_Y16:      "V4L2_PIX_FMT_Y16",
	C.V4L2_PIX_FMT_Y16_BE:   "V4L2_PIX_FMT_Y16_BE",

	C.V4L2_PIX_FMT_Y10BPACK: "V4L2_PIX_FMT_Y10BPACK",

	C.V4L2_PIX_FMT_PAL8: "V4L2_PIX_FMT_PAL8",

	C.V4L2_PIX_FMT_UV8: "V4L2_PIX_FMT_UV8",

	C.V4L2_PIX_FMT_YUYV:   "V4L2_PIX_FMT_YUYV",
	C.V4L2_PIX_FMT_YYUV:   "V4L2_PIX_FMT_YYUV",
	C.V4L2_PIX_FMT_YVYU:   "V4L2_PIX_FMT_YVYU",
	C.V4L2_PIX_FMT_UYVY:   "V4L2_PIX_FMT_UYVY",
	C.V4L2_PIX_FMT_VYUY:   "V4L2_PIX_FMT_VYUY",
	C.V4L2_PIX_FMT_Y41P:   "V4L2_PIX_FMT_Y41P",
	C.V4L2_PIX_FMT_YUV444: "V4L2_PIX_FMT_YUV444",
	C.V4L2_PIX_FMT_YUV555: "V4L2_PIX_FMT_YUV555",
	C.V4L2_PIX_FMT_YUV565: "V4L2_PIX_FMT_YUV565",
	C.V4L2_PIX_FMT_YUV32:  "V4L2_PIX_FMT_YUV32",
	C.V4L2_PIX_FMT_HI240:  "V4L2_PIX_FMT_HI240",
	C.V4L2_PIX_FMT_HM12:   "V4L2_PIX_FMT_HM12",
	C.V4L2_PIX_FMT_M420:   "V4L2_PIX_FMT_M420",

	C.V4L2_PIX_FMT_NV12: "V4L2_PIX_FMT_NV12",
	C.V4L2_PIX_FMT_NV21: "V4L2_PIX_FMT_NV21",
	C.V4L2_PIX_FMT_NV16: "V4L2_PIX_FMT_NV16",
	C.V4L2_PIX_FMT_NV61: "V4L2_PIX_FMT_NV61",
	C.V4L2_PIX_FMT_NV24: "V4L2_PIX_FMT_NV24",
	C.V4L2_PIX_FMT_NV42: "V4L2_PIX_FMT_NV42",

	C.V4L2_PIX_FMT_NV12M:        "V4L2_PIX_FMT_NV12M",
	C.V4L2_PIX_FMT_NV21M:        "V4L2_PIX_FMT_NV21M",
	C.V4L2_PIX_FMT_NV16M:        "V4L2_PIX_FMT_NV16M",
	C.V4L2_PIX_FMT_NV61M:        "V4L2_PIX_FMT_NV61M",
	C.V4L2_PIX_FMT_NV12MT:       "V4L2_PIX_FMT_NV12MT",
	C.V4L2_PIX_FMT_NV12MT_16X16: "V4L2_PIX_FMT_NV12MT_16X16",

	C.V4L2_PIX_FMT_YUV410:  "V4L2_PIX_FMT_YUV410",
	C.V4L2_PIX_FMT_YVU410:  "V4L2_PIX_FMT_YVU410",
	C.V4L2_PIX_FMT_YUV411P: "V4L2_PIX_FMT_YUV411P",
	C.V4L2_PIX_FMT_YUV420:  "V4L2_PIX_FMT_YUV420",
	C.V4L2_PIX_FMT_YVU420:  "V4L2_PIX_FMT_YVU420",
	C.V4L2_PIX_FMT_YUV422P: "V4L2_PIX_FMT_YUV422P",

	C.V4L2_PIX_FMT_YUV420M: "V4L2_PIX_FMT_YUV420M",
	C.V4L2_PIX_FMT_YVU420M: "V4L2_PIX_FMT_YVU420M",
	C.V4L2_PIX_FMT_YUV422M: "V4L2_PIX_FMT_YUV422M",
	C.V4L2_PIX_FMT_YVU422M: "V4L2_PIX_FMT_YVU422M",
	C.V4L2_PIX_FMT_YUV444M: "V4L2_PIX_FMT_YUV444M",
	C.V4L2_PIX_FMT_YVU444M: "V4L2_PIX_FMT_YVU444M",

	C.V4L2_PIX_FMT_SBGGR8:   "V4L2_PIX_FMT_SBGGR8",
	C.V4L2_PIX_FMT_SGBRG8:   "V4L2_PIX_FMT_SGBRG8",
	C.V4L2_PIX_FMT_SGRBG8:   "V4L2_PIX_FMT_SGRBG8",
	C.V4L2_PIX_FMT_SRGGB8:   "V4L2_PIX_FMT_SRGGB8",
	C.V4L2_PIX_FMT_SBGGR10:  "V4L2_PIX_FMT_SBGGR10",
	C.V4L2_PIX_FMT_SGBRG10:  "V4L2_PIX_FMT_SGBRG10",
	C.V4L2_PIX_FMT_SGRBG10:  "V4L2_PIX_FMT_SGRBG10",
	C.V4L2_PIX_FMT_SRGGB10:  "V4L2_PIX_FMT_SRGGB10",
	C.V4L2_PIX_FMT_SBGGR10P: "V4L2_PIX_FMT_SBGGR10P",
	C.V4L2_PIX_FMT_SGBRG10P: "V4L2_PIX_FMT_SGBRG10P",
	C.V4L2_PIX_FMT_SGRBG10P: "V4L2_PIX_FMT_SGRBG10P",
	C.V4L2_PIX_FMT_SRGGB10P: "V4L2_PIX_FMT_SRGGB10P",

	C.V4L2_PIX_FMT_SBGGR10ALAW8: "V4L2_PIX_FMT_SBGGR10ALAW8",
	C.V4L2_PIX_FMT_SGBRG10ALAW8: "V4L2_PIX_FMT_SGBRG10ALAW8",
	C.V4L2_PIX_FMT_SGRBG10ALAW8: "V4L2_PIX_FMT_SGRBG10ALAW8",
	C.V4L2_PIX_FMT_SRGGB10ALAW8: "V4L2_PIX_FMT_SRGGB10ALAW8",

	C.V4L2_PIX_FMT_SBGGR10DPCM8: "V4L2_PIX_FMT_SBGGR10DPCM8",
	C.V4L2_PIX_FMT_SGBRG10DPCM8: "V4L2_PIX_FMT_SGBRG10DPCM8",
	C.V4L2_PIX_FMT_SGRBG10DPCM8: "V4L2_PIX_FMT_SGRBG10DPCM8",
	C.V4L2_PIX_FMT_SRGGB10DPCM8: "V4L2_PIX_FMT_SRGGB10DPCM8",
	C.V4L2_PIX_FMT_SBGGR12:      "V4L2_PIX_FMT_SBGGR12",
	C.V4L2_PIX_FMT_SGBRG12:      "V4L2_PIX_FMT_SGBRG12",
	C.V4L2_PIX_FMT_SGRBG12:      "V4L2_PIX_FMT_SGRBG12",
	C.V4L2_PIX_FMT_SRGGB12:      "V4L2_PIX_FMT_SRGGB12",
	C.V4L2_PIX_FMT_SBGGR16:      "V4L2_PIX_FMT_SBGGR16",

	C.V4L2_PIX_FMT_MJPEG:       "V4L2_PIX_FMT_MJPEG",
	C.V4L2_PIX_FMT_JPEG:        "V4L2_PIX_FMT_JPEG",
	C.V4L2_PIX_FMT_DV:          "V4L2_PIX_FMT_DV",
	C.V4L2_PIX_FMT_MPEG:        "V4L2_PIX_FMT_MPEG",
	C.V4L2_PIX_FMT_H264:        "V4L2_PIX_FMT_H264",
	C.V4L2_PIX_FMT_H264_NO_SC:  "V4L2_PIX_FMT_H264_NO_SC",
	C.V4L2_PIX_FMT_H264_MVC:    "V4L2_PIX_FMT_H264_MVC",
	C.V4L2_PIX_FMT_H263:        "V4L2_PIX_FMT_H263",
	C.V4L2_PIX_FMT_MPEG1:       "V4L2_PIX_FMT_MPEG1",
	C.V4L2_PIX_FMT_MPEG2:       "V4L2_PIX_FMT_MPEG2",
	C.V4L2_PIX_FMT_MPEG4:       "V4L2_PIX_FMT_MPEG4",
	C.V4L2_PIX_FMT_XVID:        "V4L2_PIX_FMT_XVID",
	C.V4L2_PIX_FMT_VC1_ANNEX_G: "V4L2_PIX_FMT_VC1_ANNEX_G",
	C.V4L2_PIX_FMT_VC1_ANNEX_L: "V4L2_PIX_FMT_VC1_ANNEX_L",
	C.V4L2_PIX_FMT_VP8:         "V4L2_PIX_FMT_VP8",

	C.V4L2_PIX_FMT_CPIA1:        "V4L2_PIX_FMT_CPIA1",
	C.V4L2_PIX_FMT_WNVA:         "V4L2_PIX_FMT_WNVA",
	C.V4L2_PIX_FMT_SN9C10X:      "V4L2_PIX_FMT_SN9C10X",
	C.V4L2_PIX_FMT_SN9C20X_I420: "V4L2_PIX_FMT_SN9C20X_I420",
	C.V4L2_PIX_FMT_PWC1:         "V4L2_PIX_FMT_PWC1",
	C.V4L2_PIX_FMT_PWC2:         "V4L2_PIX_FMT_PWC2",
	C.V4L2_PIX_FMT_ET61X251:     "V4L2_PIX_FMT_ET61X251",
	C.V4L2_PIX_FMT_SPCA501:      "V4L2_PIX_FMT_SPCA501",
	C.V4L2_PIX_FMT_SPCA505:      "V4L2_PIX_FMT_SPCA505",
	C.V4L2_PIX_FMT_SPCA508:      "V4L2_PIX_FMT_SPCA508",
	C.V4L2_PIX_FMT_SPCA561:      "V4L2_PIX_FMT_SPCA561",
	C.V4L2_PIX_FMT_PAC207:       "V4L2_PIX_FMT_PAC207",
	C.V4L2_PIX_FMT_MR97310A:     "V4L2_PIX_FMT_MR97310A",
	C.V4L2_PIX_FMT_JL2005BCD:    "V4L2_PIX_FMT_JL2005BCD",
	C.V4L2_PIX_FMT_SN9C2028:     "V4L2_PIX_FMT_SN9C2028",
	C.V4L2_PIX_FMT_SQ905C:       "V4L2_PIX_FMT_SQ905C",
	C.V4L2_PIX_FMT_PJPG:         "V4L2_PIX_FMT_PJPG",
	C.V4L2_PIX_FMT_OV511:        "V4L2_PIX_FMT_OV511",
	C.V4L2_PIX_FMT_OV518:        "V4L2_PIX_FMT_OV518",
	C.V4L2_PIX_FMT_STV0680:      "V4L2_PIX_FMT_STV0680",
	C.V4L2_PIX_FMT_TM6000:       "V4L2_PIX_FMT_TM6000",
	C.V4L2_PIX_FMT_CIT_YYVYUY:   "V4L2_PIX_FMT_CIT_YYVYUY",
	C.V4L2_PIX_FMT_KONICA420:    "V4L2_PIX_FMT_KONICA420",
	C.V4L2_PIX_FMT_JPGL:         "V4L2_PIX_FMT_JPGL",
	C.V4L2_PIX_FMT_SE401:        "V4L2_PIX_FMT_SE401",
	C.V4L2_PIX_FMT_S5C_UYVY_JPG: "V4L2_PIX_FMT_S5C_UYVY_JPG",
	C.V4L2_PIX_FMT_Y8I:          "V4L2_PIX_FMT_Y8I",
	C.V4L2_PIX_FMT_Y12I:         "V4L2_PIX_FMT_Y12I",
	C.V4L2_PIX_FMT_Z16:          "V4L2_PIX_FMT_Z16",
}

//-------------------------------------------------------------------------------------------------
//PIXEL FORMAT INTERFACE IMPL
//-------------------------------------------------------------------------------------------------

type pixelFormat struct {
	name  string
	desc  string
	value uint32
}

func (p pixelFormat) Name() string {
	return p.name
}

func (p pixelFormat) Description() string {
	return p.desc
}

func (p pixelFormat) String() string {
	return fmt.Sprintf("PixelFormat[%s - %s, %d]", p.name, p.desc, p.value)
}

//-------------------------------------------------------------------------------------------------
//QUERY FORMAT
//-------------------------------------------------------------------------------------------------

func (d *device) QueryFormats() ([]PixelFormat, error) {

	result := []PixelFormat{}

	desc, err := C.queryFormats(C.int(d.file.Fd()), nil)

	if err != nil {
		return nil, err
	}

	formatCode := binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&desc.pixelformat), 4))
	name := formatToString[formatCode]
	description := string(C.GoBytes(unsafe.Pointer(&desc.description), 32))

	result = append(result, pixelFormat{name: name, desc: description, value: formatCode})

	for {
		desc, err = C.queryFormats(C.int(d.file.Fd()), desc)

		if desc == nil {
			break
		}

		if err != nil {
			return nil, err
		}

		formatCode = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&desc.pixelformat), 4))
		name = formatToString[formatCode]
		description = string(C.GoBytes(unsafe.Pointer(&desc.description), 32))

		result = append(result, pixelFormat{name: name, desc: description, value: formatCode})
	}

	return result, nil
}
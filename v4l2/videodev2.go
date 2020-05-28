package v4l2

import (
	"encoding/binary"
	"unsafe"
)

/*
THIS FILE REPRESENTS PORT OF videodev2.h HEADER FILE
FOR C LANGUAGE.
*/

type V4l2Capability struct {
	Driver       [16]uint8
	Card         [32]uint8
	BusInfo      [32]uint8
	Version      uint32
	Capabilities uint32
	DeviceCaps   uint32
	Reserved     [3]uint32
}

/* Values for 'capabilities' field */
const (
	V4L2_CAP_VIDEO_CAPTURE        = 0x00000001 /* Is a video capture device */
	V4L2_CAP_VIDEO_OUTPUT         = 0x00000002 /* Is a video output device */
	V4L2_CAP_VIDEO_OVERLAY        = 0x00000004 /* Can do video overlay */
	V4L2_CAP_VBI_CAPTURE          = 0x00000010 /* Is a raw VBI capture device */
	V4L2_CAP_VBI_OUTPUT           = 0x00000020 /* Is a raw VBI output device */
	V4L2_CAP_SLICED_VBI_CAPTURE   = 0x00000040 /* Is a sliced VBI capture device */
	V4L2_CAP_SLICED_VBI_OUTPUT    = 0x00000080 /* Is a sliced VBI output device */
	V4L2_CAP_RDS_CAPTURE          = 0x00000100 /* RDS data capture */
	V4L2_CAP_VIDEO_OUTPUT_OVERLAY = 0x00000200 /* Can do video output overlay */
	V4L2_CAP_HW_FREQ_SEEK         = 0x00000400 /* Can do hardware frequency seek  */
	V4L2_CAP_RDS_OUTPUT           = 0x00000800 /* Is an RDS encoder */

	/* Is a video capture device that supports multiplanar formats */
	V4L2_CAP_VIDEO_CAPTURE_MPLANE = 0x00001000
	/* Is a video output device that supports multiplanar formats */
	V4L2_CAP_VIDEO_OUTPUT_MPLANE = 0x00002000
	/* Is a video mem-to-mem device that supports multiplanar formats */
	V4L2_CAP_VIDEO_M2M_MPLANE = 0x00004000
	/* Is a video mem-to-mem device */
	V4L2_CAP_VIDEO_M2M = 0x00008000

	V4L2_CAP_TUNER     = 0x00010000 /* has a tuner */
	V4L2_CAP_AUDIO     = 0x00020000 /* has audio support */
	V4L2_CAP_RADIO     = 0x00040000 /* is a radio device */
	V4L2_CAP_MODULATOR = 0x00080000 /* has a modulator */

	V4L2_CAP_SDR_CAPTURE    = 0x00100000 /* Is a SDR capture device */
	V4L2_CAP_EXT_PIX_FORMAT = 0x00200000 /* Supports the extended pixel format */
	V4L2_CAP_SDR_OUTPUT     = 0x00400000 /* Is a SDR output device */

	V4L2_CAP_READWRITE = 0x01000000 /* read/write systemcalls */
	V4L2_CAP_ASYNCIO   = 0x02000000 /* async I/O */
	V4L2_CAP_STREAMING = 0x04000000 /* streaming I/O ioctls */

	V4L2_CAP_TOUCH = 0x10000000 /* Is a touch device */

	V4L2_CAP_DEVICE_CAPS = 0x80000000 /* sets device capabilities field */
)

/*
 *	F O R M A T   E N U M E R A T I O N
 */
type V4l2Fmtdesc struct {
	Index       uint32 /* Format number      */
	Typ         uint32 /* enum v4l2_buf_type */
	Flags       uint32
	Description [32]uint8 /* Description string */
	Pixelformat uint32    /* Format fourcc      */
	Reserved    [4]uint32
}

const (
	V4L2_BUF_TYPE_VIDEO_CAPTURE        = 1
	V4L2_BUF_TYPE_VIDEO_OUTPUT         = 2
	V4L2_BUF_TYPE_VIDEO_OVERLAY        = 3
	V4L2_BUF_TYPE_VBI_CAPTURE          = 4
	V4L2_BUF_TYPE_VBI_OUTPUT           = 5
	V4L2_BUF_TYPE_SLICED_VBI_CAPTURE   = 6
	V4L2_BUF_TYPE_SLICED_VBI_OUTPUT    = 7
	V4L2_BUF_TYPE_VIDEO_OUTPUT_OVERLAY = 8
	V4L2_BUF_TYPE_VIDEO_CAPTURE_MPLANE = 9
	V4L2_BUF_TYPE_VIDEO_OUTPUT_MPLANE  = 10
	V4L2_BUF_TYPE_SDR_CAPTURE          = 11
	V4L2_BUF_TYPE_SDR_OUTPUT           = 12
	/* Deprecated, do not use */
	V4L2_BUF_TYPE_PRIVATE = 0x80
)

func v4l2_fourcc(a uint32, b uint32, c uint32, d uint32) uint32 {
	return (a | (b << 8) | (c << 16) | (d << 24))
}

func v4l2_fourcc_be(a uint32, b uint32, c uint32, d uint32) uint32 {
	return (v4l2_fourcc(a, b, c, d) | (1 << 31))
}

/*
 *	E N U M S
 */
const (
	V4L2_FIELD_ANY = 0 /* driver can choose from none,
	top, bottom, interlaced
	depending on whatever it thinks
	is approximate ... */
	V4L2_FIELD_NONE       = 1 /* this device has no fields ... */
	V4L2_FIELD_TOP        = 2 /* top field only */
	V4L2_FIELD_BOTTOM     = 3 /* bottom field only */
	V4L2_FIELD_INTERLACED = 4 /* both fields interlaced */
	V4L2_FIELD_SEQ_TB     = 5 /* both fields sequential into one
	buffer, top-bottom order */
	V4L2_FIELD_SEQ_BT    = 6 /* same as above + bottom-top order */
	V4L2_FIELD_ALTERNATE = 7 /* both fields alternating into
	separate buffers */
	V4L2_FIELD_INTERLACED_TB = 8 /* both fields interlaced, top field
	first and the top field is
	transmitted first */
	V4L2_FIELD_INTERLACED_BT = 9 /* both fields interlaced, top field
	first and the bottom field is
	transmitted first */
)

/*      Pixel format         FOURCC                          depth  Description  */

/* RGB formats */
var V4L2_PIX_FMT_RGB332 uint32 = v4l2_fourcc('R', 'G', 'B', '1')      /*  8  RGB-3-3-2     */
var V4L2_PIX_FMT_RGB444 uint32 = v4l2_fourcc('R', '4', '4', '4')      /* 16  xxxxrrrr ggggbbbb */
var V4L2_PIX_FMT_ARGB444 uint32 = v4l2_fourcc('A', 'R', '1', '2')     /* 16  aaaarrrr ggggbbbb */
var V4L2_PIX_FMT_XRGB444 uint32 = v4l2_fourcc('X', 'R', '1', '2')     /* 16  xxxxrrrr ggggbbbb */
var V4L2_PIX_FMT_RGB555 uint32 = v4l2_fourcc('R', 'G', 'B', 'O')      /* 16  RGB-5-5-5     */
var V4L2_PIX_FMT_ARGB555 uint32 = v4l2_fourcc('A', 'R', '1', '5')     /* 16  ARGB-1-5-5-5  */
var V4L2_PIX_FMT_XRGB555 uint32 = v4l2_fourcc('X', 'R', '1', '5')     /* 16  XRGB-1-5-5-5  */
var V4L2_PIX_FMT_RGB565 uint32 = v4l2_fourcc('R', 'G', 'B', 'P')      /* 16  RGB-5-6-5     */
var V4L2_PIX_FMT_RGB555X uint32 = v4l2_fourcc('R', 'G', 'B', 'Q')     /* 16  RGB-5-5-5 BE  */
var V4L2_PIX_FMT_ARGB555X uint32 = v4l2_fourcc_be('A', 'R', '1', '5') /* 16  ARGB-5-5-5 BE */
var V4L2_PIX_FMT_XRGB555X uint32 = v4l2_fourcc_be('X', 'R', '1', '5') /* 16  XRGB-5-5-5 BE */
var V4L2_PIX_FMT_RGB565X uint32 = v4l2_fourcc('R', 'G', 'B', 'R')     /* 16  RGB-5-6-5 BE  */
var V4L2_PIX_FMT_BGR666 uint32 = v4l2_fourcc('B', 'G', 'R', 'H')      /* 18  BGR-6-6-6	  */
var V4L2_PIX_FMT_BGR24 uint32 = v4l2_fourcc('B', 'G', 'R', '3')       /* 24  BGR-8-8-8     */
var V4L2_PIX_FMT_RGB24 uint32 = v4l2_fourcc('R', 'G', 'B', '3')       /* 24  RGB-8-8-8     */
var V4L2_PIX_FMT_BGR32 uint32 = v4l2_fourcc('B', 'G', 'R', '4')       /* 32  BGR-8-8-8-8   */
var V4L2_PIX_FMT_ABGR32 uint32 = v4l2_fourcc('A', 'R', '2', '4')      /* 32  BGRA-8-8-8-8  */
var V4L2_PIX_FMT_XBGR32 uint32 = v4l2_fourcc('X', 'R', '2', '4')      /* 32  BGRX-8-8-8-8  */
var V4L2_PIX_FMT_RGB32 uint32 = v4l2_fourcc('R', 'G', 'B', '4')       /* 32  RGB-8-8-8-8   */
var V4L2_PIX_FMT_ARGB32 uint32 = v4l2_fourcc('B', 'A', '2', '4')      /* 32  ARGB-8-8-8-8  */
var V4L2_PIX_FMT_XRGB32 uint32 = v4l2_fourcc('B', 'X', '2', '4')      /* 32  XRGB-8-8-8-8  */

/* Grey formats */
var V4L2_PIX_FMT_GREY uint32 = v4l2_fourcc('G', 'R', 'E', 'Y')      /*  8  Greyscale     */
var V4L2_PIX_FMT_Y4 uint32 = v4l2_fourcc('Y', '0', '4', ' ')        /*  4  Greyscale     */
var V4L2_PIX_FMT_Y6 uint32 = v4l2_fourcc('Y', '0', '6', ' ')        /*  6  Greyscale     */
var V4L2_PIX_FMT_Y10 uint32 = v4l2_fourcc('Y', '1', '0', ' ')       /* 10  Greyscale     */
var V4L2_PIX_FMT_Y12 uint32 = v4l2_fourcc('Y', '1', '2', ' ')       /* 12  Greyscale     */
var V4L2_PIX_FMT_Y16 uint32 = v4l2_fourcc('Y', '1', '6', ' ')       /* 16  Greyscale     */
var V4L2_PIX_FMT_Y16_BE uint32 = v4l2_fourcc_be('Y', '1', '6', ' ') /* 16  Greyscale BE  */

/* Grey bit-packed formats */
var V4L2_PIX_FMT_Y10BPACK uint32 = v4l2_fourcc('Y', '1', '0', 'B') /* 10  Greyscale bit-packed */

/* Palette formats */
var V4L2_PIX_FMT_PAL8 uint32 = v4l2_fourcc('P', 'A', 'L', '8') /*  8  8-bit palette */

/* Chrominance formats */
var V4L2_PIX_FMT_UV8 uint32 = v4l2_fourcc('U', 'V', '8', ' ') /*  8  UV 4:4 */

/* Luminance+Chrominance formats */
var V4L2_PIX_FMT_YUYV uint32 = v4l2_fourcc('Y', 'U', 'Y', 'V')   /* 16  YUV 4:2:2     */
var V4L2_PIX_FMT_YYUV uint32 = v4l2_fourcc('Y', 'Y', 'U', 'V')   /* 16  YUV 4:2:2     */
var V4L2_PIX_FMT_YVYU uint32 = v4l2_fourcc('Y', 'V', 'Y', 'U')   /* 16 YVU 4:2:2 */
var V4L2_PIX_FMT_UYVY uint32 = v4l2_fourcc('U', 'Y', 'V', 'Y')   /* 16  YUV 4:2:2     */
var V4L2_PIX_FMT_VYUY uint32 = v4l2_fourcc('V', 'Y', 'U', 'Y')   /* 16  YUV 4:2:2     */
var V4L2_PIX_FMT_Y41P uint32 = v4l2_fourcc('Y', '4', '1', 'P')   /* 12  YUV 4:1:1     */
var V4L2_PIX_FMT_YUV444 uint32 = v4l2_fourcc('Y', '4', '4', '4') /* 16  xxxxyyyy uuuuvvvv */
var V4L2_PIX_FMT_YUV555 uint32 = v4l2_fourcc('Y', 'U', 'V', 'O') /* 16  YUV-5-5-5     */
var V4L2_PIX_FMT_YUV565 uint32 = v4l2_fourcc('Y', 'U', 'V', 'P') /* 16  YUV-5-6-5     */
var V4L2_PIX_FMT_YUV32 uint32 = v4l2_fourcc('Y', 'U', 'V', '4')  /* 32  YUV-8-8-8-8   */
var V4L2_PIX_FMT_HI240 uint32 = v4l2_fourcc('H', 'I', '2', '4')  /*  8  8-bit color   */
var V4L2_PIX_FMT_HM12 uint32 = v4l2_fourcc('H', 'M', '1', '2')   /*  8  YUV 4:2:0 16x16 macroblocks */
var V4L2_PIX_FMT_M420 uint32 = v4l2_fourcc('M', '4', '2', '0')   /* 12  YUV 4:2:0 2 lines y, 1 line uv interleaved */

/* two planes -- one Y, one Cr + Cb interleaved  */
var V4L2_PIX_FMT_NV12 uint32 = v4l2_fourcc('N', 'V', '1', '2') /* 12  Y/CbCr 4:2:0  */
var V4L2_PIX_FMT_NV21 uint32 = v4l2_fourcc('N', 'V', '2', '1') /* 12  Y/CrCb 4:2:0  */
var V4L2_PIX_FMT_NV16 uint32 = v4l2_fourcc('N', 'V', '1', '6') /* 16  Y/CbCr 4:2:2  */
var V4L2_PIX_FMT_NV61 uint32 = v4l2_fourcc('N', 'V', '6', '1') /* 16  Y/CrCb 4:2:2  */
var V4L2_PIX_FMT_NV24 uint32 = v4l2_fourcc('N', 'V', '2', '4') /* 24  Y/CbCr 4:4:4  */
var V4L2_PIX_FMT_NV42 uint32 = v4l2_fourcc('N', 'V', '4', '2') /* 24  Y/CrCb 4:4:4  */

/* two non contiguous planes - one Y, one Cr + Cb interleaved  */
var V4L2_PIX_FMT_NV12M uint32 = v4l2_fourcc('N', 'M', '1', '2')        /* 12  Y/CbCr 4:2:0  */
var V4L2_PIX_FMT_NV21M uint32 = v4l2_fourcc('N', 'M', '2', '1')        /* 21  Y/CrCb 4:2:0  */
var V4L2_PIX_FMT_NV16M uint32 = v4l2_fourcc('N', 'M', '1', '6')        /* 16  Y/CbCr 4:2:2  */
var V4L2_PIX_FMT_NV61M uint32 = v4l2_fourcc('N', 'M', '6', '1')        /* 16  Y/CrCb 4:2:2  */
var V4L2_PIX_FMT_NV12MT uint32 = v4l2_fourcc('T', 'M', '1', '2')       /* 12  Y/CbCr 4:2:0 64x32 macroblocks */
var V4L2_PIX_FMT_NV12MT_16X16 uint32 = v4l2_fourcc('V', 'M', '1', '2') /* 12  Y/CbCr 4:2:0 16x16 macroblocks */

/* three planes - Y Cb, Cr */
var V4L2_PIX_FMT_YUV410 uint32 = v4l2_fourcc('Y', 'U', 'V', '9')  /*  9  YUV 4:1:0     */
var V4L2_PIX_FMT_YVU410 uint32 = v4l2_fourcc('Y', 'V', 'U', '9')  /*  9  YVU 4:1:0     */
var V4L2_PIX_FMT_YUV411P uint32 = v4l2_fourcc('4', '1', '1', 'P') /* 12  YVU411 planar */
var V4L2_PIX_FMT_YUV420 uint32 = v4l2_fourcc('Y', 'U', '1', '2')  /* 12  YUV 4:2:0     */
var V4L2_PIX_FMT_YVU420 uint32 = v4l2_fourcc('Y', 'V', '1', '2')  /* 12  YVU 4:2:0     */
var V4L2_PIX_FMT_YUV422P uint32 = v4l2_fourcc('4', '2', '2', 'P') /* 16  YVU422 planar */

/* three non contiguous planes - Y, Cb, Cr */
var V4L2_PIX_FMT_YUV420M uint32 = v4l2_fourcc('Y', 'M', '1', '2') /* 12  YUV420 planar */
var V4L2_PIX_FMT_YVU420M uint32 = v4l2_fourcc('Y', 'M', '2', '1') /* 12  YVU420 planar */
var V4L2_PIX_FMT_YUV422M uint32 = v4l2_fourcc('Y', 'M', '1', '6') /* 16  YUV422 planar */
var V4L2_PIX_FMT_YVU422M uint32 = v4l2_fourcc('Y', 'M', '6', '1') /* 16  YVU422 planar */
var V4L2_PIX_FMT_YUV444M uint32 = v4l2_fourcc('Y', 'M', '2', '4') /* 24  YUV444 planar */
var V4L2_PIX_FMT_YVU444M uint32 = v4l2_fourcc('Y', 'M', '4', '2') /* 24  YVU444 planar */

/* Bayer formats - see http://www.siliconimaging.com/RGB%20Bayer.htm */
var V4L2_PIX_FMT_SBGGR8 uint32 = v4l2_fourcc('B', 'A', '8', '1')  /*  8  BGBG.. GRGR.. */
var V4L2_PIX_FMT_SGBRG8 uint32 = v4l2_fourcc('G', 'B', 'R', 'G')  /*  8  GBGB.. RGRG.. */
var V4L2_PIX_FMT_SGRBG8 uint32 = v4l2_fourcc('G', 'R', 'B', 'G')  /*  8  GRGR.. BGBG.. */
var V4L2_PIX_FMT_SRGGB8 uint32 = v4l2_fourcc('R', 'G', 'G', 'B')  /*  8  RGRG.. GBGB.. */
var V4L2_PIX_FMT_SBGGR10 uint32 = v4l2_fourcc('B', 'G', '1', '0') /* 10  BGBG.. GRGR.. */
var V4L2_PIX_FMT_SGBRG10 uint32 = v4l2_fourcc('G', 'B', '1', '0') /* 10  GBGB.. RGRG.. */
var V4L2_PIX_FMT_SGRBG10 uint32 = v4l2_fourcc('B', 'A', '1', '0') /* 10  GRGR.. BGBG.. */
var V4L2_PIX_FMT_SRGGB10 uint32 = v4l2_fourcc('R', 'G', '1', '0') /* 10  RGRG.. GBGB.. */
/* 10bit raw bayer packed, 5 bytes for every 4 pixels */
var V4L2_PIX_FMT_SBGGR10P uint32 = v4l2_fourcc('p', 'B', 'A', 'A')
var V4L2_PIX_FMT_SGBRG10P uint32 = v4l2_fourcc('p', 'G', 'A', 'A')
var V4L2_PIX_FMT_SGRBG10P uint32 = v4l2_fourcc('p', 'g', 'A', 'A')
var V4L2_PIX_FMT_SRGGB10P uint32 = v4l2_fourcc('p', 'R', 'A', 'A')

/* 10bit raw bayer a-law compressed to 8 bits */
var V4L2_PIX_FMT_SBGGR10ALAW8 uint32 = v4l2_fourcc('a', 'B', 'A', '8')
var V4L2_PIX_FMT_SGBRG10ALAW8 uint32 = v4l2_fourcc('a', 'G', 'A', '8')
var V4L2_PIX_FMT_SGRBG10ALAW8 uint32 = v4l2_fourcc('a', 'g', 'A', '8')
var V4L2_PIX_FMT_SRGGB10ALAW8 uint32 = v4l2_fourcc('a', 'R', 'A', '8')

/* 10bit raw bayer DPCM compressed to 8 bits */
var V4L2_PIX_FMT_SBGGR10DPCM8 uint32 = v4l2_fourcc('b', 'B', 'A', '8')
var V4L2_PIX_FMT_SGBRG10DPCM8 uint32 = v4l2_fourcc('b', 'G', 'A', '8')
var V4L2_PIX_FMT_SGRBG10DPCM8 uint32 = v4l2_fourcc('B', 'D', '1', '0')
var V4L2_PIX_FMT_SRGGB10DPCM8 uint32 = v4l2_fourcc('b', 'R', 'A', '8')
var V4L2_PIX_FMT_SBGGR12 uint32 = v4l2_fourcc('B', 'G', '1', '2') /* 12  BGBG.. GRGR.. */
var V4L2_PIX_FMT_SGBRG12 uint32 = v4l2_fourcc('G', 'B', '1', '2') /* 12  GBGB.. RGRG.. */
var V4L2_PIX_FMT_SGRBG12 uint32 = v4l2_fourcc('B', 'A', '1', '2') /* 12  GRGR.. BGBG.. */
var V4L2_PIX_FMT_SRGGB12 uint32 = v4l2_fourcc('R', 'G', '1', '2') /* 12  RGRG.. GBGB.. */
var V4L2_PIX_FMT_SBGGR16 uint32 = v4l2_fourcc('B', 'Y', 'R', '2') /* 16  BGBG.. GRGR.. */

/* compressed formats */
var V4L2_PIX_FMT_MJPEG uint32 = v4l2_fourcc('M', 'J', 'P', 'G')       /* Motion-JPEG   */
var V4L2_PIX_FMT_JPEG uint32 = v4l2_fourcc('J', 'P', 'E', 'G')        /* JFIF JPEG     */
var V4L2_PIX_FMT_DV uint32 = v4l2_fourcc('d', 'v', 's', 'd')          /* 1394          */
var V4L2_PIX_FMT_MPEG uint32 = v4l2_fourcc('M', 'P', 'E', 'G')        /* MPEG-1/2/4 Multiplexed */
var V4L2_PIX_FMT_H264 uint32 = v4l2_fourcc('H', '2', '6', '4')        /* H264 with start codes */
var V4L2_PIX_FMT_H264_NO_SC uint32 = v4l2_fourcc('A', 'V', 'C', '1')  /* H264 without start codes */
var V4L2_PIX_FMT_H264_MVC uint32 = v4l2_fourcc('M', '2', '6', '4')    /* H264 MVC */
var V4L2_PIX_FMT_H263 uint32 = v4l2_fourcc('H', '2', '6', '3')        /* H263          */
var V4L2_PIX_FMT_MPEG1 uint32 = v4l2_fourcc('M', 'P', 'G', '1')       /* MPEG-1 ES     */
var V4L2_PIX_FMT_MPEG2 uint32 = v4l2_fourcc('M', 'P', 'G', '2')       /* MPEG-2 ES     */
var V4L2_PIX_FMT_MPEG4 uint32 = v4l2_fourcc('M', 'P', 'G', '4')       /* MPEG-4 part 2 ES */
var V4L2_PIX_FMT_XVID uint32 = v4l2_fourcc('X', 'V', 'I', 'D')        /* Xvid           */
var V4L2_PIX_FMT_VC1_ANNEX_G uint32 = v4l2_fourcc('V', 'C', '1', 'G') /* SMPTE 421M Annex G compliant stream */
var V4L2_PIX_FMT_VC1_ANNEX_L uint32 = v4l2_fourcc('V', 'C', '1', 'L') /* SMPTE 421M Annex L compliant stream */
var V4L2_PIX_FMT_VP8 uint32 = v4l2_fourcc('V', 'P', '8', '0')         /* VP8 */

/*  Vendor-specific formats   */
var V4L2_PIX_FMT_CPIA1 uint32 = v4l2_fourcc('C', 'P', 'I', 'A')        /* cpia1 YUV */
var V4L2_PIX_FMT_WNVA uint32 = v4l2_fourcc('W', 'N', 'V', 'A')         /* Winnov hw compress */
var V4L2_PIX_FMT_SN9C10X uint32 = v4l2_fourcc('S', '9', '1', '0')      /* SN9C10x compression */
var V4L2_PIX_FMT_SN9C20X_I420 uint32 = v4l2_fourcc('S', '9', '2', '0') /* SN9C20x YUV 4:2:0 */
var V4L2_PIX_FMT_PWC1 uint32 = v4l2_fourcc('P', 'W', 'C', '1')         /* pwc older webcam */
var V4L2_PIX_FMT_PWC2 uint32 = v4l2_fourcc('P', 'W', 'C', '2')         /* pwc newer webcam */
var V4L2_PIX_FMT_ET61X251 uint32 = v4l2_fourcc('E', '6', '2', '5')     /* ET61X251 compression */
var V4L2_PIX_FMT_SPCA501 uint32 = v4l2_fourcc('S', '5', '0', '1')      /* YUYV per line */
var V4L2_PIX_FMT_SPCA505 uint32 = v4l2_fourcc('S', '5', '0', '5')      /* YYUV per line */
var V4L2_PIX_FMT_SPCA508 uint32 = v4l2_fourcc('S', '5', '0', '8')      /* YUVY per line */
var V4L2_PIX_FMT_SPCA561 uint32 = v4l2_fourcc('S', '5', '6', '1')      /* compressed GBRG bayer */
var V4L2_PIX_FMT_PAC207 uint32 = v4l2_fourcc('P', '2', '0', '7')       /* compressed BGGR bayer */
var V4L2_PIX_FMT_MR97310A uint32 = v4l2_fourcc('M', '3', '1', '0')     /* compressed BGGR bayer */
var V4L2_PIX_FMT_JL2005BCD uint32 = v4l2_fourcc('J', 'L', '2', '0')    /* compressed RGGB bayer */
var V4L2_PIX_FMT_SN9C2028 uint32 = v4l2_fourcc('S', 'O', 'N', 'X')     /* compressed GBRG bayer */
var V4L2_PIX_FMT_SQ905C uint32 = v4l2_fourcc('9', '0', '5', 'C')       /* compressed RGGB bayer */
var V4L2_PIX_FMT_PJPG uint32 = v4l2_fourcc('P', 'J', 'P', 'G')         /* Pixart 73xx JPEG */
var V4L2_PIX_FMT_OV511 uint32 = v4l2_fourcc('O', '5', '1', '1')        /* ov511 JPEG */
var V4L2_PIX_FMT_OV518 uint32 = v4l2_fourcc('O', '5', '1', '8')        /* ov518 JPEG */
var V4L2_PIX_FMT_STV0680 uint32 = v4l2_fourcc('S', '6', '8', '0')      /* stv0680 bayer */
var V4L2_PIX_FMT_TM6000 uint32 = v4l2_fourcc('T', 'M', '6', '0')       /* tm5600/tm60x0 */
var V4L2_PIX_FMT_CIT_YYVYUY uint32 = v4l2_fourcc('C', 'I', 'T', 'V')   /* one line of Y then 1 line of VYUY */
var V4L2_PIX_FMT_KONICA420 uint32 = v4l2_fourcc('K', 'O', 'N', 'I')    /* YUV420 planar in blocks of 256 pixels */
var V4L2_PIX_FMT_JPGL uint32 = v4l2_fourcc('J', 'P', 'G', 'L')         /* JPEG-Lite */
var V4L2_PIX_FMT_SE401 uint32 = v4l2_fourcc('S', '4', '0', '1')        /* se401 janggu compressed rgb */
var V4L2_PIX_FMT_S5C_UYVY_JPG uint32 = v4l2_fourcc('S', '5', 'C', 'I') /* S5C73M3 interleaved UYVY/JPEG */
var V4L2_PIX_FMT_Y8I uint32 = v4l2_fourcc('Y', '8', 'I', ' ')          /* Greyscale 8-bit L/R interleaved */
var V4L2_PIX_FMT_Y12I uint32 = v4l2_fourcc('Y', '1', '2', 'I')         /* Greyscale 12-bit L/R interleaved */
var V4L2_PIX_FMT_Z16 uint32 = v4l2_fourcc('Z', '1', '6', ' ')          /* Depth data 16-bit */

/* SDR formats - used only for Software Defined Radio devices */
var V4L2_SDR_FMT_CU8 uint32 = v4l2_fourcc('C', 'U', '0', '8')    /* IQ u8 */
var V4L2_SDR_FMT_CU16LE uint32 = v4l2_fourcc('C', 'U', '1', '6') /* IQ u16le */
var V4L2_SDR_FMT_CS8 uint32 = v4l2_fourcc('C', 'S', '0', '8')    /* complex s8 */
var V4L2_SDR_FMT_CS14LE uint32 = v4l2_fourcc('C', 'S', '1', '4') /* complex s14le */
var V4L2_SDR_FMT_RU12LE uint32 = v4l2_fourcc('R', 'U', '1', '2') /* real u12le */

/* Touch formats - used for Touch devices */
var V4L2_TCH_FMT_DELTA_TD16 uint32 = v4l2_fourcc('T', 'D', '1', '6') /* 16-bit signed deltas */
var V4L2_TCH_FMT_DELTA_TD08 uint32 = v4l2_fourcc('T', 'D', '0', '8') /* 8-bit signed deltas */
var V4L2_TCH_FMT_TU16 uint32 = v4l2_fourcc('T', 'U', '1', '6')       /* 16-bit unsigned touch data */
var V4L2_TCH_FMT_TU08 uint32 = v4l2_fourcc('T', 'U', '0', '8')       /* 8-bit unsigned touch data */

/* priv field value to indicates that subsequent fields are valid. */
var V4L2_PIX_FMT_PRIV_MAGIC uint32 = 0xfeedcafe

/* Flags */
var V4L2_PIX_FMT_FLAG_PREMUL_ALPHA uint32 = 0x00000001

/* Frame Size and frame rate enumeration */
/*
 *	F R A M E   S I Z E   E N U M E R A T I O N
 */
const (
	V4L2_FRMSIZE_TYPE_DISCRETE   = 1
	V4L2_FRMSIZE_TYPE_CONTINUOUS = 2
	V4L2_FRMSIZE_TYPE_STEPWISE   = 3
)

type V4l2Frmsizeenum struct {
	Index       uint32 /* Frame size number */
	PixelFormat uint32 /* Pixel format */
	Type        uint32 /* Frame size type the device supports. */

	data [6]uint32

	/*
		union {
			struct v4l2_frmsize_discrete	discrete;
			struct v4l2_frmsize_stepwise	stepwise;
		};*/

	reserved [2]uint32 /* Reserved space for future use */
}

func (f V4l2Frmsizeenum) Discrete() V4l2Frmsize_discrete {
	ptr := uintptr(unsafe.Pointer(&f))
	ptr += 12 /*skip index, pixel_format and type*/

	return *(*V4l2Frmsize_discrete)(unsafe.Pointer(ptr))
}

func (f V4l2Frmsizeenum) Stepwise() V4l2Frmsize_discrete {
	ptr := uintptr(unsafe.Pointer(&f))
	ptr += 24 /*skip index, pixel_format and type*/

	return *(*V4l2Frmsize_discrete)(unsafe.Pointer(ptr))
}

type V4l2Frmsize_discrete struct {
	Width  uint32 /* Frame width [pixel] */
	Height uint32 /* Frame height [pixel] */
}

type V4l2Frmsize_stepwise struct {
	Min_width   uint32 /* Minimum frame width [pixel] */
	Max_width   uint32 /* Maximum frame width [pixel] */
	Step_width  uint32 /* Frame width step size [pixel] */
	Min_height  uint32 /* Minimum frame height [pixel] */
	Max_height  uint32 /* Maximum frame height [pixel] */
	Step_height uint32 /* Frame height step size [pixel] */
}

/**
 * struct v4l2_format - stream data format
 * @type:	enum v4l2_buf_type; type of the data stream
 * @pix:	definition of an image format
 * @pix_mp:	definition of a multiplanar image format
 * @win:	definition of an overlaid image
 * @vbi:	raw VBI capture or output parameters
 * @sliced:	sliced VBI capture or output parameters
 * @raw_data:	placeholder for future extensions and custom formats
 */
type V4l2Format struct {
	Type uint32

	data [200]byte
	//union {
	//	struct v4l2_pix_format		pix;     /* V4L2_BUF_TYPE_VIDEO_CAPTURE */
	//	struct v4l2_pix_format_mplane	pix_mp;  /* V4L2_BUF_TYPE_VIDEO_CAPTURE_MPLANE */
	//	struct v4l2_window		win;     /* V4L2_BUF_TYPE_VIDEO_OVERLAY */
	//	struct v4l2_vbi_format		vbi;     /* V4L2_BUF_TYPE_VBI_CAPTURE */
	//	struct v4l2_sliced_vbi_format	sliced;  /* V4L2_BUF_TYPE_SLICED_VBI_CAPTURE */
	//	struct v4l2_sdr_format		sdr;     /* V4L2_BUF_TYPE_SDR_CAPTURE */
	//	__u8	raw_data[200];                   /* user-defined */
	//} fmt;
}

func (f *V4l2Format) SetPixFormat(pixformat *V4l2PixFormat) {

	f.Type = V4L2_BUF_TYPE_VIDEO_CAPTURE

	t := (*V4l2PixFormat)(unsafe.Pointer(&f.data))
	t.Width = pixformat.Width
	t.Height = pixformat.Height
	t.Pixelformat = pixformat.Pixelformat
	t.Field = pixformat.Field

	//fff := *(*[204]byte)(unsafe.Pointer(f))
	//fmt.Printf("%v\n", fff)
}

/*
 *	V I D E O   I M A G E   F O R M A T
 */
type V4l2PixFormat struct {
	Width        uint32
	Height       uint32
	Pixelformat  uint32
	Field        uint32 /* enum v4l2_field */
	Bytesperline uint32 /* for padding, zero if unused */
	Sizeimage    uint32
	Colorspace   uint32 /* enum v4l2_colorspace */
	Priv         uint32 /* private data, depends on pixelformat */
	Flags        uint32 /* format flags (V4L2_PIX_FMT_FLAG_*) */
	Uycbcr_enc   uint32 /* enum v4l2_ycbcr_encoding */
	Uantization  uint32 /* enum v4l2_quantization */
	Xfer_func    uint32 /* enum v4l2_xfer_func */
}

/*
 *	M E M O R Y - M A P P I N G   B U F F E R S
 */
type V4l2RequestBuffers struct {
	Count    uint32
	Type     uint32 /* enum v4l2_buf_type */
	Memory   uint32 /* enum v4l2_memory */
	Reserved [2]uint32
}

/*v4l2_buf_type*/
const (
	V4L2_MEMORY_MMAP    = 1
	V4L2_MEMORY_USERPTR = 2
	V4L2_MEMORY_OVERLAY = 3
	V4L2_MEMORY_DMABUF  = 4
)

/**
 * struct v4l2_buffer - video buffer info
 * @index:	id number of the buffer
 * @type:	enum v4l2_buf_type; buffer type (type == *_MPLANE for
 *		multiplanar buffers);
 * @bytesused:	number of bytes occupied by data in the buffer (payload);
 *		unused (set to 0) for multiplanar buffers
 * @flags:	buffer informational flags
 * @field:	enum v4l2_field; field order of the image in the buffer
 * @timestamp:	frame timestamp
 * @timecode:	frame timecode
 * @sequence:	sequence count of this frame
 * @memory:	enum v4l2_memory; the method, in which the actual video data is
 *		passed
 * @offset:	for non-multiplanar buffers with memory == V4L2_MEMORY_MMAP;
 *		offset from the start of the device memory for this plane,
 *		(or a "cookie" that should be passed to mmap() as offset)
 * @userptr:	for non-multiplanar buffers with memory == V4L2_MEMORY_USERPTR;
 *		a userspace pointer pointing to this buffer
 * @fd:		for non-multiplanar buffers with memory == V4L2_MEMORY_DMABUF;
 *		a userspace file descriptor associated with this buffer
 * @planes:	for multiplanar buffers; userspace pointer to the array of plane
 *		info structs for this buffer
 * @length:	size in bytes of the buffer (NOT its payload) for single-plane
 *		buffers (when type != *_MPLANE); number of elements in the
 *		planes array for multi-plane buffers
 *
 * Contains data exchanged by application and driver using one of the Streaming
 * I/O methods.
 */
type V4l2Buffer struct {
	Index     uint32
	Type      uint32
	Bytesused uint32
	Flags     uint32
	Field     uint32
	timestamp [8]byte  //struct timeval		timestamp;
	timecode  [16]byte //struct v4l2_timecode	timecode;
	Sequence  uint32

	/* memory location */
	Memory uint32
	m      [4]byte
	/*
		union {
			__u32           offset;   //4
			unsigned long   userptr; //4
			struct v4l2_plane *planes;  //4
			__s32		fd; //4
		} m;*/
	Length    uint32
	Reserved2 uint32
	Reserved  uint32
}

func (b *V4l2Buffer) Offset() uint32 {
	return binary.LittleEndian.Uint32(b.m[:])
}

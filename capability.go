package webcam

import (
	"fmt"

	"github.com/jalasoft/go-webcam/v4l2"
)

var CAP_VIDEO_CAPTURE Capability = Capability{"V4L2_CAP_VIDEO_CAPTURE", v4l2.V4L2_CAP_VIDEO_CAPTURE}
var CAP_VIDEO_OUTPUT Capability = Capability{"V4L2_CAP_VIDEO_OUTPUT", v4l2.V4L2_CAP_VIDEO_OUTPUT}
var CAP_VIDEO_OVERLAY Capability = Capability{"V4L2_CAP_VIDEO_OVERLAY", v4l2.V4L2_CAP_VIDEO_OVERLAY}
var CAP_VBI_CAPTURE Capability = Capability{"V4L2_CAP_VBI_CAPTURE", v4l2.V4L2_CAP_VBI_CAPTURE}
var CAP_VBI_OUTPUT Capability = Capability{"V4L2_CAP_VBI_OUTPUT", v4l2.V4L2_CAP_VBI_OUTPUT}
var CAP_SLICED_VBI_CAPTURE Capability = Capability{"V4L2_CAP_SLICED_VBI_CAPTURE", v4l2.V4L2_CAP_SLICED_VBI_CAPTURE}
var CAP_SLICED_VBI_OUTPUT Capability = Capability{"V4L2_CAP_SLICED_VBI_OUTPUT", v4l2.V4L2_CAP_SLICED_VBI_OUTPUT}
var CAP_RDS_CAPTURE Capability = Capability{"V4L2_CAP_RDS_CAPTURE", v4l2.V4L2_CAP_RDS_CAPTURE}
var CAP_VIDEO_OUTPUT_OVERLAY Capability = Capability{"V4L2_CAP_VIDEO_OUTPUT_OVERLAY", v4l2.V4L2_CAP_VIDEO_OUTPUT_OVERLAY}
var CAP_HW_FREQ_SEEK Capability = Capability{"V4L2_CAP_HW_FREQ_SEEK", v4l2.V4L2_CAP_HW_FREQ_SEEK}
var CAP_RDS_OUTPUT Capability = Capability{"V4L2_CAP_RDS_OUTPUT", v4l2.V4L2_CAP_RDS_OUTPUT}
var CAP_VIDEO_CAPTURE_MPLANE Capability = Capability{"V4L2_CAP_VIDEO_CAPTURE_MPLANE", v4l2.V4L2_CAP_VIDEO_CAPTURE_MPLANE}
var CAP_VIDEO_OUTPUT_MPLANE Capability = Capability{"V4L2_CAP_VIDEO_OUTPUT_MPLANE", v4l2.V4L2_CAP_VIDEO_OUTPUT_MPLANE}
var CAP_VIDEO_M2M_MPLANE Capability = Capability{"V4L2_CAP_VIDEO_M2M_MPLANE", v4l2.V4L2_CAP_VIDEO_M2M_MPLANE}
var CAP_VIDEO_M2M Capability = Capability{"V4L2_CAP_VIDEO_M2M", v4l2.V4L2_CAP_VIDEO_M2M}
var CAP_TUNER Capability = Capability{"V4L2_CAP_TUNER", v4l2.V4L2_CAP_TUNER}
var CAP_AUDIO Capability = Capability{"V4L2_CAP_AUDIO", v4l2.V4L2_CAP_AUDIO}
var CAP_RADIO Capability = Capability{"V4L2_CAP_RADIO", v4l2.V4L2_CAP_RADIO}
var CAP_MODULATOR Capability = Capability{"V4L2_CAP_MODULATOR", v4l2.V4L2_CAP_MODULATOR}
var CAP_SDR_CAPTURE Capability = Capability{"V4L2_CAP_SDR_CAPTURE", v4l2.V4L2_CAP_SDR_CAPTURE}
var CAP_EXT_PIX_FORMAT Capability = Capability{"V4L2_CAP_EXT_PIX_FORMAT", v4l2.V4L2_CAP_EXT_PIX_FORMAT}
var CAP_SDR_OUTPUT Capability = Capability{"V4L2_CAP_SDR_OUTPUT", v4l2.V4L2_CAP_SDR_OUTPUT}
var CAP_READWRITE Capability = Capability{"V4L2_CAP_READWRITE", v4l2.V4L2_CAP_READWRITE}
var CAP_ASYNCIO Capability = Capability{"V4L2_CAP_ASYNCIO", v4l2.V4L2_CAP_ASYNCIO}
var CAP_STREAMING Capability = Capability{"V4L2_CAP_STREAMING", v4l2.V4L2_CAP_STREAMING}
var CAP_TOUCH Capability = Capability{"V4L2_CAP_TOUCH", v4l2.V4L2_CAP_TOUCH}
var CAP_DEVICE_CAPS Capability = Capability{"V4L2_CAP_DEVICE_CAPS", v4l2.V4L2_CAP_DEVICE_CAPS}

var AllCapabilities = []Capability{
	CAP_VIDEO_CAPTURE,
	CAP_VIDEO_OUTPUT,
	CAP_VIDEO_OVERLAY,
	CAP_VBI_CAPTURE,
	CAP_VBI_OUTPUT,
	CAP_SLICED_VBI_CAPTURE,
	CAP_SLICED_VBI_OUTPUT,
	CAP_RDS_CAPTURE,
	CAP_VIDEO_OUTPUT_OVERLAY,
	CAP_HW_FREQ_SEEK,
	CAP_RDS_OUTPUT,
	CAP_VIDEO_CAPTURE_MPLANE,
	CAP_VIDEO_OUTPUT_MPLANE,
	CAP_VIDEO_M2M_MPLANE,
	CAP_VIDEO_M2M,
	CAP_TUNER,
	CAP_AUDIO,
	CAP_RADIO,
	CAP_MODULATOR,
	CAP_SDR_CAPTURE,
	CAP_EXT_PIX_FORMAT,
	CAP_SDR_OUTPUT,
	CAP_READWRITE,
	CAP_ASYNCIO,
	CAP_STREAMING,
	CAP_TOUCH,
	CAP_DEVICE_CAPS,
}

type v4l2Capability struct {
	cap v4l2.V4l2Capability
}

func (c v4l2Capability) Driver() string {
	return string(c.cap.Driver[:])
}

func (c v4l2Capability) Card() string {
	return string(c.cap.Card[:])
}

func (c v4l2Capability) BusInfo() string {
	return string(c.cap.BusInfo[:])
}

func (c v4l2Capability) Version() uint32 {
	return c.cap.Version
}

func (c v4l2Capability) HasCapability(cap Capability) bool {
	return (c.cap.Capabilities & cap.Value) > 0
}

func (c v4l2Capability) AllCapabilities() []Capability {
	var result []Capability

	for _, cap := range AllCapabilities {
		if c.HasCapability(cap) {
			result = append(result, cap)
		}
	}

	return result
}

func (c v4l2Capability) String() string {
	return fmt.Sprintf("Capability[driver=%s,card=%s,bus=%s,version=%d]", c.Driver(), c.Card(), c.BusInfo(), c.Version())
}

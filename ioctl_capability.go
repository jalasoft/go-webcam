package webcam

// #include "v4l2-binding.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"log"
	"unsafe"
)

var CAP_VIDEO_CAPTURE Capability = Capability{"V4L2_CAP_VIDEO_CAPTURE", C.V4L2_CAP_VIDEO_CAPTURE}
var CAP_VIDEO_OUTPUT Capability = Capability{"V4L2_CAP_VIDEO_OUTPUT", C.V4L2_CAP_VIDEO_OUTPUT}
var CAP_VIDEO_OVERLAY Capability = Capability{"V4L2_CAP_VIDEO_OVERLAY", C.V4L2_CAP_VIDEO_OVERLAY}
var CAP_VBI_CAPTURE Capability = Capability{"V4L2_CAP_VBI_CAPTURE", C.V4L2_CAP_VBI_CAPTURE}
var CAP_VBI_OUTPUT Capability = Capability{"V4L2_CAP_VBI_OUTPUT", C.V4L2_CAP_VBI_OUTPUT}
var CAP_SLICED_VBI_CAPTURE Capability = Capability{"V4L2_CAP_SLICED_VBI_CAPTURE", C.V4L2_CAP_SLICED_VBI_CAPTURE}
var CAP_SLICED_VBI_OUTPUT Capability = Capability{"V4L2_CAP_SLICED_VBI_OUTPUT", C.V4L2_CAP_SLICED_VBI_OUTPUT}
var CAP_RDS_CAPTURE Capability = Capability{"V4L2_CAP_RDS_CAPTURE", C.V4L2_CAP_RDS_CAPTURE}
var CAP_VIDEO_OUTPUT_OVERLAY Capability = Capability{"V4L2_CAP_VIDEO_OUTPUT_OVERLAY", C.V4L2_CAP_VIDEO_OUTPUT_OVERLAY}
var CAP_HW_FREQ_SEEK Capability = Capability{"V4L2_CAP_HW_FREQ_SEEK", C.V4L2_CAP_HW_FREQ_SEEK}
var CAP_RDS_OUTPUT Capability = Capability{"V4L2_CAP_RDS_OUTPUT", C.V4L2_CAP_RDS_OUTPUT}
var CAP_VIDEO_CAPTURE_MPLANE Capability = Capability{"V4L2_CAP_VIDEO_CAPTURE_MPLANE", C.V4L2_CAP_VIDEO_CAPTURE_MPLANE}
var CAP_VIDEO_OUTPUT_MPLANE Capability = Capability{"V4L2_CAP_VIDEO_OUTPUT_MPLANE", C.V4L2_CAP_VIDEO_OUTPUT_MPLANE}
var CAP_VIDEO_M2M_MPLANE Capability = Capability{"V4L2_CAP_VIDEO_M2M_MPLANE", C.V4L2_CAP_VIDEO_M2M_MPLANE}
var CAP_VIDEO_M2M Capability = Capability{"V4L2_CAP_VIDEO_M2M", C.V4L2_CAP_VIDEO_M2M}
var CAP_TUNER Capability = Capability{"V4L2_CAP_TUNER", C.V4L2_CAP_TUNER}
var CAP_AUDIO Capability = Capability{"V4L2_CAP_AUDIO", C.V4L2_CAP_AUDIO}
var CAP_RADIO Capability = Capability{"V4L2_CAP_RADIO", C.V4L2_CAP_RADIO}
var CAP_MODULATOR Capability = Capability{"V4L2_CAP_MODULATOR", C.V4L2_CAP_MODULATOR}
var CAP_SDR_CAPTURE Capability = Capability{"V4L2_CAP_SDR_CAPTURE", C.V4L2_CAP_SDR_CAPTURE}
var CAP_EXT_PIX_FORMAT Capability = Capability{"V4L2_CAP_EXT_PIX_FORMAT", C.V4L2_CAP_EXT_PIX_FORMAT}
var CAP_SDR_OUTPUT Capability = Capability{"V4L2_CAP_SDR_OUTPUT", C.V4L2_CAP_SDR_OUTPUT}
var CAP_READWRITE Capability = Capability{"V4L2_CAP_READWRITE", C.V4L2_CAP_READWRITE}
var CAP_ASYNCIO Capability = Capability{"V4L2_CAP_ASYNCIO", C.V4L2_CAP_ASYNCIO}
var CAP_STREAMING Capability = Capability{"V4L2_CAP_STREAMING", C.V4L2_CAP_STREAMING}
var CAP_TOUCH Capability = Capability{"V4L2_CAP_TOUCH", C.V4L2_CAP_TOUCH}
var CAP_DEVICE_CAPS Capability = Capability{"V4L2_CAP_DEVICE_CAPS", C.V4L2_CAP_DEVICE_CAPS}

var allCapabilities = [...]Capability{
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

//-------------------------------------------------------------------------------------------------------
//CAPABILITY INTERFACE IMPL
//-------------------------------------------------------------------------------------------------------

type capability struct {
	driver     string
	card       string
	businfo    string
	version    uint32
	cap_mask   uint32
	cap_values []Capability
}

func (c capability) Driver() string {
	return c.driver
}

func (c capability) Card() string {
	return c.card
}

func (c capability) BusInfo() string {
	return c.businfo
}

func (c capability) Version() uint32 {
	return c.version
}

func (c capability) HasCapability(cc Capability) bool {
	for _, cap := range c.cap_values {
		if cap.Value == cc.Value {
			return true
		}
	}
	return false
}

func (c capability) Capabilities() []Capability {
	result := make([]Capability, len(c.cap_values))
	copy(result, c.cap_values)
	return result
}

func (c capability) AllPossibleCapabilities() []Capability {
	all := make([]Capability, len(allCapabilities))
	copy(all, allCapabilities[:])
	return all
}

func (c capability) String() string {
	return fmt.Sprintf("Capabilities[driver=%s,card=%s,bus_info=%s,version=%v,caps=%b]", c.driver, c.card, c.businfo, c.version, c.cap_mask)
}

//--------------------------------------------------------------------------------------------------
//CAPABILTITIES QUERY METHOD IMPL
//--------------------------------------------------------------------------------------------------

func (d *device) QueryCapabilities() (Capabilities, error) {

	log.Printf("Querying capabilities for file descriptor %d\n", d.file.Fd())

	cap, err := C.queryCapability(C.int(d.file.Fd()))

	defer C.free(unsafe.Pointer(cap))

	if err != nil {
		return capability{}, err
	}

	result := capability{}

	result.driver = readString((uintptr)(unsafe.Pointer(&cap.driver)), 16)
	result.card = readString((uintptr) (unsafe.Pointer(&cap.card)), 32)
	result.businfo = readString((uintptr) (unsafe.Pointer(&cap.bus_info)), 32)
	result.version = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&cap.version), 4))
	result.cap_mask = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&cap.capabilities), 4))
	result.cap_values = convertCapabilities(result.cap_mask)

	log.Printf("Capabilities for file descriptor %d successfully read.\n", d.file.Fd())

	return result, nil
}

func convertCapabilities(caps uint32) []Capability {
	result := []Capability{}
	for _, cap := range allCapabilities {
		if (caps & cap.Value) > 0 {
			result = append(result, cap)
		}
	}
	return result
}

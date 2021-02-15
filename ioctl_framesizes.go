package webcam

// #include "v4l2-binding.h"
import "C"

import (
	"encoding/binary"
	"syscall"
	"unsafe"
)

//---------------------------------------------------------------------------------------------------
//FRAME SIZES
//---------------------------------------------------------------------------------------------------

type frameSizes struct {
	discrete []DiscreteFrameSize
	stepwise []StepwiseFrameSize
}

func (f frameSizes) Discrete() []DiscreteFrameSize {
	return f.discrete
}

func (f frameSizes) Stepwise() []StepwiseFrameSize {
	return f.stepwise
}

//---------------------------------------------------------------------------------------------------
//QUERY FRAME SIZES
//---------------------------------------------------------------------------------------------------

func (d *device) QueryFrameSizes(f PixelFormat) (FrameSizes, error) {

	raw := f.(pixelFormat)

	discrete := []DiscreteFrameSize{}
	stepwise := []StepwiseFrameSize{}

	var info *C.struct_v4l2_frmsizeenum = nil
	var err error

	for {
		info, err = C.queryFramesizes(C.int(d.file.Fd()), C.uint(raw.value), info)

		if err == syscall.EINVAL {
			break
		}

		if err != nil {
			return nil, err
		}

		if info._type == C.V4L2_FRMSIZE_TYPE_DISCRETE {
			discrete = append(discrete, newDiscreteFramesize(f, info))
		}

		if info._type == C.V4L2_FRMSIZE_TYPE_STEPWISE {
			stepwise = append(stepwise, newStepwiseFramesize(f, info))
		}
	}

	return frameSizes{discrete: discrete, stepwise: stepwise}, nil
}

func newDiscreteFramesize(pixelformat PixelFormat, info *C.struct_v4l2_frmsizeenum) DiscreteFrameSize {

	ptr := (*C.struct_v4l2_frmsize_discrete)(unsafe.Pointer(&info.anon0))

	width := binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.width), 4))
	height := binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.height), 4))

	return DiscreteFrameSize{PixelFormat: pixelformat, Width: width, Height: height}
}

func newStepwiseFramesize(pixelFormat PixelFormat, info *C.struct_v4l2_frmsizeenum) StepwiseFrameSize {
	ptr := (*C.struct_v4l2_frmsize_stepwise)(unsafe.Pointer(&info.anon0))

	result := StepwiseFrameSize{}
	result.PixelFormat = pixelFormat
	result.MinWidth = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.min_width), 4))
	result.MaxWidth = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.max_width), 4))
	result.StepWidth = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.step_width), 4))
	result.MinHeight = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.min_height), 4))
	result.MaxHeight = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.max_height), 4))
	result.StepHeight = binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&ptr.step_height), 4))

	return result
}

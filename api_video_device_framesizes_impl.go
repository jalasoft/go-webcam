package webcam

import (
	"log"
	"unsafe"

	//"github.com/jalasoft/go-webcam/v4l2"
	"webcam/v4l2"
)

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

//func (d *device) QueryFrameSizes(f PixelFormat) (FrameSizes, error) {

//discrete, stepwise, err := readFrameSizes(d.file.Fd(), f.Value)

/*
	if err != nil {
		return nil, err
	}

	return &frameSizes{discrete, stepwise}, nil*/
//	return nil, nil
//}

func readFrameSizes(fd uintptr, format uint32) ([]DiscreteFrameSize, []StepwiseFrameSize, error) {

	var discrete []DiscreteFrameSize
	var stepwise []StepwiseFrameSize

	var str v4l2.V4l2Frmsizeenum
	str.Index = 0
	str.PixelFormat = format

	for {
		ok, err := v4l2.EnumFrameSize(fd, &str)

		if err != nil {
			log.Printf("An error occured during reading frame size for index %d and format %d: %v\n", str.Index, format, err)
			return nil, nil, err
		}

		if !ok {
			return discrete, stepwise, nil
		}

		ptr := uintptr(unsafe.Pointer(&str))
		ptr += 12 /*skip index, pixel_format and type*/

		if str.Type == v4l2.V4L2_FRMSIZE_TYPE_DISCRETE {
			d := *(*DiscreteFrameSize)(unsafe.Pointer(ptr))
			discrete = append(discrete, d)
		}

		if str.Type == v4l2.V4L2_FRMSIZE_TYPE_STEPWISE {
			s := *(*StepwiseFrameSize)(unsafe.Pointer(ptr))
			stepwise = append(stepwise, s)
		}

		str.Index++
	}
}

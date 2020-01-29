package webcam

import (
	"os"
	"v4l2"
	"v4l2/ioctl"
)

type framesizes struct {
	file *os.File
}

func (f *framesizes) SupportsDiscrete(format uint32, width uint32, height uint32) (bool, error) {

	var result bool = false

	err := f.iterateFrameSizes(f.file.Fd(), format, func(str v4l2.V4l2Frmsizeenum) bool {
		if str.Type != v4l2.V4L2_FRMSIZE_TYPE_DISCRETE {
			return true
		}

		discrete := str.Discrete()

		if discrete.Width == width && discrete.Height == height {
			result = true
			return false
		}

		return true
	})

	if err != nil {
		return false, err
	}

	return result, nil
}
func (f *framesizes) AllDiscreteMJPEG() ([]DiscreteFrameSize, error) {
	return f.AllDiscrete(v4l2.V4L2_PIX_FMT_MJPEG)
}

func (f *framesizes) AllDiscrete(format uint32) ([]DiscreteFrameSize, error) {

	result := make([]DiscreteFrameSize, 0, 10)

	err := f.iterateFrameSizes(f.file.Fd(), format, func(str v4l2.V4l2Frmsizeenum) bool {

		if str.Type != v4l2.V4L2_FRMSIZE_TYPE_DISCRETE {
			return true
		}

		discrete := str.Discrete()

		result = append(result, DiscreteFrameSize{discrete.Width, discrete.Height})

		return true
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
* Callback function that accepts filled structure with frame size
 */
type frameSizeCallback func(str v4l2.V4l2Frmsizeenum) bool

/*
* local method that accepts consumer who does concrete logic for each frame size
 */
func (f *framesizes) iterateFrameSizes(fd uintptr, format uint32, callback frameSizeCallback) error {

	var index uint32 = 0
	for {

		var str v4l2.V4l2Frmsizeenum
		str.Index = index
		str.PixelFormat = format
		ok, err := ioctl.QueryFrameSize(f.file.Fd(), &str)

		if err != nil {
			return err
		}

		if !ok {
			return nil
		}

		callback(str)
		index++
	}
}

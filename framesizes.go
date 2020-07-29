package webcam

import (
	"github.com/jalasoft/go-webcam/v4l2"
)

func (d *device) SupportsDiscrete(format uint32, width uint32, height uint32) (bool, error) {

	var result bool = false

	err := d.iterateFrameSizes(d.file.Fd(), format, func(str v4l2.V4l2Frmsizeenum) bool {
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
func (d *device) AllDiscreteMJPEG() ([]DiscreteFrameSize, error) {
	return d.AllDiscrete(v4l2.V4L2_PIX_FMT_MJPEG)
}

func (d *device) AllDiscrete(format uint32) ([]DiscreteFrameSize, error) {

	result := make([]DiscreteFrameSize, 0, 10)

	err := d.iterateFrameSizes(d.file.Fd(), format, func(str v4l2.V4l2Frmsizeenum) bool {

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
func (d *device) iterateFrameSizes(fd uintptr, format uint32, callback frameSizeCallback) error {

	var index uint32 = 0
	for {

		var str v4l2.V4l2Frmsizeenum
		str.Index = index
		str.PixelFormat = format
		ok, err := v4l2.QueryFrameSize(d.file.Fd(), &str)

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

package webcam

import (
	"os"
	"v4l2"
	"v4l2/ioctl"
)

type supportedFormats struct {
	file *os.File
}

func (f supportedFormats) Supports(bufType uint32, format uint32) (bool, error) {

	var desc v4l2.V4l2Fmtdesc
	desc.Index = 0
	desc.Typ = bufType

	var index *uint32 = &desc.Index

	for {

		ok, err := ioctl.QueryFormat(f.file.Fd(), &desc)

		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}

		if desc.Pixelformat == format {
			return true, nil
		}

		*index++
	}
	return false, nil
}

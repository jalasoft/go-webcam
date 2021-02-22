package webcam

import (
	"errors"
	"fmt"
	"math"
)

const (
	DEFAULT_PIXEL_FORMAT = "V4L2_PIX_FMT_MJPEG"
)

//-------------------------------------------------------------------------------------------
//CAMERA METHOD IMPL
//-------------------------------------------------------------------------------------------

func (d *device) DiscreteFrameSize() DiscreteFrameSizeSelector {
	return discreteFrameSizeSelector{camera: d}
}

//--------------------------------------------------------------------------------------------
//IMPLEMENTATION OF DISCRETE FRAME SIZE SELECTOR
//--------------------------------------------------------------------------------------------

type discreteFrameSizeSelector struct {
	camera *device

	pixFmt     PixelFormat
	pixFmtName string
	width      uint32
	height     uint32
}

func (d discreteFrameSizeSelector) PixelFormat(pixFmt PixelFormat) DiscreteFrameSizeSelector {
	d.pixFmt = pixFmt
	return d
}

func (d discreteFrameSizeSelector) PixelFormatName(pixFmt string) DiscreteFrameSizeSelector {
	d.pixFmtName = pixFmt
	return d
}

func (d discreteFrameSizeSelector) Width(w uint32) DiscreteFrameSizeSelector {
	d.width = w
	return d
}

func (d discreteFrameSizeSelector) Height(h uint32) DiscreteFrameSizeSelector {
	d.height = h
	return d
}

func (d discreteFrameSizeSelector) Select() (DiscreteFrameSize, error) {
	pixFmt, err := d.getPixelFormat()

	if err != nil {
		return DiscreteFrameSize{}, err
	}

	w, h, err := d.getWidthHeight(pixFmt)

	if err != nil {
		return DiscreteFrameSize{}, err
	}

	return DiscreteFrameSize{pixFmt, w, h}, nil
}

func (d discreteFrameSizeSelector) getPixelFormat() (PixelFormat, error) {

	formats, err := d.camera.QueryFormats()

	if err != nil {
		return nil, err
	}

	var pixelFormat PixelFormat = d.pixFmt

	if pixelFormat == nil {
		pixelFormat = d.getFormatByName(formats)
	} else {
		d.checkFormat(formats, pixelFormat)
	}

	return pixelFormat, nil

}

func (d discreteFrameSizeSelector) getFormatByName(formats []PixelFormat) PixelFormat {

	for _, value := range formats {
		if value.Name() == d.pixFmtName {
			return value
		}
	}

	return d.getDefaultFormat(formats)
}

func (d discreteFrameSizeSelector) getDefaultFormat(formats []PixelFormat) PixelFormat {

	for _, value := range formats {
		if value.Name() == DEFAULT_PIXEL_FORMAT {
			return value
		}
	}

	return formats[0]
}

func (d discreteFrameSizeSelector) checkFormat(formats []PixelFormat, format PixelFormat) error {
	pixFmt, ok := format.(pixelFormat)

	if !ok {
		return errors.New("Provided pixel format is not one supplied by webcam.Webcam.QueryFormats()")
	}

	for _, format := range formats {
		if format == pixFmt {
			return nil
		}
	}

	return fmt.Errorf("Provided PixelFormat %v not found.", pixFmt)
}

func (d discreteFrameSizeSelector) getWidthHeight(pixFormat PixelFormat) (uint32, uint32, error) {

	frmSizes, err := d.camera.QueryFrameSizes(pixFormat)

	if err != nil {
		return 0, 0, err
	}

	discrete := frmSizes.Discrete()

	width := d.getWidth(discrete)

	matchingHeights := []DiscreteFrameSize{}

	for _, size := range discrete {
		if size.Width == width {
			matchingHeights = append(matchingHeights, size)
		}
	}

	height := d.getHeight(matchingHeights)

	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

func (d discreteFrameSizeSelector) getWidth(frmSizes []DiscreteFrameSize) uint32 {
	if d.width <= 0 {
		return frmSizes[0].Width
	}

	min_diff := float64(100000)
	best_width := uint32(0)

	for _, size := range frmSizes {
		diff := math.Abs(float64(d.width - size.Width))
		if diff < min_diff {
			min_diff = diff
			best_width = size.Width
		}
	}
	return best_width
}

func (d discreteFrameSizeSelector) getHeight(frmSizes []DiscreteFrameSize) uint32 {

	if d.height <= 0 {
		return frmSizes[0].Height
	}

	min_diff := float64(100000)
	best_width := uint32(0)

	for _, size := range frmSizes {
		diff := math.Abs(float64(d.height - size.Height))
		if diff < min_diff {
			min_diff = diff
			best_width = size.Width
		}
	}
	return best_width
}

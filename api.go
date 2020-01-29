package webcam

import (
	"errors"
	"fmt"
	"log"
	"os"
	"v4l2"
	"v4l2/ioctl"
)

func OpenVideoDevice(path string) (VideoDevice, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)

	log.Printf("Opening device %s\n", path)

	if err != nil {
		return nil, err
	}

	log.Println("Reading capability")
	cap, err := ioctl.QueryCapability(file.Fd())

	if err != nil {
		return nil, err
	}

	var dev *device = &device{file, v4l2Capability{cap}, supportedFormats{file}, &framesizes{file}, &camera{file}}

	if !dev.Capability().HasCapability(v4l2.V4L2_CAP_VIDEO_CAPTURE) {
		return nil, errors.New(fmt.Sprintf("Device %s is not a video capturing device.", dev.Name()))
	}

	if !dev.Capability().HasCapability(v4l2.V4L2_CAP_STREAMING) {
		return nil, errors.New(fmt.Sprintf("Device %s is not able to stream frames.", dev.Name()))
	}

	log.Printf("Device %s is a video device", file.Name())
	return dev, nil
}

//-------------------------------------------------------------------------
//MAIN INTERFACE
//-------------------------------------------------------------------------

type VideoDevice interface {
	Name() string
	Capability() Capability
	Formats() SupportedFormats
	FrameSizes() FrameSizes
	TakeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error)
	TakeSnapshotAsync(frameSize *DiscreteFrameSize, handler SnapshotHandler) error
	TakeSnapshotChan(frameSize *DiscreteFrameSize, ch chan Snapshot)
	Stream(framesize *DiscreteFrameSize, tick chan bool, snapshots chan<- Snapshot)
	Close() error
}

type Capability interface {
	Driver() string
	Card() string
	BusInfo() string
	Version() uint32
	HasCapability(cap uint32) bool
}

type SupportedFormats interface {
	Supports(bufType uint32, format uint32) (bool, error)
}

type FrameSizes interface {
	AllDiscrete(format uint32) ([]DiscreteFrameSize, error)
	AllDiscreteMJPEG() ([]DiscreteFrameSize, error)
	SupportsDiscrete(format uint32, width uint32, height uint32) (bool, error)
}

type DiscreteFrameSize struct {
	Width  uint32
	Height uint32
}

func (d DiscreteFrameSize) String() string {
	return fmt.Sprintf("DiscreteFrame[%dx%d]", d.Width, d.Height)
}

type Snapshot interface {
	FrameSize() *DiscreteFrameSize
	Length() uint32
	Data() []byte
}

type SnapshotHandler func(snapshot Snapshot)

package webcam

import (
	"fmt"
	"os"
)

func SearchVideoDevices() ([]string, error) {
	return searchVideoDevices()
}

func OpenVideoDevice(path string) (VideoDevice, error) {
	return openVideoDevice(path)
}

//-------------------------------------------------------------------------
//MAIN INTERFACE
//-------------------------------------------------------------------------

type VideoDevice interface {
	File() *os.File
	QueryCapabilities() (VideoDeviceCapabilities, error)
	QueryFormats() ([]PixelFormat, error)
	QueryFrameSizes(f PixelFormat) (FrameSizes, error)
	TakeSnapshot(format *PixelFormat, frameSize *DiscreteFrameSize) (Snapshot, error)
	//TakeSnapshotChan(frameSize *DiscreteFrameSize, ch chan Snapshot)
	//StreamByTicks(framesize *DiscreteFrameSize, tick chan bool, snapshots chan<- Snapshot)
	//StreamToWriter(framesize *DiscreteFrameSize, writer io.Writer, stop chan struct{})
	Close() error
}

type NameAndValue struct {
	Name  string
	Value uint32
}

//--------------------------------------------------------------------------------------
//CAPABILITIES
//--------------------------------------------------------------------------------------

type Capability NameAndValue

func (c Capability) String() string {
	return fmt.Sprintf("Capability[%s]", c.Name)
}

type VideoDeviceCapabilities interface {
	Driver() string
	Card() string
	BusInfo() string
	Version() uint32
	HasCapability(cap Capability) bool
	AllCapabilities() []Capability
}

//--------------------------------------------------------------------------------------
//FORMATS
//--------------------------------------------------------------------------------------

type PixelFormat NameAndValue

//---------------------------------------------------------------------------------------
//FRAME SIZES
//---------------------------------------------------------------------------------------

type FrameSizeType uint32

type FrameSizes interface {
	Discrete() []DiscreteFrameSize
	Stepwise() []StepwiseFrameSize
}

type DiscreteFrameSize struct {
	Width  uint32
	Height uint32
}

type StepwiseFrameSize struct {
	MinWidth   uint32
	MaxWidth   uint32
	StepWidth  uint32
	MinHeight  uint32
	MaxHeight  uint32
	StepHeight uint32
}

func (d DiscreteFrameSize) String() string {
	return fmt.Sprintf("DiscreteFrame[%dx%d]", d.Width, d.Height)
}

//----------------------------------------------------------------------------------------
//SNAPSHOT
//----------------------------------------------------------------------------------------

type Snapshot interface {
	FrameSize() *DiscreteFrameSize
	Length() uint32
	Data() []byte
}

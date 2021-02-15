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
	QueryCapabilities() (Capabilities, error)
	QueryFormats() ([]PixelFormat, error)
	QueryFrameSizes(f PixelFormat) (FrameSizes, error)
	TakeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error)
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

type Capabilities interface {
	Driver() string
	Card() string
	BusInfo() string
	Version() uint32
	HasCapability(cap Capability) bool
	Capabilities() []Capability
	AllPossibleCapabilities() []Capability
}

//--------------------------------------------------------------------------------------
//FORMATS
//--------------------------------------------------------------------------------------

type PixelFormat interface {
	Name() string
	Description() string
}

//---------------------------------------------------------------------------------------
//FRAME SIZES
//---------------------------------------------------------------------------------------

type FrameSizes interface {
	Discrete() []DiscreteFrameSize
	Stepwise() []StepwiseFrameSize
}

type DiscreteFrameSize struct {
	PixelFormat PixelFormat
	Width       uint32
	Height      uint32
}

func (d DiscreteFrameSize) String() string {
	return fmt.Sprintf("DiscreteFrame[%dx%d]", d.Width, d.Height)
}

type StepwiseFrameSize struct {
	PixelFormat PixelFormat
	MinWidth    uint32
	MaxWidth    uint32
	StepWidth   uint32
	MinHeight   uint32
	MaxHeight   uint32
	StepHeight  uint32
}

func (s StepwiseFrameSize) String() string {
	return fmt.Sprintf("StepwiseFrame[min_w=%d,max_w=%d,min_h=%d,max_height=%d,step_w=%d,step_h=%d]", s.MinWidth, s.MaxWidth, s.MinHeight, s.MaxHeight, s.StepWidth, s.StepHeight)
}

//----------------------------------------------------------------------------------------
//SNAPSHOT
//----------------------------------------------------------------------------------------

type Snapshot interface {
	Data() []byte
}

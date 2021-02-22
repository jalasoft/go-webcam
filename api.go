package webcam

import (
	"fmt"
	"os"
)

func FindWebcams() ([]WebcamInfo, error) {
	return findWebcams()
}

func OpenWebcam(path string) (Webcam, error) {
	return openWebcam(path)
}

//-------------------------------------------------------------------------
//MAIN INTERFACE
//-------------------------------------------------------------------------

type WebcamInfo struct {
	File *os.File
	Name string
}

func (i WebcamInfo) String() string {
	return fmt.Sprintf("Webcam[%s:%s]", i.Name, i.File.Name())
}

type Webcam interface {
	File() *os.File
	QueryCapabilities() (Capabilities, error)
	QueryFormats() ([]PixelFormat, error)
	QueryFrameSizes(f PixelFormat) (FrameSizes, error)
	DiscreteFrameSize() DiscreteFrameSizeSelector
	TakeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error)
	StreamSnapshots(framesize *DiscreteFrameSize, snapChan chan Snapshot, errChan chan error, stop chan bool)
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
//FRAME SIZE SELECTOR
//----------------------------------------------------------------------------------------

type DiscreteFrameSizeSelector interface {
	PixelFormat(pixFmt PixelFormat) DiscreteFrameSizeSelector
	PixelFormatName(pixFmt string) DiscreteFrameSizeSelector
	Width(w uint32) DiscreteFrameSizeSelector
	Height(h uint32) DiscreteFrameSizeSelector
	Select() (DiscreteFrameSize, error)
}

//----------------------------------------------------------------------------------------
//SNAPSHOT
//----------------------------------------------------------------------------------------

type Snapshot interface {
	Data() []byte
}

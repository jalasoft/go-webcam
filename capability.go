package webcam

import (
	"fmt"
	"v4l2"
)

type v4l2Capability struct {
	cap v4l2.V4l2Capability
}

func (c v4l2Capability) Driver() string {
	return string(c.cap.Driver[:])
}

func (c v4l2Capability) Card() string {
	return string(c.cap.Card[:])
}

func (c v4l2Capability) BusInfo() string {
	return string(c.cap.BusInfo[:])
}

func (c v4l2Capability) Version() uint32 {
	return c.cap.Version
}

func (c v4l2Capability) HasCapability(cap uint32) bool {
	return (c.cap.Capabilities & cap) > 0
}

func (c v4l2Capability) String() string {
	return fmt.Sprintf("Capability[driver=%s,card=%s,bus=%s,version=%d]", c.Driver(), c.Card(), c.BusInfo(), c.Version())
}

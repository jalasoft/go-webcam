package webcam

//"fmt"

//"github.com/jalasoft/go-webcam/v4l2"

//"github.com/jalasoft/go-webcam/v4l2"

/*
func (d *device) QueryCapabilities() (VideoDeviceCapabilities, error) {
	cap, err := v4l2.QueryCapability(d.file.Fd())

	if err != nil {
		return nil, err
	}

	return videoDeviceCapabilities{cap}, nil
}

type videoDeviceCapabilities struct {
	cap v4l2.V4l2Capability
}

func (c videoDeviceCapabilities) Driver() string {
	return string(c.cap.Driver[:])
}

func (c videoDeviceCapabilities) Card() string {
	return string(c.cap.Card[:])
}

func (c videoDeviceCapabilities) BusInfo() string {
	return string(c.cap.BusInfo[:])
}

func (c videoDeviceCapabilities) Version() uint32 {
	return c.cap.Version
}

func (c videoDeviceCapabilities) HasCapability(cap Capability) bool {
	return (c.cap.Capabilities & cap.Value) > 0
}

func (c videoDeviceCapabilities) AllCapabilities() []Capability {
	var result []Capability

	for _, cap := range AllCapabilities {
		if c.HasCapability(cap) {
			result = append(result, cap)
		}
	}

	return result
}

func (c videoDeviceCapabilities) String() string {
	return fmt.Sprintf("Details[driver=%s,card=%s,bus=%s,version=%d]", c.Driver(), c.Card(), c.BusInfo(), c.Version())
}*/

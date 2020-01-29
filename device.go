package webcam

import (
	"log"
	"os"
)

type device struct {
	file       *os.File
	capability v4l2Capability
	formats    supportedFormats
	framesizes *framesizes
	camera     *camera
}

func (d *device) Name() string {
	return d.file.Name()
}

func (d *device) Capability() Capability {
	return d.capability
}

func (d *device) Formats() SupportedFormats {
	return d.formats
}

func (d *device) FrameSizes() FrameSizes {
	return d.framesizes
}

func (d *device) TakeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error) {
	return d.camera.takeSnapshot(frameSize)
}

func (d *device) TakeSnapshotAsync(frameSize *DiscreteFrameSize, handler SnapshotHandler) error {
	return d.camera.takeSnapshotAsync(frameSize, handler)
}

func (d *device) TakeSnapshotChan(frameSize *DiscreteFrameSize, ch chan Snapshot) {
	d.camera.takeSnapshotChan(frameSize, ch)
}

func (d *device) Stream(framesize *DiscreteFrameSize, ticks chan bool, snapshots chan<- Snapshot) {
	stream := &stream{file: d.file, frameSize: framesize}
	stream.stream(ticks, snapshots)
}

func (d *device) Close() error {
	log.Printf("Closing video device.\n")
	return d.file.Close()
}

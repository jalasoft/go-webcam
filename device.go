package webcam

import (
	"io"
	"log"
	"os"

	"github.com/jalasoft/go-webcam/v4l2"
)

type device struct {
	file       *os.File
	capability v4l2Capability
}

func (d *device) Name() string {
	return d.file.Name()
}

func (d *device) Capabilities() Capabilities {
	return d.capability
}

func (d *device) Formats() SupportedFormats {
	return d
}

func (d *device) Supports(bufType uint32, format uint32) (bool, error) {

	var desc v4l2.V4l2Fmtdesc
	desc.Index = 0
	desc.Typ = bufType

	var index *uint32 = &desc.Index

	for {

		ok, err := v4l2.QueryFormat(d.file.Fd(), &desc)

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

func (d *device) FrameSizes() FrameSizes {
	return d
}

func (d *device) TakeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error) {
	cam := &camera{file: d.file, frameSize: frameSize}
	return cam.takeSnapshot()
}

func (d *device) TakeSnapshotAsync(frameSize *DiscreteFrameSize, handler SnapshotHandler) error {
	cam := &camera{file: d.file, frameSize: frameSize}
	return cam.takeSnapshotAsync(handler)
}

func (d *device) TakeSnapshotChan(frameSize *DiscreteFrameSize, ch chan Snapshot) {
	cam := &camera{file: d.file, frameSize: frameSize}
	cam.takeSnapshotChan(ch)
}

func (d *device) StreamByTicks(framesize *DiscreteFrameSize, ticks chan bool, snapshots chan<- Snapshot) {
	cam := &camera{file: d.file, frameSize: framesize}
	cam.streamDrivenByTicks(ticks, snapshots)
}

func (d *device) StreamToWriter(framesize *DiscreteFrameSize, writer io.Writer, stop chan struct{}) {
	cam := &camera{file: d.file, frameSize: framesize}
	cam.streamToWriter(writer, stop)
}

func (d *device) Close() error {
	log.Printf("Closing video device.\n")
	return d.file.Close()
}

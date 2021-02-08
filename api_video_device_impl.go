package webcam

import (
	"log"
	"os"
	//"github.com/jalasoft/go-webcam/v4l2"
)

type device struct {
	file *os.File
}

func (d *device) File() *os.File {
	return d.file
}

func (d *device) TakeSnapshot(format *PixelFormat, frameSize *DiscreteFrameSize) (Snapshot, error) {
	cam := &camera{file: d.file, format: format, frameSize: frameSize}

	if err := cam.open(); err != nil {
		return nil, err
	}

	defer cam.close()

	if err := cam.queueDequeue(); err != nil {
		return nil, err
	}

	snap := &snapshot{cam.frameSize, cam.data, cam.length}

	return snap, nil
}

/*
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
*/

func (d *device) Close() error {
	log.Printf("Closing video device.\n")
	return d.file.Close()
}

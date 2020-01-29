package webcam

import (
	"fmt"
	"log"
	"os"
	"v4l2"
)

//-----------------------------------------------------
//SNAPSHOT
//-----------------------------------------------------

type snapshot struct {
	framesize *DiscreteFrameSize
	data      []byte
	length    uint32
}

func (s *snapshot) FrameSize() *DiscreteFrameSize {
	return s.framesize
}

func (s *snapshot) Data() []byte {
	return s.data
}

func (s *snapshot) Length() uint32 {
	return s.length
}

//-----------------------------------------------------
//STILL CAMERA
//-----------------------------------------------------

type camera struct {
	file *os.File
}

func (s *camera) takeSnapshotChan(frameSize *DiscreteFrameSize, ch chan Snapshot) {

	stream := &stream{file: s.file, frameSize: frameSize}

	if err := stream.open(); err != nil {
		log.Fatalf("%v\n", err)
	}

	defer stream.close()

	snap, err := stream.snapshot()

	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}

	ch <- snap
	close(ch)
}

func (s *camera) takeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error) {

	var sn *snapshot

	err := s.takeSnapshotAsync(frameSize, func(snap Snapshot) {
		var dataCopy []byte = make([]byte, snap.Length())
		copy(dataCopy, snap.Data())
		sn = &snapshot{snap.FrameSize(), dataCopy, snap.Length()}
	})

	if err != nil {
		return nil, err
	}

	return sn, nil
}

func (s *camera) takeSnapshotAsync(frameSize *DiscreteFrameSize, handler SnapshotHandler) error {

	log.Printf("Setting up frame size %dx%d", frameSize.Width, frameSize.Height)
	if err := setFrameSize(s.file.Fd(), frameSize, v4l2.V4L2_PIX_FMT_MJPEG); err != nil {
		return err
	}

	log.Printf("Frame size set up")
	log.Printf("Requesting buffer")
	if err := requestMmapBuffer(s.file.Fd()); err != nil {
		return err
	}
	log.Printf("Buffer requested successfully")
	log.Printf("Querying mmap buffer")
	offset, length, err := queryMmapBuffer(s.file.Fd())

	if err != nil {
		return err
	}

	log.Printf("Mmap buffer obtained. Offset=%v, length=%v\n", offset, length)
	log.Printf("Retrieving mapped memory block, offset=%d, length=%d", offset, length)

	data, err := mapBuffer(s.file.Fd(), offset, length)
	if err != nil {
		return err
	}

	log.Println("Activating streaming")
	if err := activateStreaming(s.file.Fd()); err != nil {
		return err
	}

	log.Println("Queueing buffer")
	var buffer v4l2.V4l2Buffer
	buffer.Index = uint32(0)
	buffer.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	buffer.Memory = v4l2.V4L2_MEMORY_MMAP

	if err := queueBuffer(s.file.Fd(), &buffer); err != nil {
		return err
	}
	log.Println(fmt.Sprintf("Buffer filled with %d bytes", buffer.Length))

	log.Println("Dequeuing the buffer")
	if err := dequeueBuffer(s.file.Fd(), &buffer); err != nil {
		return err
	}

	snapshot := &snapshot{frameSize, data, length}
	handler(snapshot)

	log.Printf("Releasing mapped memory block")
	if err := munmapBuffer(data); err != nil {
		return err
	}

	log.Println("Deactivating streaming")
	if err := deactivateStreaming(s.file.Fd()); err != nil {
		return err
	}

	return nil
}

//--------------------------------------------------------------------------------------------------
//STREAMING
//--------------------------------------------------------------------------------------------------

type stream struct {
	file      *os.File
	frameSize *DiscreteFrameSize
	length    uint32
	data      []byte
}

func (s *stream) stream(ticks chan bool, snapshots chan<- Snapshot) {

	if err := s.open(); err != nil {
		log.Fatalf("%v\n", err)
	}

	defer func() {
		if err := s.close(); err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	defer close(snapshots)

	for range ticks {
		fmt.Printf("snapshot\n")

		snap, err := s.snapshot()

		if err != nil {
			close(snapshots)
			ticks <- false
			log.Fatalf("%v\n", err)
		}

		ticks <- true
		snapshots <- snap
	}
}

func (s *stream) open() error {
	log.Printf("Setting up frame size %dx%d", s.frameSize.Width, s.frameSize.Height)
	if err := setFrameSize(s.file.Fd(), s.frameSize, v4l2.V4L2_PIX_FMT_MJPEG); err != nil {
		return err
	}

	log.Printf("Frame size set up")
	log.Printf("Requesting buffer")
	if err := requestMmapBuffer(s.file.Fd()); err != nil {
		return err
	}
	log.Printf("Buffer requested successfully")
	log.Printf("Querying mmap buffer")
	offset, length, err := queryMmapBuffer(s.file.Fd())

	if err != nil {
		return err
	}

	s.length = length

	log.Printf("Mmap buffer obtained. Offset=%v, length=%v\n", offset, length)
	log.Printf("Retrieving mapped memory block, offset=%d, length=%d", offset, length)

	data, err := mapBuffer(s.file.Fd(), offset, length)
	if err != nil {
		return err
	}

	s.data = data

	log.Println("Activating streaming")
	if err := activateStreaming(s.file.Fd()); err != nil {
		return err
	}

	return nil
}

func (s *stream) snapshot() (Snapshot, error) {

	log.Println("Queueing buffer")
	var buffer v4l2.V4l2Buffer
	buffer.Index = uint32(0)
	buffer.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	buffer.Memory = v4l2.V4L2_MEMORY_MMAP

	if err := queueBuffer(s.file.Fd(), &buffer); err != nil {
		return nil, err
	}
	log.Println(fmt.Sprintf("Buffer filled with %d bytes", buffer.Length))

	log.Println("Dequeuing the buffer")
	if err := dequeueBuffer(s.file.Fd(), &buffer); err != nil {
		return nil, err
	}

	snapshot := &snapshot{s.frameSize, s.data, s.length}
	return snapshot, nil
}

func (s *stream) close() error {
	log.Printf("Releasing mapped memory block")
	if err := munmapBuffer(s.data); err != nil {
		return err
	}

	log.Println("Deactivating streaming")
	if err := deactivateStreaming(s.file.Fd()); err != nil {
		return err
	}

	return nil
}

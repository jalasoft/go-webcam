package webcam

import (
	"os"
	//"github.com/jalasoft/go-webcam/v4l2"
	//"github.com/jalasoft/go-webcam/v4l2"
	//"webcam/v4l2"
)

//-----------------------------------------------------
//SNAPSHOT
//-----------------------------------------------------

/*
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
}*/

//-----------------------------------------------------
//CAMERA
//-----------------------------------------------------

type camera struct {
	file      *os.File
	format    *PixelFormat
	frameSize *DiscreteFrameSize
	length    uint32
	data      []byte
}

func (s *camera) open() error { /*
		log.Printf("Setting up pixel format %s with frame size %dx%d for file %v", s.format.Name, s.frameSize.Width, s.frameSize.Height, s.file.Name())
		if err := setFrameSize(s.file.Fd(), s.frameSize, s.format.Value); err != nil {
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
	*/
	return nil
}

/*
func (c *camera) takeSnapshotChan(ch chan Snapshot) {

	if err := c.open(); err != nil {
		log.Fatal(err)
	}

	defer c.close()

	if err := c.queueDequeue(); err != nil {
		log.Fatal(err)
	}

	snap := &snapshot{c.frameSize, c.data, c.length}
	ch <- snap
	close(ch)
}*/

/*
func (c *camera) takeSnapshot() (Snapshot, error) {

	if err := c.open(); err != nil {
		return nil, err
	}

	defer c.close()

	if err := c.queueDequeue(); err != nil {
		return nil, err
	}

	snap := &snapshot{c.frameSize, c.data, c.length}

	return snap, nil
}*/
/*
func (c *camera) streamToWriter(writer io.Writer, stop chan struct{}) {

	if err := c.open(); err != nil {
		log.Fatalf("%v\n", err)
	}

	defer func() {
		if err := c.close(); err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	for {
		select {
		case <-stop:
			break

		default:
			if err := c.queueDequeue(); err != nil {
				log.Fatal(err)
			}
			writer.Write(c.data)
		}
	}

	log.Printf("Stream skoncil")
}
*/
/*
func (c *camera) streamDrivenByTicks(ticks chan bool, snapshots chan<- Snapshot) {

	if err := c.open(); err != nil {
		log.Fatalf("%v\n", err)
	}

	defer func() {
		if err := c.close(); err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	for range ticks {
		log.Printf("snapshot requested\n")

		if err := c.queueDequeue(); err != nil {
			//name := formatToString[formatCode]
			log.Fatalf("chyba %v\n", err)
		}

		snap := &snapshot{c.frameSize, c.data, c.length}
		snapshots <- snap
	}

	log.Println("Ticks finished")
	close(snapshots)
}*/

//--------------------------------------------------------------------------------------
//IMPLEMENTATION DETAILS
//--------------------------------------------------------------------------------------
/*
func (s *camera) queueDequeue() error {

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

	//snapshot := &snapshot{s.frameSize, s.data, s.length}
	//return snapshot, nil

	return nil
}

func (s *camera) close() error {
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

//-----------------------HELPER FUNCTIONS----------------------------------------------
/*
func setFrameSize(fd uintptr, frameSize *DiscreteFrameSize, pixelFormat uint32) error {
	var format v4l2.V4l2Format

	var pixFormat v4l2.V4l2PixFormat
	pixFormat.Width = frameSize.Width
	pixFormat.Height = frameSize.Height
	pixFormat.Pixelformat = pixelFormat
	//pixFormat.Field = v4l2.V4L2_FIELD_NONE

	format.SetPixFormat(&pixFormat)

	return v4l2.SetFrameSize(fd, &format)
}*/

/*
func requestMmapBuffer(fd uintptr) error {

	var request v4l2.V4l2RequestBuffers
	request.Count = 1
	request.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	request.Memory = v4l2.V4L2_MEMORY_MMAP

	return v4l2.RequestBuffer(fd, &request)
}*/
/*
func queryMmapBuffer(fd uintptr) (uint32, uint32, error) {

	buffer := &v4l2.V4l2Buffer{}
	buffer.Index = uint32(0)
	buffer.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	buffer.Memory = v4l2.V4L2_MEMORY_MMAP

	v4l2.QueryBuffer(fd, buffer)

	return buffer.Offset(), buffer.Length, nil
}

func mapBuffer(fd uintptr, offset uint32, length uint32) ([]byte, error) {
	return syscall.Mmap(int(fd), int64(offset), int(length), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
}

func munmapBuffer(data []byte) error {
	return syscall.Munmap(data)
}

func activateStreaming(fd uintptr) error {
	return v4l2.ActivateStreaming(fd, v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE)
}

func deactivateStreaming(fd uintptr) error {
	return v4l2.DeactivateStreaming(fd, v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE)
}

func queueBuffer(fd uintptr, buffer *v4l2.V4l2Buffer) error {

	if err := v4l2.QueueBuffer(fd, buffer); err != nil {
		return err
	}

	return nil
}

func dequeueBuffer(fd uintptr, buffer *v4l2.V4l2Buffer) error {

	err := v4l2.DequeueBuffer(fd, buffer)

	if err != nil {
		return err
	}

	return nil
}
*/

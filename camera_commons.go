package webcam

import (
	"syscall"
	"github.com/jalasoft/go-webcam/v4l2"
)

func setFrameSize(fd uintptr, frameSize *DiscreteFrameSize, pixelFormat uint32) error {
	var format v4l2.V4l2Format

	var pixFormat v4l2.V4l2PixFormat
	pixFormat.Width = frameSize.Width
	pixFormat.Height = frameSize.Height
	pixFormat.Pixelformat = pixelFormat
	pixFormat.Field = v4l2.V4L2_FIELD_NONE

	format.SetPixFormat(&pixFormat)

	return v4l2.SetFrameSize(fd, &format)
}

func requestMmapBuffer(fd uintptr) error {

	var request v4l2.V4l2RequestBuffers
	request.Count = 1
	request.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	request.Memory = v4l2.V4L2_MEMORY_MMAP

	return v4l2.RequestBuffer(fd, &request)
}

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

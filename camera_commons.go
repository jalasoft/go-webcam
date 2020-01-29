package webcam

import (
	"syscall"
	"v4l2"
	"v4l2/ioctl"
)

func setFrameSize(fd uintptr, frameSize *DiscreteFrameSize, pixelFormat uint32) error {
	var format v4l2.V4l2Format

	var pixFormat v4l2.V4l2PixFormat
	pixFormat.Width = frameSize.Width
	pixFormat.Height = frameSize.Height
	pixFormat.Pixelformat = pixelFormat
	pixFormat.Field = v4l2.V4L2_FIELD_NONE

	format.SetPixFormat(&pixFormat)

	return ioctl.SetFrameSize(fd, &format)
}

func requestMmapBuffer(fd uintptr) error {

	var request v4l2.V4l2RequestBuffers
	request.Count = 1
	request.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	request.Memory = v4l2.V4L2_MEMORY_MMAP

	return ioctl.RequestBuffer(fd, &request)
}

func queryMmapBuffer(fd uintptr) (uint32, uint32, error) {

	buffer := &v4l2.V4l2Buffer{}
	buffer.Index = uint32(0)
	buffer.Type = v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE
	buffer.Memory = v4l2.V4L2_MEMORY_MMAP

	ioctl.QueryBuffer(fd, buffer)

	return buffer.Offset(), buffer.Length, nil
}

func mapBuffer(fd uintptr, offset uint32, length uint32) ([]byte, error) {
	return syscall.Mmap(int(fd), int64(offset), int(length), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
}

func munmapBuffer(data []byte) error {
	return syscall.Munmap(data)
}

func activateStreaming(fd uintptr) error {
	return ioctl.ActivateStreaming(fd, v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE)
}

func deactivateStreaming(fd uintptr) error {
	return ioctl.DeactivateStreaming(fd, v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE)
}

func queueBuffer(fd uintptr, buffer *v4l2.V4l2Buffer) error {

	if err := ioctl.QueueBuffer(fd, buffer); err != nil {
		return err
	}

	return nil
}

func dequeueBuffer(fd uintptr, buffer *v4l2.V4l2Buffer) error {

	err := ioctl.DequeueBuffer(fd, buffer)

	if err != nil {
		return err
	}

	return nil
}

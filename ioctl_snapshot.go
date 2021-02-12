package webcam

// #include "v4l2-binding.h"
import "C"
import (
	"encoding/binary"
	"log"
	"unsafe"
)

func (d *device) TakeSnapshot(format *PixelFormat, frameSize *DiscreteFrameSize) (Snapshot, error) {

	err := setFrameSize(d, format, frameSize)

	if err != nil {
		return nil, err
	}

	err = requestBuffer(d)

	if err != nil {
		return nil, err
	}

	var requestedBuffer *C.struct_v4l2_buffer

	requestedBuffer, err = queryBuffer(d)

	if err != nil {
		return nil, err
	}

	mappedMemoryPtr, err := mmap(d, requestedBuffer)

	if err != nil {
		return nil, err
	}


	defer C.free(unsafe.Pointer(requestedBuffer))

	var buffer *C.struct_v4l2_buffer
	
	buffer, err = C.newBuffer()

	if err != nil {
		return nil, err
	}

	defer C.free(unsafe.Pointer(buffer))

	
	err = streamOn(d)

	if err != nil {
		return nil, err
	}

	//queue dequeue

	err = streamOff(d)

	if err != nil {
		return nil, err
	}

	err = munmap(mappedMemoryPtr, requestedBuffer)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func setFrameSize(d *device, format *PixelFormat, frameSize *DiscreteFrameSize) error {

	raw := (*format).(pixelFormat)
	_, err := C.setDiscreteFrameSize(C.int(d.file.Fd()), C.uint(raw.value), C.uint(frameSize.Width), C.uint(frameSize.Height))

	return err
}

func requestBuffer(d *device) error {

	_, err := C.requestBuffer(C.int(d.file.Fd()))

	return err
}

func queryBuffer(d *device) (*C.struct_v4l2_buffer, error) {

	buff, err := C.queryBuffer(C.int(d.file.Fd()))

	if err != nil {
		return nil, err
	}

	return buff, nil
}

func mmap(d *device, req_buffer *C.struct_v4l2_buffer) (uintptr, error) {


	ptr, err := C.mmap2(C.int(d.file.Fd()), req_buffer)

	if err != nil {
		return 0, err
	}

	var p uintptr = (uintptr)(unsafe.Pointer(ptr))

	return p, nil
}

func munmap(memPtr uintptr, req_buffer *C.struct_v4l2_buffer) error {
	length := binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(&req_buffer.length), 4))
	_, err := C.munmap2(unsafe.Pointer(memPtr), C.uint(length));

	log.Printf("CAJK: %d, %v", length, err)
	return err
}

func streamOn(d *device) error {

	_, err := C.streamOn(C.int(d.file.Fd()))

	return err
}

func streamOff(d *device) error {
	_, err := C.streamOff(C.int(d.file.Fd()))

	return err
}
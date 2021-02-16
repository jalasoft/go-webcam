package webcam

// #include "v4l2-binding.h"
import "C"
import (
	"encoding/binary"
	"log"
	"unsafe"
)

//-----------------------------------------------------------------------------
//SNAPSHOT INTERFACE IMPL
//-----------------------------------------------------------------------------

type snapshot struct {
	data []byte
}

func (s *snapshot) Data() []byte {
	return s.data
}

//------------------------------------------------------------------------------
//TAKE SNAPSHOT
//------------------------------------------------------------------------------

func (d *device) TakeSnapshot(frameSize *DiscreteFrameSize) (Snapshot, error) {

	err := setFrameSize(d, frameSize)

	if err != nil {
		return nil, err
	}

	_, err = C.requestBuffer(C.int(d.file.Fd()))

	if err != nil {
		return nil, err
	}

	var requestedBuffer *C.struct_v4l2_buffer

	requestedBuffer, err = C.queryBuffer(C.int(d.file.Fd()))

	if err != nil {
		return nil, err
	}

	defer C.free(unsafe.Pointer(requestedBuffer))

	mappedMemoryPtr, err := mmap(d, requestedBuffer)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = munmap(mappedMemoryPtr, requestedBuffer)

		if err != nil {
			log.Printf("Cannot munma memory region: %v\n", err)
		}
	}()

	var buffer *C.struct_v4l2_buffer

	buffer, err = C.newBuffer()

	if err != nil {
		return nil, err
	}

	defer C.free(unsafe.Pointer(buffer))

	_, err = C.streamOn(C.int(d.file.Fd()))

	if err != nil {
		return nil, err
	}

	_, err = C.queueBuffer(C.int(d.file.Fd()), buffer)

	if err != nil {
		return nil, err
	}

	_, err = C.dequeueBuffer(C.int(d.file.Fd()), buffer)

	if err != nil {
		return nil, err
	}

	bytes := copyBytes(mappedMemoryPtr, requestedBuffer)

	_, err = C.streamOff(C.int(d.file.Fd()))

	if err != nil {
		return nil, err
	}

	return &snapshot{bytes}, nil
}

func setFrameSize(d *device, frameSize *DiscreteFrameSize) error {

	raw := frameSize.PixelFormat.(pixelFormat)
	_, err := C.setDiscreteFrameSize(C.int(d.file.Fd()), C.uint(raw.value), C.uint(frameSize.Width), C.uint(frameSize.Height))

	return err
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
	_, err := C.munmap2(unsafe.Pointer(memPtr), C.uint(length))

	return err
}

func copyBytes(memPtr uintptr, req_buffer *C.struct_v4l2_buffer) []byte {
	return C.GoBytes(unsafe.Pointer(memPtr), C.int(req_buffer.length))
}

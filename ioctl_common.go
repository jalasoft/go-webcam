package webcam

import (
	"C"
	"unsafe"
	"strings"
)

func readString(ptr uintptr, len uint8) string {
	original := string(C.GoBytes(unsafe.Pointer(ptr), C.int(len)))
	return strings.TrimRight(original, "\u0000")
}